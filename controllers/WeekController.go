package controllers

import (
	"fjsdxy-plus/helper"
	"fjsdxy-plus/models"
	"time"
)

//错误代码102xx
type WeekController struct {
	BaseController
}

// @Title 教学周历
// @Description 获取教务处教学周历数据
// @Param term path string true "学期 如2018-2019-2"
// @Success 200 {string} 获取结果
// @router /pull/:term [get]
func (this *WeekController) Pull() {
	term := this.GetString(":term")
	if term == "" {
		this.Error(10200, "学期不能为空")
	}
	week := models.NewWeek()
	week.Term = term
	if err := week.Pull(); err != nil {
		this.Error(10201, err.Error())
	}
	this.Success("获取成功")
}

// @Title 指定周历数据
// @Description 获取指定周历数据
// @Param Authorization header string true "token"
// @Param today path string true "日期 如2018-04-04"
// @Success 200 {string} 获取结果
// @router /get/:today [get]
func (this *WeekController) GetWeek() {
	today := this.GetString(":today")
	if today == "" {
		this.Error(10202, "日期不能为空")
	}
	week := models.NewWeek()
	week.Today = today
	if err := week.GetWeek(); err != nil {
		this.Error(10203, err.Error())
	}
	this.Success("获取成功", map[string]interface{}{
		"term":   week.Term,   //学期
		"weekly": week.Weekly, //周次
		"week":   week.Week,   //星期
		"today":  week.Today,  //日期
	})
}

// @Title 当天周历
// @Description 获取当天周历
// @Success 200 {string} 获取结果
// @router /get_today [get]
func (this *WeekController) GetToday() {
	today := time.Now().Format("2006-01-02")
	toweek := time.Now().Weekday()

	this.Success("获取成功", map[string]interface{}{
		"week":  helper.GetWeek(toweek), //星期
		"today": today,                  //日期
	})
}

// @Title 当周校历
// @Description 获取当周校历
// @Param Authorization header string true "token"
// @Param date path string true "日期 如2018-04-04"
// @Success 200 {string} 获取结果
// @router /get_week_info/:date [get]
func (this *WeekController) GetWeekInfo() {
	date := this.GetString(":date")
	if date == "" {
		this.Error(10600, "日期不能为空")
	}
	week := models.NewWeek()
	week.Today = date
	if err := week.GetWeek(); err != nil {
		this.Error(10203, err.Error())
	}
	weeks := []models.Week{}
	if num, _ := models.DB().QueryTable(models.GetTable("week")).Filter("term", week.Term).Filter("weekly", week.Weekly).All(&weeks); num == 0 {
		this.Error(10202, "校历没有当周")
	}
	type ret struct {
		Week string `json:"week"`
		Time string `json:"time"`
	}
	var weekInfos []ret
	for _, v := range weeks {
		if v.Week != 0 && v.Week != 6 {
			weekInfos = append(weekInfos, ret{
				Week: helper.GetWeek2(v.Week),
				Time: v.Today[5:len(v.Today)],
			})
		}
	}
	this.Success("获取成功", map[string]interface{}{
		"weekly": week.Weekly,
		"isWeek": week.Week - 1,
		"weeks":  weekInfos,
	})
}
