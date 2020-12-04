package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"main.go/pkg/config"
	"main.go/routers"
	"net/http"
	"os"
	"time"
)

//自定义个日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 可以通过上下文对象，设置一些依附在上下文对象里面的键/值数据
		c.Set("example", "123456")
		// 在这里处理请求到达控制器函数之前的逻辑
		//调用下一个中间件，或者控制器处理函数，具体得看注册多少个中间件
		c.Next()

		//在这里可以处理请求返回给用户之前的逻辑
		latency := time.Since(t)
		log.Println(latency)

		// 例如 查询请求状态码
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	// 设置日志等级
	log.SetLevel(log.DebugLevel)
	// 设置日志输出到什么地方去
	// 将日志输出到标准输出，就是直接在控制台打印出来
	log.SetOutput(os.Stdout)
	// 设置为true则显示日志在代码什么位置打印
	//log.SetReportCaller(true)
	// 设置日志以json格式输出， 如果不设置默认以text格式输出
	log.SetFormatter(&log.JSONFormatter{})


	// 打印日志
	log.Debug("调试信息")
	log.Info("提示信息")
	log.Warn("警告信息")
	log.Error("错误信息")
	//log.Panic("致命错误")
	//
	// 为日志加上字段信息，log.Fields 其实就是map[string]interface{}类别的别名
	log.WithFields(log.Fields{
		"user_id": 1001,
		"ip" : "123.12.12.11",
		"request_id" : "kana012uasdb8a918gad712",
	}).Info("用户登录失败")


	r := gin.New()
	r.Use(Logger())

	// 设置一个get请求的路由，url为/hello ,处理函数(或者叫控制器函数) 是一个闭包函数
	r.GET("/hello", func(context *gin.Context) {
		// 通过请求上下文对象Context, 直接往客户端返回一个json
		context.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	r.GET("/cookie", func(c *gin.Context) {
		c.SetCookie("site_cookie", "cookievalue", 3600, "/","localhost", false, true)
	})

	r.GET("getCookie", func(c *gin.Context) {
		data, err := c.Cookie("site_cookie")
		if err != nil {
			c.String(http.StatusOK, data)
			return
		}
		c.String(http.StatusOK, data)
	})
	r.GET("deleteCookie", func(c *gin.Context) {
		c.SetCookie("site_cookie", "cookievalue", -1, "/", "localhost", false, true)
		c.String(http.StatusOK, "删除成功")
	})

	//注册路由
	routers.Init(r)

	r.Run(fmt.Sprintf("%s:%d", config.Server.Address, config.Server.Port))
}
