package customerService

import (
	"dbo-test/util"
	"github.com/gin-gonic/gin"
)

type customer struct {
	util.Auth
}

var Customer = customer{}

func (e customer) View(c *gin.Context) {
	e.JWTValidations(c, GetDetail)
}

func (e customer) List(c *gin.Context) {
	e.JWTValidations(c, GetList)
}

func (e customer) Insert(c *gin.Context) {
	e.JWTValidations(c, InsertCustomer)
}

func (e customer) Update(c *gin.Context) {
	e.JWTValidations(c, UpdateCustomer)
}

func (e customer) Delete(c *gin.Context) {
	e.JWTValidations(c, DeleteCustomer)
}
