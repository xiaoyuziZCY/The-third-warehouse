package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"xiaoyuzi/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//获取name，age，sax
	//数据对比
	fmt.Println("======================================")
	//根据对比结果进行判断处理
	name := c.Ctx.Input.Query("name")
	if name != "xiaoyuzi" {
		c.Ctx.ResponseWriter.Write([]byte("对不起，校验错误"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("校验成功"))
	//c.Data["Website"] = "beego的主页"
	//c.Data["Email"] = "1162735365@QQ.com"
	//c.TplName = "index.tpl"
}
func (c *MainController) Post(){
	//name := c.Ctx.Request.FormValue("name")
	//if name != "xiaoyuzi" {
	//	c.Ctx.WriteString("对不起，校验错误")
	//	return
	//}
	fmt.Println("asdasdasdasdasdadasdads")
	//c.Ctx.WriteString("校验成功")
	var person =models.Person{Name:string("xiaoyuzi")}
	datebytes,err :=ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil{
		c.Ctx.WriteString("错误1")
		return
	}
	err = json.Unmarshal(datebytes,&person)
	if err !=nil {
		c.Ctx.WriteString("错误2")
		return
	}
	c.Ctx.WriteString("成功")
}