package main

import (
	_ "BlogProject/routers"
	"BlogProject/utils"
	"github.com/astaxie/beego"
	//_ "github.com/russross/blackfriday" // Markdown文档替换
	//_ "github.com/PuerkitoBio/goquery" // 类似jQuery DOM操作
	//_ "github.com/sourcegraph/syntaxhighlight" // 代码的语法高亮显示
)

func main() {
	utils.InitMysql()
	//beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

