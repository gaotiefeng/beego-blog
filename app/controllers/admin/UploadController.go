package admin

import (
	"beego/app/constants"
	"beego/app/service/Qiniu"
)

type UploadController struct {
	LoginController
}

func (this *UploadController) QiniuToken(){
	token := Qiniu.GetToken()
	data := map[string]interface{}{}
	data["token"] = token
	this.ResponseSuccess(constants.SUCCESS,"1",data)
}

func (this *UploadController) FileImg() {

}