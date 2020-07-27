package routers

import (
	"beego/app/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin/login/index", &admin.LoginController{},"get:Index")
	beego.Router("/admin/login/login", &admin.LoginController{},"post:Login")


	beego.Router("/admin/index/index",&admin.IndexController{},"get:Index")
	beego.Router("/admin/index/test",&admin.IndexController{},"get:Test")
	beego.Router("/admin/article/list",&admin.ArticleController{},"get:List")

}