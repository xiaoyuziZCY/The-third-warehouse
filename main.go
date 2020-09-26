package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "xiaoyuzi/db_mysql"
	_ "xiaoyuzi/routers"
)

func main() {
	config := beego.AppConfig
	appname:=config.String("appname")
	httpport,err:=config.Int("httpport")
	if err !=nil {
		panic("错误5")
	}
	fmt.Println(httpport)
	fmt.Println("应用名称",appname)
	beego.Run()
}

