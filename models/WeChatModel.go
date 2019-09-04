package models

import (
	"encoding/json"
	"errors"
	"fjsdxy-plus/helper"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/xlstudio/wxbizdatacrypt"
	"io/ioutil"
	"net/http"
)

type Wechat struct {
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

func NewWechat() *Wechat {
	return &Wechat{}
}
func (this *Wechat) Login(code string) error {
	//获取openid
	_ = this.GetOpenid(code)
	if this.Openid == "" || this.SessionKey == "" {
		return errors.New("获取openid失败")
	}
	wechat := NewWechat()
	err := DB().QueryTable(GetTable("wechat")).Filter("openid", this.Openid).RelatedSel().One(wechat)
	//微信用户未注册，注册用户
	if wechat.User == nil {
		user := NewUser()
		user.Phone = ""
		user.Password = helper.GetMd5("123456")
		user.NickName = this.NickName
		user.Avator = this.AvatarUrl
		this.User = user
		_, err = DB().Insert(this.User)
		_, err = DB().Insert(this)
		if err != nil {
			return errors.New("用户注册失败")
		}
	} else {
		//已注册用户，更新sessionKey
		wechat.SessionKey = this.SessionKey
		if _, err := DB().Update(wechat); err != nil {
			return errors.New("登录失败")
		}
		*this = *wechat
	}

	return nil
}

// 从微信服务器获取openid等信息
func (this *Wechat) GetOpenid(code string) error {

	c := http.Client{}
	resp, err := c.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", beego.AppConfig.String("wechat::appid"), beego.AppConfig.String("wechat::secret"), code))
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
func (this *Wechat) GetInfo(encryptedData, iv string) (err error) {
	appID := beego.AppConfig.String("wechat::appid")
	pc := wxbizdatacrypt.WxBizDataCrypt{AppID: appID, SessionKey: this.SessionKey}
	result, err := pc.Decrypt(encryptedData, iv, false) //第三个参数解释： 需要返回 JSON 数据类型时 使用 true, 需要返回 map 数据类型时 使用 false
	if err != nil {
		helper.Logger.Error(err.Error())
		return errors.New("获取信息失败")
	}
	resData := result.(map[string]interface{})
	this.AvatarUrl = resData["avatarUrl"].(string)
	this.NickName = resData["nickName"].(string)
	this.Gender = resData["gender"].(float64)

	if this.User.NickName == "" && this.User.Avator == "" {
		this.User.NickName = resData["nickName"].(string)
		this.User.Avator = resData["avatarUrl"].(string)
	}

	if _, err := DB().Update(this.User); err != nil {
		helper.Logger.Error(err.Error())
		return errors.New("获取信息失败")
	}
	if _, err := DB().Update(this); err != nil {
		helper.Logger.Error(err.Error())
		return errors.New("获取信息失败")
	}
	return nil
}

//微信加密数据解密
func (this *Wechat) decryptData(aesCipher, aesKey, aesIv []byte) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	var err error
	var result []byte
	if result, err = helper.AesDecrypt(aesCipher, aesKey, aesIv); err != nil {
		helper.Logger.Error(err.Error())
		return nil, errors.New("解密失败")
	}

	if err = json.Unmarshal(result, &data); err != nil {
		helper.Logger.Error(err.Error())
		return nil, errors.New("获取信息失败")
	}
	return data, nil
}
