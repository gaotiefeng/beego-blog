package service

import (
	"beego/app/service/Dao"
	"github.com/astaxie/beego"
)

/**
获取配置值
 */
func GetByConfigKey(key string) (value string) {
	if key == "" {
		beego.Error("getByconfigkey is empty")
		return ""
	}
	config :=  Dao.FindByConfigKey(key)
	return config.ConfigValue
}