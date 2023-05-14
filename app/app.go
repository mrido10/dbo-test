package app

import (
	"dbo-test/config"
	"dbo-test/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func Start() {
	route := gin.Default()
	controller.Token(route, "token")
	controller.Customer(route, "customer")
	if err := route.Run(":" + config.Configure.Server.Port); err != nil {
		log.Fatal(err)
	}
}
