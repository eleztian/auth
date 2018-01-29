package main

import (
	_ "etecity/auth/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"etecity/auth/models"
)

func init() {
	logs.Register("mongo", func() logs.Logger {
		return &models.MongoDBLogger{}
	})
	beego.SetLogger("mongo", `{"db_name":"llog"}`)
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "111.231.215.178:6379,100,"
	//if beego.BConfig.RunMode == "dev" {
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	beego.Debug("sss")
	beego.Run()

}
