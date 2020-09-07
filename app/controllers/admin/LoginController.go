package admin

import (
	"beego/app/constants"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}


func (this *LoginController) Index()  {
	this.TplName = "admin/login.html"
}

func (this *LoginController) Login()  {
	userName := this.Input().Get("user_name")
	password := this.GetString("password")
	var count int64
	var json = make(map[string]interface{})
	json["userName"] = userName
	json["password"] = password
	this.ResponseSuccess(constants.SUCCESS,"登录成功",json,count)
}

func (this *LoginController) ResponseError (code int,message string, data interface{}) () {

	var json interface{}
	json = map[string]interface{}{"code":code,"msg":message,"data":data}
	this.Data["json"] = json
	this.ServeJSON()
}

func (this *LoginController) ResponseSuccess (code int,message string, data interface{},count int64) () {

	var json interface{}
	json = map[string]interface{}{"code":code,"msg":message,"count":count,"data":data}
	this.Data["json"] = json
	this.ServeJSON()
}
