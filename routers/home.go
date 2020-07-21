package routers

import (
	"beego/app/controllers/home"
	"github.com/astaxie/beego"
)

func init()  {
	beego.Router("/blog/index",&home.BlogController{},"get:Index")
	beego.Router("/blog/detail",&home.BlogController{},"get:Detail")
}
