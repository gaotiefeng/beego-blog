package main

import (
	"beego/app/models"
	"beego/app/service/Filter"
	_ "beego/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init()  {

	host := beego.AppConfig.String("mysqlhost")
	prot := beego.AppConfig.String("mysqlprot")
	user := beego.AppConfig.String("mysqluser")
	pass := beego.AppConfig.String("mysqlpass")
	db := beego.AppConfig.String("mysqldb")

	logs.SetLogger(logs.AdapterConsole, `{"level":1,"color":true}`)
	logs.Debug(host)
	dataSrouce := user + ":" + pass + "@tcp("+ host +":"+ prot +")/" + db + "?charset=utf8"

	orm.RegisterDataBase("default","mysql",dataSrouce,60)

	orm.RegisterModel(new(models.User))


	beego.SetStaticPath("/resources","resources")
}


func main(){

	orm.Debug = true
	beego.InsertFilter("/admin/*",beego.BeforeRouter,Filter.AdminFilter)
	//打开session
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.Run()
}



