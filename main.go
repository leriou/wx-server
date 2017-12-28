package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"server-hub/models"
	_ "server-hub/routers"
	"github.com/go-redis/redis"
)

func init() {
	// set default database
	db_host := beego.AppConfig.String("jdbc.host")
	db_uname :=  beego.AppConfig.String("jdbc.username")
	db_pwd :=  beego.AppConfig.String("jdbc.password")
	db_database :=  beego.AppConfig.String("jdbc.database")
	db_charset :=  beego.AppConfig.String("jdbc.charset")

	conn := db_uname+":"+db_pwd+"@tcp("+db_host+":3306)/"+db_database+"?charset="+db_charset
	orm.RegisterDataBase("default", "mysql",conn, 30)
	// create table
	//orm.RunSyncdb("default", false, true)

	client := redis.NewClient(&redis.Options{
				Addr:     "localhost:6379",
				Password: "", // no password set
				DB:       0,  // use default DB
			})
}

func main() {
	beego.Run()
}
