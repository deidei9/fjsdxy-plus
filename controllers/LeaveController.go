package controllers

import (
	"fjsdxy-plus/models"
	"github.com/pig0224/fjsdxy/xg/leave"
	"strconv"
)

type LeaveController struct {
	BaseController
}

// @Title 全部请假
// @Description 获取全部请假
// @Param Authorization header string true "token"
// @Success 200 {string} 获取结果
// @router /get_list [get]
func (this *LeaveController) GetLeave() {
	student := models.NewStudent()
	student.User = models.NewUser()
	student.User.Id = this.UserId
	if err := student.GetInfo(); err != nil {
		this.Error(10700, err.Error())
	}
	c, err := student.LoginXG()
	if err != nil {
		this.Error(10701, err.Error())
	}
	data, err := leave.Get(c)
	if err != nil {
		this.Error(10702, err.Error())
	}
	this.Success("获取成功", data)
}

// @Title 撤销假条
// @Description 撤销假条
// @Param Authorization header string true "token"
// @Param leaveId path string true "leaveId"
// @Success 200 {string} 获取结果
// @router /revoke/:leaveId [get]
func (this *LeaveController) Revokes() {
	leaveId := this.GetString(":leaveId")
	if leaveId == "" {
		this.Error(10600, "Id不能为空")
	}
	student := models.NewStudent()
	student.User = models.NewUser()
	student.User.Id = this.UserId
	if err := student.GetInfo(); err != nil {
		this.Error(10700, err.Error())
	}
	c, err := student.LoginXG()
	if err != nil {
		this.Error(10701, err.Error())
	}
	Id, _ := strconv.Atoi(leaveId)
	err = leave.Revoke(Id, c)
	if err != nil {
		this.Error(10702, err.Error())
	}
	this.Success("撤销成功")
}

// @Title 申请假条
// @Description 申请日常假条
// @Param Authorization header string true "token"
// @Param LeaveBeginDate formData string true "开始时间"
// @Param LeaveBeginTime formData string true "开始时刻"
// @Param LeaveEndDate formData string true "结束时间"
// @Param LeaveEndTime formData string true "结束时刻"
// @Param LeaveType formData string true "请假类型"
// @Param OutAddress formData string true "外出地点"
// @Param AreaWide formData string true "外出区域"
// @Param OutMoveTel formData string true "联系人电话"
// @Param Relation formData string true "联系人关系"
// @Param OutName formData string true "联系人姓名"
// @Param StuMoveTel formData string true "本人联系电话"
// @Param LeaveThing formData string true "请假事由"
// @Success 200 {string} json web token
// @router /apply [post]
func (this *LeaveController) Apply() {
	LeaveBeginDate := this.GetString("LeaveBeginDate")
	LeaveBeginTime := this.GetString("LeaveBeginTime")
	LeaveEndDate := this.GetString("LeaveEndDate")
	LeaveEndTime := this.GetString("LeaveEndTime")
	LeaveType := this.GetString("LeaveType")
	OutAddress := this.GetString("OutAddress")
	AreaWide := this.GetString("AreaWide")
	OutMoveTel := this.GetString("OutMoveTel")
	Relation := this.GetString("Relation")
	OutName := this.GetString("OutName")
	StuMoveTel := this.GetString("StuMoveTel")
	LeaveThing := this.GetString("LeaveThing")
	if LeaveBeginDate == "" || LeaveBeginTime == "" || LeaveEndDate == "" || LeaveEndTime == "" || LeaveType == "" || OutAddress == "" || AreaWide == "" || OutMoveTel == "" || Relation == "" || OutName == "" || LeaveThing == "" || StuMoveTel == "" {
		this.Error(10500, "参数不正确")
	}
	student := models.NewStudent()
	student.User = models.NewUser()
	student.User.Id = this.UserId
	if err := student.GetInfo(); err != nil {
		this.Error(10700, err.Error())
	}
	c, err := student.LoginXG()
	if err != nil {
		this.Error(10701, err.Error())
	}
	err = leave.Apply(leave.LeaveInfo{
		LeaveBeginDate: LeaveBeginDate,
		LeaveBeginTime: LeaveBeginTime,
		LeaveEndDate:   LeaveEndDate,
		LeaveEndTime:   LeaveEndTime,
		LeaveType:      LeaveType,
		OutAddress:     OutAddress,
		AreaWide:       AreaWide,
		OutMoveTel:     OutMoveTel,
		Relation:       Relation,
		OutName:        OutName,
		StuMoveTel:     StuMoveTel,
		LeaveThing:     LeaveThing,
	}, c)
	if err != nil {
		this.Error(10702, err.Error())
	}
	this.Success("提交成功")
}
