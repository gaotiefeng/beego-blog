package Admin

import (
	"beego/app/models"
	"beego/app/service/Dao"
	"time"
)

//文章更新或添加
func ArticleSave(id int,name string,content string,image string,classId int) (num int64,err error)  {

	//struct object
	article := models.Article{}
	//对结构体对象赋值
	article.Id = id
	article.Content = content
	article.Name = name
	article.Image = image
	article.ClassId = classId
	article.CreatedAt = time.Now()
	if id == 0 {
		num = 0
		err = Dao.ArticleDaoInsert(article)
	}else {
		num,err = Dao.ArticleDaoUpdate(article)
	}
	return num,err
}