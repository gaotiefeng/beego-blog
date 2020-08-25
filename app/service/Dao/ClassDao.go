package Dao

import (
	"beego/app/models"
	"github.com/astaxie/beego/orm"
)

/**
查询一级分类
 */
func ClassFindParent(pid int) (*[]models.Class,error) {
	classModel := new([]models.Class)

	o := orm.NewOrm()
	qs := o.QueryTable("class")
	_,err  := qs.Filter("pid",pid).OrderBy("sort desc").All(classModel)

	return classModel,err
}
/**
递归查询分类
 */
func GetClass(pid int) []*models.ClassTreeList {
	o := orm.NewOrm()
	var class []models.Class
	_,_ = o.QueryTable("class").Filter("pid", pid).OrderBy("sort").All(&class)
	treeList := []*models.ClassTreeList{}
	for _, v := range class{
		child := GetClass(v.ClassId)
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

