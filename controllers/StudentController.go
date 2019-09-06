package controllers

import (
	"fjsdxy-plus/helper"
	"fjsdxy-plus/models"
)

//错误代码103xx
type StudentController struct {
	BaseController
}

// @Title 绑定学号
// @Description 绑定水院学号
// @Param Authorization header string true "token"
// @Param studentId formData string true "学号"
// @Param password formData string true "密码"
// @Success 200 {string} 绑定结果&学生信息
// @router /bind [post]
func (this *StudentController) Bind() {

	studenId := this.GetString("studentId")
	password := helper.GenToken(this.GetString("password"))
	if studenId == "" || password == "" {
		this.Error(10300, "学号或者密码不能为空")
	}
	student := models.NewStudent()
	student.User = models.NewUser()
	student.User.Id = this.UserId
	student.StudentId = studenId
	student.Password = password
	if err := student.Bind(); err != nil {
		this.Error(10301, err.Error())
	}
	this.Success("绑定成功", map[string]interface{}{
		"studentName": student.StudentName,       //学生姓名
		"className":   student.Classes.ClassName, //班级名称
		"college":     student.Classes.College,   //院系
		"major":       student.Classes.Major,     //专业
	})
}

// @Title 解绑学号
// @Description 解除绑定学号
// @Param Authorization header string true "token"
// @Success 200 {string} 解除结果
// @router /unBind [post]
func (this *StudentController) UnBind() {
	student := models.NewStudent()
	student.User = models.NewUser()
	student.User.Id = this.UserId
	if err := student.UnBind(); err != nil {
		this.Error(10302, err.Error())
	}
	this.Success("解绑成功")
}

// @Title 学生信息
// @Description 获取绑定学生信息
// @Param Authorization header string true "token"
// @Success 200 {string} 学生信息
// @router /getInfo [get]
func (this *StudentController) GetInfo() {
	student := models.NewStudent()
	student.User = models.NewUser()
	student.User.Id = this.UserId
	if err := student.GetInfo(); err != nil {
		this.Error(10303, err.Error())
	}
	this.Success("获取成功", map[string]interface{}{
		"studentName": student.StudentName,       //学生姓名
		"className":   student.Classes.ClassName, //班级名称
		"college":     student.Classes.College,   //院系
		"major":       student.Classes.Major,     //专业
	})
}

// @Title 更新学生信息
// @Description 获取最新学生信息
// @Param Authorization header string true "token"
// @Success 200 {string} 学生信息
// @router /pull [get]
func (this *StudentController) PullInfo() {
	student := models.NewStudent()
	student.User = models.NewUser()
	student.User.Id = this.UserId
	if err := student.PullInfo(); err != nil {
		this.Error(10304, err.Error())
	}
	this.Success("更新成功", map[string]interface{}{
		"studentName": student.StudentName,       //学生姓名
		"className":   student.Classes.ClassName, //班级名称
		"college":     student.Classes.College,   //院系
		"major":       student.Classes.Major,     //专业
	})
}

// @Title 上课通知
// @Description 设置上课通知
// @Param Authorization header string true "token"
// @Success 200 {string} 设置结果
// @router /class_notice [post]
func (this *StudentController) ClassNotice() {
	student := models.NewStudent()
	student.User = models.NewUser()
	student.User.Id = this.UserId
	if err := student.ClassNotice(); err != nil {
		this.Error(10305, err.Error())
	}
	this.Success("设置成功", student.IsNotice)
}
