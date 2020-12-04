package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()
	//engine.GET("hello", func(context *gin.Context) {
	//	fmt.Println("请求路径：", context.FullPath())
	//	context.Writer.Write([]byte("hello, gin\n"))
	//})
	engine.LoadHTMLGlob("./html/*")
	// 加载静态文件 映射  第一个参数是请求，第二个是映射的路径
	//engine.Static("/img", "./img")
	routerGroup := engine.Group("/user")

	routerGroup.POST("/register", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var register Register
		if err := context.ShouldBind(&register); err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(register.UserName)
		fmt.Println(register.Phone)
		context.Writer.Write([]byte(register.UserName + "Register"))
	})

	routerGroup.POST("/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var register Register
		if err := context.ShouldBind(&register); err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(register.UserName)
		fmt.Println(register.Phone)
		context.Writer.Write([]byte(register.UserName + "登录"))
	})

	engine.Handle("GET", "/hello", func(context *gin.Context) {
		// fullpath 获取解析的url
		//path := context.FullPath()
		//fmt.Println(path)
		//// DefaultQuery url带的参数  两个参数  key，默认值
		////name := context.DefaultQuery("name", "hello")
		//name := context.Query("name")
		//fmt.Println(name)

		var student Student
		err := context.ShouldBindQuery(&student)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(student.Name, student.Classes)

		// 输出
		context.Writer.Write([]byte("Hello," + student.Name))
	})

	engine.Handle("POST", "login", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		// PostForm 解析post
		//username := context.PostForm("username")
		// bool类型 exist是否获取了这个值
		username, exist := context.GetPostForm("username")
		if exist {
			fmt.Println(username)
			return
		}
		password := context.PostForm("password")
		fmt.Println(username, password)
		context.Writer.Write([]byte(username + "登录"))
	})

	engine.POST("/register", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var register Register
		if err := context.ShouldBind(&register); err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(register.UserName)
		fmt.Println(register.Phone)
		context.Writer.Write([]byte(register.UserName + "Register"))
	})

	engine.POST("addstudent", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var person Person
		if err := context.BindJSON(&person); err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("姓名", person.Name)
		fmt.Println("年龄", person.Age)
		context.Writer.Write([]byte("添加记录" + person.Name))
	})

	engine.DELETE("user/:id", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		userID := context.Param("id")
		fmt.Println(userID)
		context.Writer.Write([]byte("delete 用户" + userID))
	})

	engine.GET("/hellostring", func(context *gin.Context) {
		fullPath := "请求路径" + context.FullPath()
		fmt.Println(fullPath)
		context.Writer.WriteString(fullPath)
	})

	engine.GET("/hellojson", func(context *gin.Context) {
		fullPath := "请求路径" + context.FullPath()
		fmt.Println(fullPath)
		context.JSON(200,map[string]interface{}{
			"code": 0,
			"message": "请求成功",
			"data": fullPath,
		})
	})

	engine.GET("/jsonstruct", func(context *gin.Context) {
		fullPath := "请求路径" + context.FullPath()
		fmt.Println(fullPath)
		resp := Response{Code: 1, Message: "OK", Data: fullPath}
		context.JSON(200, &resp)

	})

	engine.GET("/hellohtml", func(context *gin.Context) {
		fullPath := "请求路径" + context.FullPath()
		fmt.Println(fullPath)
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "测试",
		})
	})

	if err := engine.Run(":8090"); err != nil {
		log.Fatal(err.Error())
	}
}

type Student struct {
	Name string `form:"name"`
	Classes string `form:"classes"`
}

type Register struct {
	UserName string `form:"name"`
	Phone string `form:"phone"`
	Password string `form:"password"`
}

type Person struct {
	Name string `form:"name"`
	Sex string `form:"sex"`
	Age int64 `form:"age"`
}

type Response struct {
	Code int
	Message string
	Data interface{}
}