package helper

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var redis cache.Cache

func InitCache() {
	var err error
	host := beego.AppConfig.DefaultString("redis::host", "localhost")
	port := beego.AppConfig.DefaultString("redis::port", "6379")
	password := beego.AppConfig.DefaultString("redis::password", "")
	config := make(map[string]interface{})
	config["key"] = "fjsdxy"
	config["conn"] = host + ":" + port
	config["dbNum"] = "0"
	config["password"] = password
	c, _ := json.Marshal(config)
	redis, err = cache.NewCache("redis", string(c))
	if err != nil {
		Logger.Error("cache初始化失败")
	}
}
