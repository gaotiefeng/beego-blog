package controllers

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

type ToolsController struct {
	BaseController
}

func init() {
	//跨域设置
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
}
/*
加密工具
 */
func (this *ToolsController)Encryption()  {

	Type := this.GetString("type","md5")
	str := this.GetString("string","123456")

	var data map[string]interface{}
	data = make(map[string]interface{})
	data["type"] = Type
	if Type == "md5" {
		data["encryp"] = md5V(str)
	}else if Type == "sha1" {
		data["encryp"] = sha1V(str)
	}

	this.ResponseSuccess("加密",data)
}

func md5V(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func sha1V(str string) string {
	h := sha1.New() // md5加密类似md5.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(str))
	//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来对现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	//SHA1 值经常以 16 进制输出，使用%x 来将散列结果格式化为 16 进制字符串。
	return fmt.Sprintf("%x", bs)
}

