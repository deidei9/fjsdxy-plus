package controllers

import (
	"fjsdxy-plus/helper"
	"fjsdxy-plus/models"
	"strconv"
)

//错误代码105xx
type QQController struct {
	BaseController
}

// @Title QQ登录
// @Description 提供QQ用户授权登录
// @Param code formData string true "临时登录凭证"
// @Success 200 {string} json web token
// @router /login [post]
func (this *QQController) Login() {
	code := this.GetString("code")
	if code == "" {
		this.Error(10500, "code不能为空")
	}
	qq := models.NewQQ()
	err := qq.Login(code)
	if err != nil {
		this.Error(10501, "登录失败")
	}
	uid := qq.User.Id
	token := helper.GenToken(strconv.Itoa(uid))
	this.Success("登录成功", map[string]interface{}{
		"token": token,
	})
}

// @Title 获取QQ用户信息
// @Description 根据QQ提供的encryptedData解密用户信息
// @Param Authorization header string true "token"
// @Param encryptedData formData string true "加密数据"
// @Param iv formData string true "解密IV"
// @Success 200 {string} json web token
// @Failure 400 获取失败
// @router /getInfo [post]
func (this *QQController) GetInfo() {
	encryptedData := this.GetString("encryptedData")
	if encryptedData == "" {
		this.Error(10502, "encryptedData不能为空")
	}
	iv := this.GetString("iv")
	if iv == "" {
		this.Error(10503, "iv不能为空")
	}
	qq := models.NewQQ()
	if err := models.DB().QueryTable(models.GetTable("qq")).Filter("User__Id", this.UserId).RelatedSel().One(qq); err != nil {
		this.Error(10504, "用户未登录")
	}
	if err := qq.GetInfo(encryptedData, iv); err != nil {
		this.Error(10505, "获取用户信息失败")
	}

	this.Success("获取成功", qq)
}
