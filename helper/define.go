//定义一些常量和变量
package helper

import (
	"github.com/astaxie/beego"
	"sync"
)

type ConfigCate string

const (
	//fjsdxy-plus Version
	VERSION = "v1.0"
	//Cache Config
	CACHE_CONF = `{"CachePath":"./cache/runtime","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}`

	DEFAULT_STATIC_EXT    = ".txt,.html,.ico,.jpeg,.png,.gif,.xml"
	DEFAULT_COOKIE_SECRET = "fjsdxy"
)

var (
	//develop mode
	Debug = beego.AppConfig.String("runmode") == "dev"
	//配置文件的全局map
	ConfigMap sync.Map
)
