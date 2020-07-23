package Dao

import (
	"beego/app/models"
	"github.com/astaxie/beego/orm"
)


func ArticleDaoFind(id int)  (models.Article,error){
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err := o.Read(&article)

	return article,err
}

func ArticleDaoList(offset int,limit int) (int64,*[]models.Article,error) {
	article := new(models.Article)
	o := orm.NewOrm()
	articleModel := new([]models.Article)
	count,err := o.QueryTable(article).Limit(limit,offset).All(articleModel)

	return count,articleModel,err
}
