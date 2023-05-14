package tokenService

import (
	"dbo-test/util"
	"github.com/gin-gonic/gin"
)

type token struct {
	util.Auth
}

var Token = token{}

func (e token) GenerateToken(c *gin.Context) {
	e.IgnoreValidateJWT(c, GenerateToken)
}
