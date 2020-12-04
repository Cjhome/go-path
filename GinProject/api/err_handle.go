package api

import (
	"github.com/gin-gonic/gin"
	"main.go/pkg/e"
	"net/http"
)

func ErrorHandle(c *gin.Context) (interface{}, error) {
	query := c.Query("q")
	if query == "" {
		return nil, e.ApiError{
			Status: http.StatusOK,
			Code: 404,
			Message: "q的参数不能为空",
		}
	}

	if query == "test" {
		return nil, e.ApiError{
			Status: http.StatusOK,
			Code: 404,
			Message: "q的参数不能为test",
		}
	}
	return query, nil
}
