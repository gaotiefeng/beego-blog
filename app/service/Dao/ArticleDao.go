package Dao

import (
	"beego/app/models"
	"github.com/astaxie/beego"
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
func ArticleDaoList(offset int,limit int,classId int,childId int) (int64,*[]models.Article,error) {
	article := new(models.Article)
	o := orm.NewOrm()
	articleModel := new([]models.Article)

	qs := o.QueryTable(article)
	beego.Info(classId)
	if classId != 0 {
		qs = qs.Filter("class_id",classId)
	}
	if childId != 0 {
		qs = qs.Filter("child_id",childId)
	}
	count,err := qs.Limit(limit,offset).OrderBy("-created_at").All(articleModel)

	return count,articleModel,err
}
//最近热文
func ArticleDaoHotList(offset int,limit int,classId int,childId int) (int64,*[]models.Article,error) {
	article := new(models.Article)
	o := orm.NewOrm()
	articleModel := new([]models.Article)

	qs := o.QueryTable(article)
	beego.Info(classId)
	if classId != 0 {
		qs = qs.Filter("class_id",classId)
	}
	if childId != 0 {
		qs = qs.Filter("child_id",childId)
	}
	count,err := qs.Limit(limit,offset).OrderBy("-click").All(articleModel)

	return count,articleModel,err
}
//总条数
func ArticleDaoCount(classId int,childId int) (count int64) {

	article := new(models.Article)
	o := orm.NewOrm()

	qs := o.QueryTable(article)
	if classId != 0 {
		qs = qs.Filter("class_id",classId)
	}
	if childId != 0 {
		qs = qs.Filter("child_id",childId)
	}
	count,_ = qs.Count()
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
//更新点击量
func ArticleDaoUpdateClick(article models.Article) (num int64,err error) {
	o := orm.NewOrm()

	num,err = o.Update(&article, "Click")

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
//上一篇
func ArticleDaoPre(id int) (*models.Article,error)  {
	o := orm.NewOrm()
	article := new(models.Article)
	err := o.QueryTable(article).Filter("id__lt",id).OrderBy("-id").One(article)

	return article,err
}
//下一篇
func ArticleDaoNext(id int) (*models.Article,error)  {
	o := orm.NewOrm()
	article := new(models.Article)
	err := o.QueryTable(article).Filter("id__gt",id).OrderBy("id").One(article)

	return article,err
}