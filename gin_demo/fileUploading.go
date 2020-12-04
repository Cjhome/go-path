package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	// 为multipart forms 设置较低的内存限制(默认 3MIB)
	r.MaxMultipartMemory = 8 << 20 // 8MIB
	r.POST("upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		log.Println(file.Filename)
		dst := "static/upload"
		// 上传文件至指定目录
		context.SaveUploadedFile(file, dst)
		context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	r.POST("manyUpload", func(context *gin.Context) {
		form, _ := context.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			//上传文件至指定目录
			dst := "static/upload"
			context.SaveUploadedFile(file, dst)
		}
		context.String(http.StatusOK, fmt.Sprintf("%d files upload!", len(files)))
	})

	r.Run(":8080")
}