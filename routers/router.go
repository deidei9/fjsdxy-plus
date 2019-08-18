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
	var apiVersion = beego.AppConfig.String("ApiVersion")
	ns := beego.NewNamespace(apiVersion,
		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		//beego.NSNamespace("/user",
		//	beego.NSInclude(
		//		&controllers.UserController{},
		//	),
		//),
		beego.NSNamespace("/course",
			beego.NSInclude(
				&controllers.CourseController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
