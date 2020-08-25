package routers

import (
	"beego/app/controllers"
	"beego/app/controllers/home"
	"github.com/astaxie/beego"
)

func init()  {

	//beego.Router("/",&controllers.BaseController{},"get:Welcome")
	beego.Router("/",&home.BlogController{},"get:Index")

	beego.Router("/404",&controllers.BaseController{},"get:Notfound")
}