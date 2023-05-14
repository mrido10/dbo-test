package controller

import (
	"dbo-test/service/customerService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Customer(route *gin.Engine, module string) {
	route.GET(fmt.Sprintf("/%s", module), customerService.Customer.List)
	route.POST(fmt.Sprintf("/%s", module), customerService.Customer.Insert)
	route.PUT(fmt.Sprintf("/%s", module), customerService.Customer.Update)
	route.DELETE(fmt.Sprintf("/%s", module), customerService.Customer.Delete)
	route.GET(fmt.Sprintf("/%s/:id", module), customerService.Customer.View)
}
