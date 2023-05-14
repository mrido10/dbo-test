package controller

import (
	"dbo-test/service/tokenService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Token(route *gin.Engine, module string) {
	route.POST(fmt.Sprintf("/%s/generate", module), tokenService.Token.GenerateToken)
}
