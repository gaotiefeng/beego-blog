package routers

import (
	"beego/app/controllers/home"
	"github.com/astaxie/beego"
)

func init()  {
	beego.Router("/login/index",&home.LoginController{},"get:Index")

	beego.Router("/blog/detail",&home.BlogController{},"get:Detail")

	beego.Router("/article/list",&home.ArticleController{},"post:List")
}
