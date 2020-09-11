package main

import (
	_ "BlogProject/routers"
	"BlogProject/utils"
	"github.com/astaxie/beego"
)

func main() {
	utils.InitMysql()
	//beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

