package controller

import (
	"dbo-test/service/orderService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Order(route *gin.Engine, module string) {
	route.POST(fmt.Sprintf("/%s", module), orderService.Order.Insert)
	route.GET(fmt.Sprintf("/%s", module), orderService.Order.List)
	route.PUT(fmt.Sprintf("/%s", module), orderService.Order.Update)
	route.DELETE(fmt.Sprintf("/%s", module), orderService.Order.Delete)
	route.GET(fmt.Sprintf("/%s/:id", module), orderService.Order.View)
}
