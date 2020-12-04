package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

// 自定义模版功能
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	router := gin.Default()
	// 自定义模版渲染器
	html := template.Must(template.ParseFiles("file1", "file2"))
	router.SetHTMLTemplate(html)
	// 自定义分隔符
	router.Delims("{[{","}]}")
	// 自定义模版功能
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLFiles("./testdata/template/raw.html")

	router.LoadHTMLGlob("/path/to/templates")
	//router.LoadHTMLGlob("html/**、*") //使用不同目录下名称相同的模版
	router.LoadHTMLGlob("html/*")

	router.GET("raw", func(context *gin.Context) {
		context.HTML(http.StatusOK, "/raw", map[string]interface{}{
			"now": time.Date(2017,07,01,0,0,0, 0, time.UTC()),
		})
	})

	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.Run(":8888")
}
