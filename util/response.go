package util

import (
	"dbo-test/model/dto"
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, statusCode int, msg string, data interface{}) {
	c.JSON(statusCode, dto.Response{
		Message: msg,
		Data:    data,
		Code:    statusCode,
		Status:  statusCode == 200,
	})
}
