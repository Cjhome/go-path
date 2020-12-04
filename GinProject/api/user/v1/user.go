package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/dto"
	"main.go/pkg/e"
	v1 "main.go/requests/v1"
	"main.go/services/user"
)

func RegisterHandle(c *gin.Context) (interface{}, error) {
	request := v1.RegisterRequest{}

	err := c.ShouldBind(&request)
	if err != nil {
		return nil, e.ApiError{
			Status: 422,
			Code: 40004,
			Message: "参数校验失败",
			Data: err,
		}
	}

	userDto := dto.UserDto{
		Username: request.Username,
		Password: request.Password,
	}
	service := user.UserService{}
	err = service.Register(userDto)
	if err != nil {
		return nil, e.ApiError{
			Status: 422,
			Code: 4005,
			Message: err.Error(),
		}
	}
	return "注册成功", nil
}

func LoginHandle(c *gin.Context) (interface{}, error) {
	request := v1.LoginRequest{}

	err := c.ShouldBind(&request)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40004,
			Message: "参数校验失败",
			Data: err,
		}
	}

	userDto := dto.UserDto{
		Username: request.Username,
		Password: request.Password,
	}

	service := user.UserService{}
	err = service.Login(userDto)

	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40005,
			Message: err.Error(),
		}
	}

	return "登陆成功", nil
}
