package middleware

import (
	"fjsdxy-plus/helper"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

// 注册中间件
func init() {

	//TOKEN中间件
	var FilterToken = func(ctx *context.Context) {
		var apiVersion = "/" + beego.AppConfig.String("ApiVersion")
		var skipRouter = []string{
			apiVersion + "/wechat/login",
			apiVersion + "/qq/login",
			apiVersion + "/week/pull",
			apiVersion + "/week/get_today",
		}

		if !isSkip(ctx.Request.RequestURI, skipRouter) {
			token := ctx.Input.Header("Authorization")
			if helper.CheckToken(token) == "" {
				ctx.Output.Body([]byte("Not Authorization"))
			}
		}
	}

	beego.InsertFilter("/*", beego.BeforeRouter, FilterToken)
}

//判断是否跳过
func isSkip(url string, params []string) bool {
	path := strings.Split(url, "/")
	for _, router := range params {
		p := strings.Split(router, "/")
		if path[1] == p[1] && path[2] == p[2] && path[3] == p[3] {
			return true
		}
	}
	return false
}
