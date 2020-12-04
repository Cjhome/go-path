package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	r := gin.Default()

	//gin.H 是map[string]interface{} 的一种快捷方式
	r.GET("someJSON", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hey",
			"status": http.StatusOK,
		})
	})

	r.GET("moreJSON", func(context *gin.Context) {
		//你也可以使用一个结构体
		var msg struct{
			Name string `json:"user"`
			Message string
			Number int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// 注意 msg.Name 在 JSON 中变成了 “user”
		// 将输出：{"user":"Lena", "Message":"hey", "Number": 123}
		context.JSON(http.StatusOK, msg)
	})

	r.GET("someXML", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{"messge": "Hey", "status": http.StatusOK})
	})
	r.GET("someYML", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
	r.GET("someProtoBuf", func(context *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf 的具体定义写在testdata/protoexample 文件中。
		data := &protoexample.Test{
			Label: &label,
			Reps: reps,
		}
		//请注意，数据在响应中变为二进制数据
		// 将输入被protoexample.Test protobuf 序列化了的数据
		context.ProtoBuf(http.StatusOK, data)
	})
	r.Run(":8080")
}