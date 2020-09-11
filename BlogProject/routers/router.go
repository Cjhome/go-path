package routers

import (
	"BlogProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
    beego.Router("register", &controllers.RegisterController{})
    beego.Router("login", &controllers.LoginController{})
    beego.Router("exit", &controllers.ExitController{})
}