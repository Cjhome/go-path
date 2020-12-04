package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 你也可以使用自己的SecureJSON前缀
	// r.SecureJsonPrefix("")]}',\n")
	r.GET("someJSON", func(context *gin.Context) {
		name := []string{"lena", "austin", "foo"}

		// 将输出 while(1);["lena","austin", "foo"]
		context.SecureJSON(http.StatusOK, name)
	})
	r.Run(":8080")
}