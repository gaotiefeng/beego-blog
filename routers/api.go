package routers

import (
	"beego/app/controllers/api"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/api/user/login",&api.UserController{},"get:Login")
	beego.Router("/api/user/register",&api.UserController{},"get:Register")
	beego.Router("/api/user/find",&api.UserController{},"get:Find")
	beego.Router("/api/user/list",&api.UserController{},"get:List")
	beego.Router("/api/user/update",&api.UserController{},"get:Update")
	beego.Router("/api/user/delete",&api.UserController{},"get:Delete")


}