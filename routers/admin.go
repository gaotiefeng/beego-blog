package routers

import (
	"beego/app/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	//登录
	beego.Router("/admin/login", &admin.LoginController{},"get:Index")
	beego.Router("/admin/login/login", &admin.LoginController{},"post:Login")
	//面板页
	beego.Router("/admin/index/index",&admin.IndexController{},"get:Index")
	//分类
	beego.Router("/admin/class/child",&admin.ClassController{},"post:GetChild")
	beego.Router("/admin/class-list",&admin.ClassController{},"get:List")

	beego.Router("/admin/class-index",&admin.ClassController{},"get:Index")
	beego.Router("/admin/class-add",&admin.ClassController{},"get:Add")

	//文章
	beego.Router("/admin/article-index",&admin.ArticleController{},"get:Index")
	beego.Router("/admin/article/save",&admin.ArticleController{},"post:Save")
	beego.Router("/admin/article/list",&admin.ArticleController{},"get:List")
	beego.Router("/admin/article/add",&admin.ArticleController{},"get:Add")
	beego.Router("/admin/article/del",&admin.ArticleController{},"post:Del")
	beego.Router("/admin/article/data-all",&admin.ArticleController{},"post:DataAll")

	//七牛
	beego.Router("/admin/upload/get-token",&admin.UploadController{},"get:QiniuToken")
	//图片上传
	beego.Router("/admin/upload/image",&admin.UploadController{},"post:UploadFileImg")
	beego.Router("/admin/upload/edit-image",&admin.UploadController{},"*:EditUploadFileImg")

}