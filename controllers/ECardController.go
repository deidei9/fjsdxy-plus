package controllers

import (
	"fjsdxy-plus/models"
	"github.com/pig0224/fjsdxy/cas/ecard"
)

//错误代码107xx
type ECardController struct {
	BaseController
}

// @Title 学生信息
// @Description 获取绑定学生信息
// @Param Authorization header string true "token"
// @Success 200 {string} 学生信息
// @router /get [get]
func (this *ECardController) Get() {
	student := models.NewStudent()
	student.User = models.NewUser()
	student.User.Id = this.UserId
	if err := student.GetInfo(); err != nil {
		this.Error(10700, err.Error())
	}
	c, err := student.LoginCAS()
	if err != nil {
		this.Error(10701, err.Error())
	}
	data, err := ecard.Get(c)
	if err != nil {
		this.Error(10702, err.Error())
	}
	this.Success("获取成功", data)
}
