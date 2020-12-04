package routers

import (
	"github.com/gin-gonic/gin"
	"main.go/api"
	"main.go/pkg/e"
)


func test(r *gin.Engine) {
	r.GET("/test", api.Test)

	r.GET("/json", api.Json)

	r.GET("/xml", api.Xml)

	//统一错误处理测试
	r.GET("/error", e.ErrorWrapper(api.ErrorHandle))
}