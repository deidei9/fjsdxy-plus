//定义一些常量和变量
package helper

import (
	"github.com/astaxie/beego"
)

const (
	//fjsdxy-plus Version
	VERSION = "v1.0"
	//Cache Config
	CACHE_CONF = `{"CachePath":"./cache/runtime","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}`

	DEFAULT_STATIC_EXT    = ".txt,.html,.ico,.jpeg,.png,.gif,.xml"
	DEFAULT_COOKIE_SECRET = "fjsdxy"

	WEIXIN = "wechat"
	QQ     = "qq"
)

var (
	//develop mode
	Debug = beego.AppConfig.String("runmode") == "dev"
)
