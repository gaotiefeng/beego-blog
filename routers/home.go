package routers

import (
	"beego/app/controllers"
	"beego/app/controllers/home"
	"github.com/astaxie/beego"
)

func init()  {
	beego.Router("/",&home.IndexController{},"get:Index")//首页
	beego.Router("/index",&home.ArticleController{},"get:Index")//文章列表
	beego.Router("/article/detail",&home.ArticleController{},"get:Detail") //文章详情

	beego.Router("/login/index",&home.LoginController{},"get:Index")
	beego.Router("/article/list",&home.ArticleController{},"post:List")
	beego.Router("/tools/encryption",&controllers.ToolsController{},"post:Encryption")
}
