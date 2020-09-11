package routers

import (
	"beego-quickstart/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Get("get", func(context *context.Context) {
		beego.Info("基础路由中的get请求")
		context.Output.Body([]byte("基础路由中的get请求 get method"))
	})
	beego.Get("getUserInfo", func(context *context.Context) {
		beego.Info("获取用户信息")
		context.Output.Body([]byte("获取用户信息 get method"))
	})
    beego.Post("post", func(c *context.Context) {
		beego.Info("基础路由中的post请求")
		c.Output.Body([]byte("基础路由中的post请求 post method"))
	})

    // 固定路由
    beego.Router("GetInfo", &controllers.MainController{})

    // 正则路由
	beego.Router("/getUser/:id:int", &controllers.RegController{})
    beego.Router("/getUser2/:id[0-9]+", &controllers.RegController{})
    beego.Router("/getUser3/*.*", &controllers.RegController{})

    // 自定义路由
	beego.Router("GetInfo4", &controllers.MainController{}, "GET:GetUserInfo")
}
