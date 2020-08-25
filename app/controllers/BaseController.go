package controllers

import (
	//"fmt"
	"beego/app/constants"
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func TypeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func (c *BaseController) Welcome() {

	//c.Data["json"] = map[string]interface{}{"code": 200, "message": "hello word"}
	//c.ServeJSON()
	c.Data["content"] = "welcome to blog"
	c.Data["url"] = "/blog/index"
	c.TplName = "welcome.html"
}

func (c *BaseController) Notfound() {

	c.TplName = "404.tpl"
}

func Success(code int,message string,data interface{}) (json interface{}){

	json = map[string]interface{}{"code":code,"msg":message,"data":data}

	return json
}

func Error(code int,message string, data interface{}) (json interface{}) {

	json = map[string]interface{}{"code":code,"msg":message,"data":data}
	return json
}

func (this *BaseController) ResponseError (code int,message string, data interface{}) () {

	var json interface{}
	json = map[string]interface{}{"code":code,"msg":message,"data":data}
	this.Data["json"] = json
	this.ServeJSON()
}

func (this *BaseController) ResponseSuccess (message string, data interface{}) () {

	var json interface{}
	json = map[string]interface{}{"code":constants.SUCCESS,"msg":message,"data":data}
	this.Data["json"] = json
	this.ServeJSON()
}