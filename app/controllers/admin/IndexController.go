package admin

type IndexController struct {
	LoginController
}

//
func (this *IndexController) Index()  {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/index.html"
}

