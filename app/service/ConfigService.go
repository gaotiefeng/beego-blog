package service

import (
	"beego/app/models"
	"beego/app/service/Dao"
	"github.com/astaxie/beego"
)

/**
获取配置值
 */
func GetByConfigKey(key string) (value string) {
	if key == "" {
		beego.Error("get Byconfigkey is empty")
		return ""
	}
	config :=  Dao.FindByConfigKey(key)
	return config.ConfigValue
}

/**
获取所有配置值
*/
func GetByConfig() (int64,*[]models.Config,error) {

	count,configModel,err :=  Dao.FindConfigAll()
	return count,configModel,err
}