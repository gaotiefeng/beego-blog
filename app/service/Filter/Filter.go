package Filter

import (
	"github.com/astaxie/beego/context"
	"strings"
)

//admin登录过滤器
var AdminFilter = func(ctx *context.Context){
	_, ok := ctx.Input.Session("adminId").(int)
	ok2 := strings.Contains(ctx.Request.RequestURI, "/admin/login")
	if !ok && !ok2{
		ctx.Redirect(302, "/admin/login")
	}
}
