package admin

import (
	"beego/app/constants"
	"beego/app/service/Qiniu"
	"beego/tool/file"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type UploadController struct {
	LoginController
}
type Sizer interface {
	Size() int64
}
const (
	LOCAL_FILE_DIR    = "static/uploads/file"
	EDIT_FILE_DIR= "static/uploads/file_edit"
	MIN_FILE_SIZE     = 1       // bytes
	MAX_FILE_SIZE     = 5000000 // bytes
	IMAGE_TYPES       = "(jpg|gif|p?jpeg|(x-)?png)"
	ACCEPT_FILE_TYPES = IMAGE_TYPES
	EXPIRATION_TIME   = 300 // seconds
	THUMBNAIL_PARAM   = "=s80"
)
type FileInfo struct {
	Url          string `json:"url,omitempty"`
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Size         int64  `json:"size"`
	Error        string `json:"error,omitempty"`
	DeleteUrl    string `json:"deleteUrl,omitempty"`
	DeleteType   string `json:"deleteType,omitempty"`
}
var (
	imageTypes      = regexp.MustCompile(IMAGE_TYPES)
	acceptFileTypes = regexp.MustCompile(ACCEPT_FILE_TYPES)
)
//获得七牛token
func (this *UploadController) QiniuToken(){
	token := Qiniu.GetToken()
	data := map[string]interface{}{}
	data["token"] = token
	this.ResponseSuccess(constants.SUCCESS,"1",data)
}
func (fi *FileInfo) ValidateType() (valid bool) {
	if acceptFileTypes.MatchString(fi.Type) {
		return true
	}
	fi.Error = "Filetype not allowed"
	return false
}
//图片上传//本地上传
func (this *UploadController) UploadFileImg() {
	f, h, err := this.GetFile("file")
	t := time.Now()
	path := LOCAL_FILE_DIR+t.Format("2006-01-02")

	var json interface{}
	defer f.Close()
	if err != nil {
		fmt.Println("getfile err ", err)
		json = map[string]interface{}{"code":500,"msg":"文件错误","state":"ERROR"}
		this.Success(json)
		return
	} else {
		var Url string
		ext := file.Ext(h.Filename)
		fi := &FileInfo{
			Name: h.Filename,
			Type: ext,
		}
		if !fi.ValidateType() {
			json = map[string]interface{}{"code":500,"msg":"invalid file type","state":"ERROR"}
			this.Success(json)
			return
		}
		var fileSize int64
		if sizeInterface, ok := f.(Sizer); ok {
			fileSize = sizeInterface.Size()
			fmt.Println(fileSize)
		}
		fileExt := strings.TrimLeft(ext, ".")
		fileSaveName := fmt.Sprintf("%s_%d%s", fileExt, time.Now().Unix(), ext)
		imgPath := fmt.Sprintf("%s/%s", path, fileSaveName)

		file.InsureDir(path)

		this.SaveToFile("file", imgPath) // 保存位置在 static/upload,没有文件夹要先创建
		if err == nil {
			Url = "/" + imgPath
		}
		json = map[string]interface{}{"code":200,"msg":"上传成功","src":Url,"state":"SUCCESS","title":h.Filename,"size":h.Size,"url":Url,"type":ext}

		this.Success(json)
	}
}
//编辑器图片上传
//本地上传
func (this *UploadController) EditUploadFileImg() {

	var json interface{}
	f, h, err := this.GetFile("edit_file")
	if err != nil {
		json = map[string]interface{}{"code":200,"msg":"配置数据|错误","imageUrlPrefix":constants.HOST,"state":"ERROR","imageActionName":"edit-image","imagePath":EDIT_FILE_DIR,"imageFieldName":"edit_file","imageMaxSize":50048,"imageAllowFiles":[7]string{".xls",".pdf",".png", ".jpg", ".jpeg", ".gif", ".bmp"}}
		this.Success(json)
		return
	}
	t := time.Now()
	path := EDIT_FILE_DIR+t.Format("2006-01-02")

	defer f.Close()
	if err != nil {
		fmt.Println("getfile err ", err)
		json = map[string]interface{}{"code":500,"msg":"文件错误","state":"ERROR","imageActionName":"edit-image","imagePath":EDIT_FILE_DIR}
		this.Success(json)
		return
	} else {
		fmt.Println("getfile err 1", err)
		var Url string
		ext := file.Ext(h.Filename)
		fi := &FileInfo{
			Name: h.Filename,
			Type: ext,
		}
		if !fi.ValidateType() {
			json = map[string]interface{}{"code":500,"msg":"invalid file type","state":"ERROR"}
			this.Success(json)
			return
		}
		var fileSize int64
		if sizeInterface, ok := f.(Sizer); ok {
			fileSize = sizeInterface.Size()
			fmt.Println(fileSize)
		}
		fileExt := strings.TrimLeft(ext, ".")
		fileSaveName := fmt.Sprintf("%s_%d%s", fileExt, time.Now().Unix(), ext)
		imgPath := fmt.Sprintf("%s/%s", path, fileSaveName)

		file.InsureDir(path)

		this.SaveToFile("edit_file", imgPath) // 保存位置在 static/upload,没有文件夹要先创建
		if err == nil {
			Url = "/" + imgPath
		}
		json = map[string]interface{}{"code":200,"msg":"上传成功","src":Url,"state":"SUCCESS","title":h.Filename,"size":h.Size,"url":Url,"type":ext}

		this.Success(json)
	}
}