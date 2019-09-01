package controllers

import (
	"fjsdxy-plus/models"
	"strconv"
)

//错误代码101xx
type ExamController struct {
	BaseController
}

// @Title 学生成绩
// @Description 获取学生成绩
// @Param Authorization header string true "token"
// @Success 200 {string} 成绩结果
// @router /get [get]
func (this *ExamController) GetInfo() {
	student := models.NewStudent()
	student.User = models.NewUser()
	student.User.Id = this.UserId
	if err := student.GetInfo(); err != nil {
		this.Error(10101, err.Error())
	}
	exam := models.NewExam()
	exam.Student = student
	data, err := exam.GetByStu()
	if err != nil {
		this.Error(10102, err.Error())
	}
	if len(data) == 0 {
		this.Success("无数据", map[string]interface{}{})
	}
	this.Success("获取成功", data)
}

// @Title 更新成绩
// @Description 获取最新学生成绩
// @Param Authorization header string true "token"
// @Success 200 {string} 成绩结果
// @router /pull [get]
func (this *ExamController) PullExam() {
	student := models.NewStudent()
	student.User = models.NewUser()
	student.User.Id = this.UserId
	if err := student.GetInfo(); err != nil {
		this.Error(10103, err.Error())
	}
	exam := models.NewExam()
	exam.Student = student
	num, err := exam.PullExam()
	if err != nil {
		this.Error(10104, err.Error())
	}

	this.Success("更新" + strconv.FormatInt(num, 10) + "条数据")
}
