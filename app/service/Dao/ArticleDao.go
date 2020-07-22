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

	/*var list [] orm.Params
	count,err := o.QueryTable(article).Limit(limit,offset).Values(&list,"id","name","mobile","password","created_at")*/

	return count,articleModel,err
}
