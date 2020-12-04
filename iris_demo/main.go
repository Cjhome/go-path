package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	//Method GET
	// Resource http://localhost:8080
	//app.Get("/", func(ctx iris.Context) {
	//	ctx.HTML("<h1>Welcome Hello Word!</h1>")
	//})
	////same as app.Handle("GET", "/ping", [...])
	////Method GET
	////Resource http://localhost:8080/ping
	//app.Get("/ping", func(ctx iris.Context) {
	//	ctx.WriteString("pong")
	//})
	//
	////Method GET
	////Resource http://localhost:8080/hello
	//app.Get("/hello", func(ctx iris.Context) {
	//	ctx.JSON(iris.Map{"message": "Hello Word"})
	//})

	app.Contr

	app.Run(iris.Addr(":8080"))
}