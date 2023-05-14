package util

import (
	"dbo-test/model"
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, statusCode int, msg string, data interface{}) {
	c.JSON(statusCode, model.Response{
		Message: msg,
		Data:    data,
		Code:    statusCode,
		Status:  statusCode == 200,
	})
}
