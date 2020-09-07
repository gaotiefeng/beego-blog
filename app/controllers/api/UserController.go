package api

import (
	"beego/app/constants"
	"beego/app/controllers"
	"beego/app/service/Dao"
	"beego/app/service/Formatter"
	"beego/app/validation"
	"github.com/astaxie/beego"
)

type UserController struct {
	controllers.BaseController
}


func (this *UserController) Login()  {
	//TODO request verify
	mobile := this.GetString("mobile")
	password := this.GetString("password")

	err, user := Dao.UserMobile(mobile)
	json := controllers.Error(constants.SERVERERROR,"请求参数",mobile)
	if err != nil {
		json = controllers.Error(404,"请输入正确手机号",err)
	}else {
		if password == user.Password{
			json = controllers.Success(200,"登陆成功",user)
		}else {
			json = controllers.Error(constants.SERVERERROR,"密码错误",mobile)
		}
	}

	this.Data["json"] = map[string]interface{}{"code":200,"data":json}
	this.ServeJSON()
}

func (this *UserController) Register() {
	mobile := this.GetString("mobile")
	name := this.GetString("name")
	password := this.GetString("password","123456")
	//验证
	err := validation.RegisterValidation(name,mobile)
	if err != nil {
		beego.Info("验证器", err)
		this.ResponseError(constants.SERVERERROR, "验证器", err)
	}
	json := map[string]interface{}{}
	json["password"] = password

	this.ResponseSuccess("注册接口",json)

}

func (this *UserController) Find() {

	id, _ := this.GetInt("id")
	if id == 0 {
		this.ResponseError(constants.SERVERERROR,"id 不能为空",id)
	}
	err,user := Dao.UserFind(id)

	find := Formatter.UserFormatter(user)
	if err != nil {
		beego.Info("查询失败",err)
		this.ResponseError(constants.SERVERERROR,"查询失败",err)
	}else {
		this.ResponseSuccess("查询成功",find)
	}
}

func (this *UserController) List()  {
	offset, _ := this.GetInt("offset",0)
	limit, _ := this.GetInt("limit",10)

	count,list,err := Dao.UserDaoList(offset,limit)

	if err != nil {
		this.ResponseError(constants.SERVERERROR,"查询失败",err)
	}
	data := map[string]interface{}{"count":count,"items":list}

	this.ResponseSuccess("查询成功",data)
}


func (this *UserController) Update()  {

	id,_ := this.GetInt("id")
	name := this.GetString("name")
	if id == 0 {
		this.ResponseError(constants.SERVERERROR,"id 不能为空",id)
	}

	this.ResponseSuccess("更新成功",name)
}

func (this *UserController) Delete() {

	id,_ := this.GetInt("id")

	this.ResponseSuccess("删除成功",id)
}
