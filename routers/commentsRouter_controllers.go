package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:CourseController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "GetCourse",
			Router:           `/get/:date`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:CourseController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "GetNextClass",
			Router:           `/get_next`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:ECardController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:ECardController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/get`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:ExamController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:ExamController"],
		beego.ControllerComments{
			Method:           "GetInfo",
			Router:           `/get`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:ExamController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:ExamController"],
		beego.ControllerComments{
			Method:           "PullExam",
			Router:           `/pull`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:LeaveController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:LeaveController"],
		beego.ControllerComments{
			Method:           "Apply",
			Router:           `/apply`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:LeaveController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:LeaveController"],
		beego.ControllerComments{
			Method:           "GetLeave",
			Router:           `/get_list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:LeaveController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:LeaveController"],
		beego.ControllerComments{
			Method:           "Revokes",
			Router:           `/revoke/:leaveId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:StudentController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:StudentController"],
		beego.ControllerComments{
			Method:           "Bind",
			Router:           `/bind`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:StudentController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:StudentController"],
		beego.ControllerComments{
			Method:           "ClassNotice",
			Router:           `/class_notice`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:StudentController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:StudentController"],
		beego.ControllerComments{
			Method:           "GetInfo",
			Router:           `/getInfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:StudentController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:StudentController"],
		beego.ControllerComments{
			Method:           "PullInfo",
			Router:           `/pull`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:StudentController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:StudentController"],
		beego.ControllerComments{
			Method:           "UnBind",
			Router:           `/unBind`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:UserController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetInfo",
			Router:           `/getInfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:WechatController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:WechatController"],
		beego.ControllerComments{
			Method:           "GetInfo",
			Router:           `/getInfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:WechatController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:WechatController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:WeekController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:WeekController"],
		beego.ControllerComments{
			Method:           "GetWeek",
			Router:           `/get/:today`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:WeekController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:WeekController"],
		beego.ControllerComments{
			Method:           "GetToday",
			Router:           `/get_today`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:WeekController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:WeekController"],
		beego.ControllerComments{
			Method:           "GetWeekInfo",
			Router:           `/get_week_info/:date`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fjsdxy-plus/controllers:WeekController"] = append(beego.GlobalControllerRouter["fjsdxy-plus/controllers:WeekController"],
		beego.ControllerComments{
			Method:           "Pull",
			Router:           `/pull/:term`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
