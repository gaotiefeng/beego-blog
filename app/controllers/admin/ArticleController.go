package admin

import (
	"beego/app/constants"
	"beego/app/service/Dao"
)

type ArticleController struct {
	LoginController
}


func (this *ArticleController) List()  {

	offset, _ := this.GetInt("offset",0)
	limit, _ := this.GetInt("limit",10)

	count,list,err := Dao.ArticleDaoList(offset,limit)

	if err != nil {
		this.ResponseError(constants.SERVERERROR,"查询失败",err)
	}
	//data := map[string]interface{}{"count":count,"items":list}

	this.ResponseSuccess("查询成功",list,count)
}
