package main

import (
	_ "beego_project/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	/**
	import _ 执行这个包下的init函数
	 */
	o := orm.NewOrm()
	o.Using("default")

	//profile := new(Profile)
	//profile.Age = 30
	beego.Run()
}
