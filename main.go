package main

import (
	"douban/controller"
	"douban/tool"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := tool.GetCfg()
	router := gin.Default()
	router.Use(controller.Cors())
	routerEngine(router)
	router.Run(":" + cfg.AppPort)
}

func routerEngine(engine *gin.Engine) {
	new(controller.UserController).Router(engine)

}
