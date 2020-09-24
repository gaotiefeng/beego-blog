package admin

type IndexController struct {
	LoginController
}

//
func (this *IndexController) Index()  {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/index.html"
}

func (this *IndexController) Test()  {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/article/list.html"
}