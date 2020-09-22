package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego的主页"
	c.Data["Email"] = "1162735365@QQ.com"
	c.TplName = "index.tpl"
}
