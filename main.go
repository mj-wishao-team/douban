package main

import (
	"github.com/gin-gonic/gin"
	"movie.douban/controller"
)

func main() {
	router := gin.Default()
	router.Use(controller.Cors())
	routerEngine(router)
	router.Run(":9090")
}

func routerEngine(engine *gin.Engine) {
	new(controller.UserController).Router(engine)
}
