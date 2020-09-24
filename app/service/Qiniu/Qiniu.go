package Qiniu

import (
	"beego/conf"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

func GetToken() (token string){
	bucket := conf.Bucket
	putPoliy := storage.PutPolicy{
		Scope: bucket,
	}
	accessKey := conf.AccessKey
	secretKey := conf.SecretKey
	mac := qbox.NewMac(accessKey,secretKey)
	upToken := putPoliy.UploadToken(mac)
	return upToken
}
