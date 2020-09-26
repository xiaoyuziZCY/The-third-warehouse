package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"xiaoyuzi/db_mysql"
	"xiaoyuzi/models"
)

type Registercontrooler struct {
beego.Controller
}
func(r *Registercontrooler)Post(){
fmt.Println(r ==nil)
fmt.Println(r.Ctx ==nil)
fmt.Println(r.Ctx.Request ==nil)
bodyBytes,err :=ioutil.ReadAll(r.Ctx.Request.Body)
if err !=nil{
	r.Ctx.WriteString("数据错误1")
	return
}
	var user models.User
	err = json.Unmarshal(bodyBytes,&user)
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("数据错误2")
		return
	}
	id, err := db_mysql.InsertUser(user)
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("用户保存失败.")
		return
	}
	fmt.Println(id)
	res:=models.Responseresult{
		Code:0,
		Message:"运行成功",
		Data:nil,
	}
	r.Data["json"] = &res
	r.ServeJSON()
	fmt.Println("姓名:",)
	fmt.Println("年龄：",)
	fmt.Println("性别:",)
}
