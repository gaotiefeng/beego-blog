package routers

import (
	"beego/app/controllers"
	"github.com/astaxie/beego"
)

func init()  {

	beego.Router("/",&controllers.BaseController{},"get:Welcome")
	beego.Router("/404",&controllers.BaseController{},"get:Notfound")
}