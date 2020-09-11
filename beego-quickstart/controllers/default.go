package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type RegController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func (c *MainController) GetUserInfo() {
	beego.Info("获取用户信息")
	username := c.GetString("username")
	userid := c.GetString("userid")
	c.Ctx.Output.Body([]byte(username + ":" + userid))
}

func (c *RegController) Get() {
	beego.Info("RegController Get")
	//id := c.Ctx.Input.Param(":id")
	//beego.Info("Id" + id)
	//c.Ctx.ResponseWriter.Write([]byte("id:" + id))

	path := c.Ctx.Input.Param(":path")
	beego.Info(path)
	ext := c.Ctx.Input.Param(":ext")
	beego.Info(ext)
}

func (c *RegController) Post() {
	beego.Info("RegController Post")
	c.Ctx.Output.Body([]byte("正则路由 Post "))
}
