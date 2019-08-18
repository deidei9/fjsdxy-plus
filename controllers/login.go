package controllers

import (
	"fjsdxy-plus/models"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	BaseController
}

// @Title login
// @Description 用户登录
// @Param	body1	body2 	models.User	false	"body for user content"
// @Success 200 123
// @Failure 403 body is empty
// @router / [get]
func (this *LoginController) Get() {
	//users := models.NewUser()
	o := orm.NewOrm()
	student := models.NewWeek()
	o.QueryTable(models.GetTable("week")).Filter("Id", 1).One(student)
	//this.Data["json"], _ = models.GetById(
	//	"student",
	//	1,
	//	"id",
	//	"password",
	//	"user__phone",
	//	"classes__college",
	//)

	this.Data["json"] = student
	this.ServeJSON()

}
