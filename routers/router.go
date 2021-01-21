package routers

import (
	"beego/app/controllers"
	"github.com/astaxie/beego"
)

func init()  {

	//beego.Router("/",&controllers.BaseController{},"get:Welcome")

	beego.Router("/404",&controllers.BaseController{},"get:Notfound")
	beego.Router("/500",&controllers.BaseController{},"get:Error")
	beego.Router("/505",&controllers.BaseController{},"get:Defend")

}