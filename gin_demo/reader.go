package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("someDataFromReader", func(context *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			context.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLenght := response.ContentLength
		contextType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment;filename="gopher.png"`,
		}
		context.DataFromReader(http.StatusOK, contentLenght, contextType, reader, extraHeaders)
	})
}
