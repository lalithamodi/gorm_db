package main

import (
	"main/config"
	"main/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8081")

}
