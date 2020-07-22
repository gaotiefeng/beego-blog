package admin

import (
	"beego/app/constants"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) ResponseError (code int,message string, data interface{}) () {

	var json interface{}
	json = map[string]interface{}{"code":code,"msg":message,"data":data}
	this.Data["json"] = json
	this.ServeJSON()
}

func (this *LoginController) ResponseSuccess (message string, data interface{},count int64) () {

	var json interface{}
	json = map[string]interface{}{"code":constants.SUCCESS,"msg":message,"count":count,"data":data}
	this.Data["json"] = json
	this.ServeJSON()
}
