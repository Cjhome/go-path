package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	User string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login", func(context *gin.Context) {
		// 你可以使用显式绑定 multipart form
		// c.ShouldBindWith(&form, binding.Form)
		//或者简单地使用shouldBind 方法自动绑定
		var form LoginForm
		if context.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				context.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})
	router.Run(":8080")
}