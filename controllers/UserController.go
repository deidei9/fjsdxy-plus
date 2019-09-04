package controllers

import (
	"fjsdxy-plus/models"
)

//错误代码104xx
type UserController struct {
	BaseController
}

// @Title 用户信息
// @Description 获取用户信息
// @Param Authorization header string true "token"
// @Success 200 {string} 用户信息
// @router /getInfo [get]
func (this *UserController) GetInfo() {
	//查询用户信息
	user := models.NewUser()
	user.Id = this.UserId
	if err := user.GetInfo(); err != nil {
		this.Error(10400, err.Error())
	}
	//查询学生信息
	student := models.NewStudent()
	student.User = user
	if err := student.GetInfo(); err == nil {
		this.Success("获取成功", map[string]interface{}{
			"userId":   user.Id,
			"nickName": user.NickName,
			"avator":   user.Avator,
			"state":    user.State,
			"student_info": map[string]interface{}{
				"isNotice":     student.IsNotice,        //上课通知 1是 0否
				"student_name": student.StudentName,     //班级名称
				"college":      student.Classes.College, //院系
				"major":        student.Classes.Major,   //专业
			},
		})
	}
	this.Success("获取成功", map[string]interface{}{
		"userId":       user.Id,
		"nickName":     user.NickName,
		"avator":       user.Avator,
		"state":        user.State,
		"student_info": "", //未绑定学号
	})
}

// @Title 审核状态
// @Description 获取小程序审核状态
// @Param platform path string true "审核平台"
// @Success 200 {string} 审核状态
// @router /support/:platform [get]
func (this *UserController) GetCheck() {
	platform := this.GetString(":platform")
	if platform == "" {
		this.Error(10600, "平台不能为空")
	}
	isCheck := models.NewConfig().GetConfigBool(platform, "isCheck", false)
	this.Success("操作成功", map[string]interface{}{
		"state": isCheck,
	})
}
