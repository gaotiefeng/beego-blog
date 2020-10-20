package Admin

import (
	"beego/app/models"
	"beego/app/service/Dao"
	"time"
)

func ClassSave(id int,name string,sort int,pid int) (num int64,err error) {
	//struct object
	class := models.Class{}
	//对结构体对象赋值
	class.ClassId = id
	class.Name = name
	class.Sort = sort
	class.Pid = pid
	class.CreatedAt = time.Now()
	if id == 0 {
		num = 0
		err = Dao.ClassDaoInsert(class)
	}else {
		num,err = Dao.ClassDaoUpdate(class)
	}
	return num,err
}
