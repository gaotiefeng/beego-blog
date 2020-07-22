package Dao

import (
	"beego/app/models"
	"github.com/astaxie/beego/orm"
)

func ArticleDaoList(offset int,limit int) (int64,*[]models.Article,error) {
	article := new(models.Article)
	o := orm.NewOrm()
	articleModel := new([]models.Article)
	count,err := o.QueryTable(article).Limit(limit,offset).All(articleModel)

	return count,articleModel,err
}
