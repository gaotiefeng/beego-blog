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
func ArticleDaoList(offset int,limit int,classId int) (int64,*[]models.Article,error) {
	article := new(models.Article)
	o := orm.NewOrm()
	articleModel := new([]models.Article)

	qs := o.QueryTable(article)
	if classId !=0 {
		qs.Filter("class_id",classId)
	}
	count,err := qs.Limit(limit,offset).All(articleModel)

	return count,articleModel,err
}
//总条数
func ArticleDaoCount() (count int64) {

	article := new(models.Article)
	o := orm.NewOrm()
	count,_ = o.QueryTable(article).Count()

	return count
}
//添加
func ArticleDaoInsert(article models.Article) (err error) {
	//orm object
	o := orm.NewOrm()
	//insert
	_,err = o.Insert(&article)
	return err
}
//更新
func ArticleDaoUpdate(article models.Article) (num int64,err error) {
	o := orm.NewOrm()

	num,err = o.Update(&article, "Name","Content","Image","ClassId","ChildId")

	return num,err
}
//删除
func ArticleDaoDelete(id int) (num int64, err error) {
	//orm object
	o := orm.NewOrm()
	//struct object
	article := models.Article{}
	article.Id = id
	num, err = o.Delete(&article)

	return num, err
}
//文章列表
func ArticleDaoDataAll() (int64,*[]models.Article,error) {
	article := new(models.Article)
	o := orm.NewOrm()
	articleModel := new([]models.Article)
	count,err := o.QueryTable(article).All(articleModel)

	return count,articleModel,err
}