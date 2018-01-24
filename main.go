package main

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "server-hub/models"
	_ "server-hub/routers"
	// "github.com/go-redis/redis"
)

func init() {
	db,_ := beego.AppConfig.GetSection("mysql")
	
	// set default database
	conn := db["username"]+":"+db["password"]+"@tcp("+db["host"]+":"+db["port"]+")/"+db["database"]
	conn += "?charset="+db["charset"]
	orm.RegisterDataBase("default", "mysql",conn, 30)
	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}
