##### beego

[![beego](https://img.shields.io/badge/go-beego-blue)](https://github.com/astaxie/beego)

###### 启动

```
cp config/Database.go.dist config/Database.go
go run main.go
```

###### router

```
    /user/register
    /user/find
    /user/list
    /user/update
    /user/delete
```

###### 博客列表
```
    /blog/index
    /blog/detail
```

##### 后台 ui layui
```
    /admin/index/index
```

[beego 文档地址](https://beego.me/docs/install/bee.md)

```安装 bee 工具
go get github.com/beego/bee
#热更新controller 会自动更新
bee run
```