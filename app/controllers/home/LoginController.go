package home

import (
	"beego/tool/http"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Index()  {
	http.Post()
	this.ServeJSON()
}

