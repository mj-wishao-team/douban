package controller

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type MovieController struct {
}

func (M *MovieController) Router(engine *gin.Engine) {
	engine.GET("api/movie/subject/:id", getMovie)
	engine.POST("api/movie/comment", JWTAuthMiddleware(), putMovieComment)
}

//获取电影信息
func getMovie(ctx *gin.Context) {
	Id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		tool.RespErrorWithData(ctx, err)
		fmt.Println("getMovie_ParseInt ERR is", err)
		return
	}

	movies, err := service.GetMovieById(Id)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("getMovie_GetMovieById ERR is", err)
		return
	}
	tool.RespSuccessfulWithData(ctx, movies)
}
