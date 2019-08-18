package middleware

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// 注册中间件
func init() {

	//授权中间件
	var FilterAuth = func(ctx *context.Context) {
		var apiVersion = "/" + beego.AppConfig.String("ApiVersion") + "/"
		var skipRouter = []string{
			apiVersion + "login",
		}
		if !isSkip(ctx.Request.RequestURI, skipRouter) {
			ctx.Output.Body([]byte("未授权操作"))
		}
	}

	beego.InsertFilter("/*", beego.BeforeRouter, FilterAuth)
}

//判断是否跳过
func isSkip(url string, params []string) bool {
	for _, router := range params {
		if url == router {
			return true
		}
	}
	return false
}
