package Dao

import (
	"beego/app/models"
	"github.com/astaxie/beego/orm"
)
//分类信息
func ClassFirst(id int)  (models.Class,error){
	o := orm.NewOrm()
	class := models.Class{ClassId: id}
	err := o.Read(&class)

	return class,err
}
/**
0为一级分类
查寻pid的子集
 */
func ClassFindParent(pid int) (*[]models.Class,error) {
	classModel := new([]models.Class)

	o := orm.NewOrm()
	qs := o.QueryTable("class")
	_,err  := qs.Filter("pid",pid).OrderBy("-sort").All(classModel)

	return classModel,err
}
/**
递归查询分类
 */
func ClassGetAll(pid int) []*models.ClassTreeList {
	o := orm.NewOrm()
	var class []models.Class
	_,_ = o.QueryTable("class").Filter("pid", pid).OrderBy("-sort").All(&class)
	treeList := []*models.ClassTreeList{}
	for _, v := range class{
		child := ClassGetAll(v.ClassId)
		node := &models.ClassTreeList{
			ClassId:v.ClassId,
			Name:v.Name,
			Sort:v.Sort,
			Pid:v.Pid,
		}
		node.Children = child
		treeList = append(treeList, node)
	}
	return treeList
}
/**
查寻所有分类
*/
func ClassFindAll() (*[]models.Class,error) {
	classModel := new([]models.Class)

	o := orm.NewOrm()
	qs := o.QueryTable("class")
	_,err  := qs.OrderBy("-sort").All(classModel)

	return classModel,err
}
/**
查寻所有分类
*/
func ClassList(offset int,limit int) (*[]models.Class,error) {
	classModel := new([]models.Class)

	o := orm.NewOrm()
	qs := o.QueryTable("class")
	_,err  := qs.Limit(limit).Offset(offset).OrderBy("-sort").All(classModel)

	return classModel,err
}
//总条数
func ClassDaoCount() (count int64) {

	class := new(models.Class)
	o := orm.NewOrm()
	count,_ = o.QueryTable(class).Count()

	return count
}
//添加
func ClassDaoInsert(class models.Class) (err error) {
	//orm object
	o := orm.NewOrm()
	//insert
	_,err = o.Insert(&class)
	return err
}
//更新
func ClassDaoUpdate(class models.Class) (num int64,err error) {
	o := orm.NewOrm()

	num,err = o.Update(&class, "Name","Pid","Sort")

	return num,err
}

func ClassDaoDel(id int)  (num int64,err error){
	class := models.Class{ClassId: id}
	o := orm.NewOrm()
	num,err = o.Delete(&class)
	return num,err
}