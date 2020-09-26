package routers

import (
	"xiaoyuzi/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/register",&controllers.Registercontrooler{})
}
