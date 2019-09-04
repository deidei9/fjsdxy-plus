package models

import (
	"strconv"
)

type Config struct {
	Id   int    `orm:"auto;pk;index"`
	Cate string `orm:"unique;size(30)"`
	Key  string `orm:"size(30)"`
	Val  string `orm:"size(255)"`
}

func NewConfig() *Config {
	return &Config{}
}

//获取配置
//@param            cate            配置分类
//@param            key             键
//@param			def				default，即默认值
//@return           val             值
func (this *Config) GetConfig(cate string, key string, def ...string) string {
	if err := DB().QueryTable(GetTable("config")).Filter("Cate", cate).Filter("Key", key).One(this); err == nil {
		return this.Val
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

//获取配置
//@param            cate            配置分类
//@param            key             键
//@param			def				default，即默认值
//@return           val             值
func (this *Config) GetConfigBool(cate string, key string, def ...bool) (val bool) {
	value := this.GetConfig(cate, key)

	if value == "true" || value == "1" {
		val = true
	} else {
		if len(def) > 0 {
			val = def[0]
		}
	}
	return
}

//获取配置
//@param            cate            配置分类
//@param            key             键
//@param			def				default，即默认值
//@return           val             值
func (this *Config) GetConfigInt64(cate string, key string, def ...int64) (val int64) {
	val, err := strconv.ParseInt(this.GetConfig(cate, key), 10, 64)
	if err != nil {
		if len(def) > 0 {
			val = def[0]
		}
	}
	return
}

//获取配置
//@param            cate            配置分类
//@param            key             键
//@param			def				default，即默认值
//@return           val             值
func (this *Config) GetConfigFloat64(cate string, key string, def ...float64) (val float64) {
	val, err := strconv.ParseFloat(this.GetConfig(cate, key), 64)
	if err != nil {
		if len(def) > 0 {
			val = def[0]
		}
	}
	return
}
