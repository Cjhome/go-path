package routers

import (
	"beego_project/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Get("get", func(context *context.Context) {
		beego.Info("基础路由中的get请求")
		context.Output.Body([]byte("基础路由中的get请求 get method"))
	})
}
