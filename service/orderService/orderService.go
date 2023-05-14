package orderService

import (
	"dbo-test/util"
	"github.com/gin-gonic/gin"
)

type order struct {
	util.Auth
}

var Order = order{}
func (e order) Insert(c *gin.Context) {
	e.JWTValidations(c, InsertOrder)
}
func (e order) List(c *gin.Context) {
	e.JWTValidations(c, GetList)
}
