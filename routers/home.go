package routers

import (
	"beego/app/controllers"
	"beego/app/controllers/home"
	"github.com/astaxie/beego"
)

func init()  {
	beego.Router("/",&home.BlogController{},"get:Index")
	beego.Router("/index",&home.IndexController{},"get:Index")

	beego.Router("/login/index",&home.LoginController{},"get:Index")
	beego.Router("/article/detail",&home.ArticleController{},"get:Detail")
	beego.Router("/article/list",&home.ArticleController{},"post:List")
	beego.Router("/tools/encryption",&controllers.ToolsController{},"post:Encryption")
}
