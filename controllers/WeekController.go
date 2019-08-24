package controllers

import "fjsdxy-plus/models"

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
