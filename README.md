# wx-server

这是一个用于个人演示和学习的项目,基于Beego框架进行开发,只提供API服务

## 安装

使用go官方推荐的`dep`来做包管理

安装dep:`brew innstall dep`

安装依赖:`dep ensure`

## 运行

增加配置

1. 在 server-hub 目录里面创建目录 conf, 新增配置文件 app.conf

app.conf内容

        appname = server-hub
        httpport = 8087
        runmode = dev

        [mysql]
        host = localhost
        username = root
        password = password
        database = wx_mp
        port = 3306
        charset = utf8

运行`go run main.go` or `bee run` 

## 开发计划

目前的开发计划有

1. 为微信公众号提供服务
2. 用户系统
3. 权限系统
4. 后台任务
5. 机器学习功能

## 微信公众号服务

功能进度:

1. 接入微信公众号(已完成)
2. 消息接收和回复框架(已完成)
3. 根据用户消息进行功能性回复(已完成)
4. 其他高级功能(因为需要300RMB的认证费用,所以暂停了)
5. 用户系统(开发中)
6. 权限系统

## 其他

1. 项目代码结构优化(计划中)
2. 数据库和消息队列服务解耦(计划中)