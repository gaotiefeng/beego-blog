package Dao

import (
	"beego/app/models"
	"github.com/astaxie/beego/orm"
)

//文章信息
func ArticleDaoFind(id int)  (models.Article,error){
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err := o.Read(&article)

	return article,err
}
//文章列表
func ArticleDaoList(offset int,limit int) (int64,*[]models.Article,error) {
	article := new(models.Article)
	o := orm.NewOrm()
	articleModel := new([]models.Article)
	count,err := o.QueryTable(article).Limit(limit,offset).All(articleModel)

	return count,articleModel,err
}
//总条数
func ArticleDaoCount() (count int64) {

	article := new(models.Article)
	o := orm.NewOrm()
	count,_ = o.QueryTable(article).Count()

	return count
}