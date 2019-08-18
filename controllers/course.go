package controllers

import "fmt"

type CourseController struct {
	BaseController
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (this *CourseController) Get() {
	fmt.Println(123456789111)
	this.Data["json"] = 123
	this.ServeJSON()
}
