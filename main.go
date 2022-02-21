package main

import (
	"douban/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(controller.Cors())
	router.Static("./static", "./static")
	routerEngine(router)
	router.Run(":9090")
}

func routerEngine(engine *gin.Engine) {
	new(controller.UserController).Router(engine)
	new(controller.CommentController).Router(engine)
	new(controller.MovieController).Router(engine)
	new(controller.CelebrityController).Router(engine)
	new(controller.SearchController).Router(engine)
	new(controller.MyLookController).Router(engine)
	new(controller.DiscussController).Router(engine)
	new(controller.ReplyController).Router(engine)
}
