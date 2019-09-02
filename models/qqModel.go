package models

import (
	"encoding/json"
	"errors"
	"fjsdxy-plus/helper"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
)

type qq struct {
	Id         int     `orm:"auto;pk;index" json:"id"`
	Openid     string  `orm:"size(255)" json:"openId"`
	Unionid    string  `orm:"size(255)" json:"unionid"`
	SessionKey string  `orm:"size(255)" json:"session_key"`
	NickName   string  `orm:"size(30)" json:"nickName"`
	AvatarUrl  string  `orm:"size(255)" json:"avatarUrl"`
	Gender     float64 `orm:"digits(1);decimals(0)" json:"gender"`
	User       *User   `orm:"rel(one)" json:"user_info"`
	Model
}

func NewQQ() *qq {
	return &qq{}
}
func (this *qq) Login(code string) error {
	//获取openid
	_ = this.GetOpenid(code)
	if this.Openid == "" || this.SessionKey == "" {
		return errors.New("获取openid失败")
	}
	qq := NewQQ()
	err := DB().QueryTable(GetTable("qq")).Filter("openid", this.Openid).RelatedSel().One(qq)
	//QQ用户未注册，注册用户
	if qq.User == nil {
		user := NewUser()
		user.Phone = ""
		user.Password = helper.GetMd5("123456")
		user.NickName = this.NickName
		user.Avator = this.AvatarUrl
		this.User = user
		fmt.Println(this.User.State)
		_, err = DB().Insert(this.User)
		_, err = DB().Insert(this)
		if err != nil {
			return errors.New("用户注册失败")
		}
	} else {
		//已注册用户，更新sessionKey
		qq.SessionKey = this.SessionKey
		if _, err := DB().Update(qq); err != nil {
			return errors.New("登录失败")
		}
		*this = *qq
	}

	return nil
}

// 从微信服务器获取openid等信息
func (this *qq) GetOpenid(code string) error {

	c := http.Client{}

	resp, err := c.Get(fmt.Sprintf("https://api.q.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", beego.AppConfig.String("qq::appid"), beego.AppConfig.String("qq::secret"), code))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(body, &this)
	return nil
}

//效验微信用户信息
func (this *qq) GetInfo(encryptedData, iv string) (err error) {

	aesKey, _ := helper.GetBase64(this.SessionKey)
	aesIv, _ := helper.GetBase64(iv)
	aesCipher, _ := helper.GetBase64(encryptedData)
	result, err := this.decryptData(aesCipher, aesKey, aesIv)
	if err != nil {
		return err
	}
	watermark := result["watermark"].(map[string]interface{})
	appid := watermark["appid"].(string)

	if appid != beego.AppConfig.String("qq::appid") {
		return nil
	}
	this.AvatarUrl = result["avatarUrl"].(string)
	this.NickName = result["nickName"].(string)
	this.Gender = result["gender"].(float64)

	if this.User.NickName == "" && this.User.Avator == "" {
		this.User.NickName = result["nickName"].(string)
		this.User.Avator = result["avatarUrl"].(string)
	}

	if _, err := DB().Update(this.User); err != nil {
		return errors.New("获取信息失败")
	}
	if _, err := DB().Update(this); err != nil {
		return errors.New("获取信息失败")
	}
	return nil
}

//QQ加密数据解密
func (this *qq) decryptData(aesCipher, aesKey, aesIv []byte) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	var err error
	var result []byte
	if result, err = helper.AesDecrypt(aesCipher, aesKey, aesIv); err != nil {
		return nil, errors.New("解密失败")
	}

	if err = json.Unmarshal(result, &data); err != nil {
		fmt.Println(err)
		return nil, errors.New("获取信息失败")
	}

	return data, nil
}
