package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.POST("/form_post", func(context *gin.Context) {
		message :=context.PostForm("message")
		nick := context.DefaultPostForm("nick", "anonymous")

		context.JSON(200, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})
	router.Run(":8080")
}