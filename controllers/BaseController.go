package controllers

import (
	"fjsdxy-plus/helper"
	"github.com/astaxie/beego"
	"strconv"
)

type BaseController struct {
	beego.Controller
	UserId int
}

//初始化函数
func (this *BaseController) Prepare() {
	//获取用户ID
	this.GetUserId()
}

//获取登录的用户ID
func (this *BaseController) GetUserId() {
	token := this.Ctx.Input.Header("Authorization") //获取token
	uid, _ := strconv.Atoi(helper.CheckToken(token))
	if uid != 0 {
		this.UserId = uid
	}
}

//响应json
func (this *BaseController) ResponseJson(data map[string]interface{}) {
	this.Data["json"] = data
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) Error(code int, msg string) {
	ret := map[string]interface{}{"status": code, "msg": msg}
	this.ResponseJson(ret)
}

func (this *BaseController) Success(msg string, data ...interface{}) {
	ret := map[string]interface{}{"status": 200, "msg": msg}
	if len(data) > 0 {
		ret["data"] = data[0]
	}
	this.ResponseJson(ret)
}
