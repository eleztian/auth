package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"log"
	"time"
)

type TestController struct {
	beego.Controller
}

func (t *TestController)Test()  {
	redies, err := cache.NewCache("redis", `{"conn":"111.231.215.178:6379", "key":""}`)
	if err != nil {
		log.Fatal(err)
	}
	redies.Put("user", "tab", time.Duration(1*time.Second))
}
