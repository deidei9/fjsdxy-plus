package routers

import (
	"github.com/astaxie/beego"
	"github.com/pig0224/fjsdxy-plus/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
