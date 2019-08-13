package controllers

import (
	"github.com/astaxie/beego"
	"github.com/pig0224/fjsdxy"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
