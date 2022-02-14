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
	engine.GET("api/movie", GetMovieList)
	engine.GET("api/movie/explore", GetMovieListByTag)
	engine.GET("api/movie/chart", getMovieLeaderboard)
}

//获取单个电影信息
func getMovie(ctx *gin.Context) {
	Id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
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

//选电影
func GetMovieListByTag(ctx *gin.Context) {
	tag := ctx.PostForm("tag")
	//分类
	sort := ctx.PostForm("sort")
	//个数限制
	limit, _ := strconv.Atoi(ctx.PostForm("limit"))

	MovieList, err := service.GetMovieListByTag(tag, sort, limit)
	if err != nil {
		tool.RespErrorWithData(ctx, "未找到相关条件")
		return
	}
	tool.RespSuccessfulWithData(ctx, MovieList)
}

//获取电影列表
func GetMovieList(ctx *gin.Context) {

}

//排行榜
func getMovieLeaderboard(ctx *gin.Context) {
	var limit int = 20
	service.GetMovieLeaderboard(limit)
}
