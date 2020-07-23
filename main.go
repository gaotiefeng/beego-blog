package main

import (
	"beego/app/models"
	"beego/conf"
	_ "beego/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init()  {

	dataSrouce := conf.MYSQL_USERNAME + ":" + conf.MYSQL_PASSWORD + "@tcp("+ conf.MYSQL_HOST+":"+ conf.MYSQL_PORT+")/" + conf.MYSQL_DATABASE + "?charset=utf8"

	orm.RegisterDataBase("default","mysql",dataSrouce,60)

	orm.RegisterModel(new(models.User))


	beego.SetStaticPath("/resources","resources")
}


func main(){

	logs.SetLogger(logs.AdapterConsole, `{"level":1,"color":true}`)
	/*logs.Async(1e3)
	logs.SetLogger(logs.AdapterFile, `{"filename":"runtime/logs/beego.log"}`)*/

	beego.Run()
}
