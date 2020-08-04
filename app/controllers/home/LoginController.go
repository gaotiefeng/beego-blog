package home

import (
	"beego/client"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Index()  {
	client.Post()
	this.ServeJSON()
}

