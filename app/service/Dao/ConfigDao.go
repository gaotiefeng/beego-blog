package Dao

import (
	"beego/app/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

func FindByConfigKey(key string) (config models.Config)  {

	config = models.Config{}
	o := orm.NewOrm()
	qs := o.QueryTable("config")
	err  := qs.Filter("config_key",key).One(&config)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(config.Id, config.ConfigValue)
	}

	return config
}