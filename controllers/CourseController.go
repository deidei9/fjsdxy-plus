package controllers

import (
	"fjsdxy-plus/helper"
	"fjsdxy-plus/models"
	"time"
)

//错误代码106xx
type CourseController struct {
	BaseController
}

// @Title 获取课程表
// @Description 获取课程表数据
// @Param Authorization header string true "token"
// @Param date path string true "日期 如2018-04-04"
// @Success 200 {string} 获取结果
// @router /get/:date [get]
func (this *CourseController) GetCourse() {
	date := this.GetString(":date")

	if date == "" {
		this.Error(10600, "日期不能为空")
	}
	course := models.NewCourse()
	course.Classes = models.NewClasses()
	course.Classes.Student = models.NewStudent()
	course.Classes.Student.User = models.NewUser()
	course.Classes.Student.User.Id = this.UserId
	//获取当前学期
	week := models.NewWeek()
	if err := models.DB().QueryTable(models.GetTable("week")).Filter("today", date).One(week, "Term", "Weekly"); err != nil {
		this.Error(10601, "本周不在校历内")
	}
	course.Term = week.Term
	course.Weekly = int(week.Weekly)
	//获取本地课程表
	data, err := course.GetCourseHost()
	if err != nil {
		//获取校务课表数据
		data, err := course.GetCourse()
		if err != nil {
			this.Error(10602, err.Error())
		}
		this.Success("获取成功", data)
	}
	this.Success("获取成功", data)
}

// @Title 下节课课程
// @Description 获取下节课课程
// @Param Authorization header string true "token"
// @Success 200 {string} 获取结果
// @router /get_next [get]
func (this *CourseController) GetNextClass() {
	//today := time.Now().Format("2006-01-02")
	today := "2019-09-10"
	week := models.NewWeek()
	week.Today = today
	if err := week.GetWeek(); err != nil {
		this.Error(10603, "您还在度假呢")
	}
	if week.Weekly == 5 || week.Weekly == 6 {
		this.Error(10604, "今天休息呢")
	}
	//获取当前时间课程位置
	hour := time.Now().Hour()  //小时
	min := time.Now().Minute() //分钟
	positon := helper.Time2Pos(int(week.Week), hour, min)
	//获取今天课程数据
	course := models.NewCourse()
	course.Classes = models.NewClasses()
	course.Classes.Student = models.NewStudent()
	course.Classes.Student.User = models.NewUser()
	course.Classes.Student.User.Id = this.UserId
	course.Term = week.Term
	course.Weekly = int(week.Weekly)
	//获取校务课表数据
	_, err := course.GetCourseHost()
	if err != nil {

		//获取校务课表数据
		_, err := course.GetCourse()
		if err != nil {
			this.Error(10605, err.Error())
		}
	}
	//匹配下节课
	for _, v := range positon {
		course.Position = v
		if err := course.GetNextClass(); err == nil {
			break
		}
	}
	next := map[string]interface{}{}
	if course.Id != 0 {
		next = map[string]interface{}{
			"name":    course.Name,
			"time":    helper.GetClassTime(course.Position),
			"class":   course.Class,
			"teacher": course.Teacher,
		}
	}
	this.Success("获取成功", map[string]interface{}{
		"date":   today,
		"week":   helper.GetWeek(week.Week),
		"weekly": week.Weekly,
		"course": next,
	})
}
