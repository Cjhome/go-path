package api

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers/hello"
	"main.go/requests"
	"net/http"
)

func Test(c *gin.Context)  {
	// 实例化一个TestRequest 结构体。用于接收参数
	testStruct := requests.TestRequest{}

	//接收请求参数
	err := c.ShouldBind(&testStruct)
	// 判断参数校验是否通过，如果不通过，把错误返回前端
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": requests.Translate(err),
		})
		return
	}

	// 调用HelloService
	var service hello.HelloContract

	// 这里使用的是接口定义了变量
	service = &hello.GreeterService{}
	//调用服务的方法处理业务
	result := service.SayHello(testStruct.Username)

	// 校验通过，返回请求参数
	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}