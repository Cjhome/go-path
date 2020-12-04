package v1

import (
	v1 "main.go/api/user/v1"
	"main.go/pkg/e"
)

import "github.com/gin-gonic/gin"

func UserRouter(r *gin.Engine) {
	r.POST("/register", e.ErrorWrapper(v1.RegisterHandle))


	//登陆
	r.POST("/login", e.ErrorWrapper(v1.LoginHandle))
}
