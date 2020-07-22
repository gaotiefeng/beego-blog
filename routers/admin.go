package routers

import (
	"beego/app/controllers"
	"beego/app/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/",&controllers.BaseController{},"get:Welcome")
	beego.Router("/admin/index/index",&admin.IndexController{},"get:Index")
}