package models

import "github.com/astaxie/beego/orm"

type Config struct {
	Id int	`orm:"column(id);auto" description:"id" json:"id"`
	ConfigKey string	`orm:"column(config_key)" description:"config_key" json:"config_key"`
	ConfigValue string	`orm:"column(config_value)" description:"config_value" json:"config_value"`
}

func (u *Config) TableName() string {
	// db table name
	return "config"
}

func init() {
	orm.RegisterModel(new(Config))
}