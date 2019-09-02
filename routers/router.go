// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"fjsdxy-plus/controllers"
	_ "fjsdxy-plus/routers/middleware"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/wechat",
			beego.NSInclude(
				&controllers.WechatController{},
			),
		),
		beego.NSNamespace("/qq",
			beego.NSInclude(
				&controllers.QQController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/student",
			beego.NSInclude(
				&controllers.StudentController{},
			),
		),
		beego.NSNamespace("/week",
			beego.NSInclude(
				&controllers.WeekController{},
			),
		),
		beego.NSNamespace("/exam",
			beego.NSInclude(
				&controllers.ExamController{},
			),
		),
		beego.NSNamespace("/course",
			beego.NSInclude(
				&controllers.CourseController{},
			),
		),
		beego.NSNamespace("/ecard",
			beego.NSInclude(
				&controllers.ECardController{},
			),
		),
		beego.NSNamespace("/leave",
			beego.NSInclude(
				&controllers.LeaveController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
