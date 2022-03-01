package controller

import (
	"douban/dao"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type MovieController struct {
}

func (M *MovieController) Router(engine *gin.Engine) {
	engine.GET("/api/movie/subject/:mid", getMovie)
	engine.GET("/api/movie", GetMovieList)
	engine.GET("/api/movie/sort", GetMovieListByTag)
	engine.GET("/api/movie/chart", getMovieLeaderboard)
}

//获取单个电影信息
func getMovie(ctx *gin.Context) {

	accessToken := ctx.PostForm("access_token")
	refreshToken := ctx.PostForm("refresh_token")

	Id, err := strconv.ParseInt(ctx.Param("mid"), 10, 64)
	accessToken, refreshToken, uid := service.JudgeToken(accessToken, refreshToken, ctx)
	if err != nil {
		tool.RespErrorWithData(ctx, err)
		fmt.Println("getMovie_ParseInt ERR is", err)
		return
	}

	if uid != 0 && accessToken != "" && refreshToken != "" {

		var Online = make(map[string]interface{})
		rp := dao.NewRedisStore("movie_"+ctx.Param("mid")+"_"+strconv.FormatInt(uid, 10), time.Hour*24, ctx)
		if flag := rp.GetRedisPages(rp.PreKey, &Online); flag == true {
			tool.RespSuccessfulWithData(ctx, Online)
			return
		}

		SC, err := service.GetShortCommentByUidAndMid(uid, Id)
		if err != nil {
			tool.RespInternalError(ctx)
			fmt.Println(err)
			return
		}
		movies, err := service.GetMovieById(Id)
		if err != nil {
			tool.RespInternalError(ctx)
			fmt.Println("getMovie_GetMovieById ERR is", err)
			return
		}

		Discussion, err := service.GetDiscussionList("time", Id)
		if err != nil {
			tool.RespInternalError(ctx)
			fmt.Println("getMovie_GetDiscussionList ERR is", err)
			return
		}

		LC, err := service.GetMovieReviews(Id)
		if err != nil {
			tool.RespInternalError(ctx)
			fmt.Println("getMovie_GetLargeCommentSlice ERR is", err)
			return
		}
		sC, err := service.GetMovieComment(Id)
		if err != nil {
			tool.RespInternalError(ctx)
			fmt.Println("getMovie_GetShortCommentSlice ERR is", err)
			return
		}
		Online = map[string]interface{}{
			"Movies":         movies,
			"MyShortComment": SC,
			"acess_token":    accessToken,
			"refresh_token":  refreshToken,
			"discussion":     Discussion,
			"comment":        sC,
			"reviews":        LC,
		}
		rp.SetRedisPages(rp.PreKey, &Online)
		tool.RespSuccessfulWithData(ctx, Online)
	} else {
		var tourist = make(map[string]interface{})
		rp := dao.NewRedisStore("movie_"+ctx.Param("mid"), time.Hour*24, ctx)
		if flag := rp.GetRedisPages(rp.PreKey, &tourist); flag == true {
			tool.RespSuccessfulWithData(ctx, tourist)
			return
		}

		movies, err := service.GetMovieById(Id)
		if err != nil {
			tool.RespInternalError(ctx)
			fmt.Println("getMovie_GetMovieById ERR is", err)
			return
		}

		Discussion, err := service.GetDiscussionList("time", Id)
		if err != nil {
			tool.RespInternalError(ctx)
			fmt.Println("getMovie_GetDiscussionList ERR is", err)
			return
		}

		LC, err := service.GetMovieReviews(Id)
		if err != nil {
			tool.RespInternalError(ctx)
			fmt.Println("getMovie_GetLargeCommentSlice ERR is", err)
			return
		}

		sC, err := service.GetMovieComment(Id)
		if err != nil {
			tool.RespInternalError(ctx)
			fmt.Println("getMovie_GetShortCommentSlice ERR is", err)
			return
		}

		tourist = map[string]interface{}{
			"Movies":     movies,
			"discussion": Discussion,
			"comment":    sC,
			"reviews":    LC,
		}
		rp.SetRedisPages(rp.PreKey, &tourist)
		tool.RespSuccessfulWithData(ctx, tourist)
	}
}

//选电影
func GetMovieListByTag(ctx *gin.Context) {
	tag := ctx.PostForm("tag")
	//分类
	sort := ctx.PostForm("sort")
	//从第几行开始
	start, _ := strconv.Atoi(ctx.PostForm("start"))

	MovieList, err := service.GetMovieListByTag(tag, sort, start)
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

	start, err := strconv.Atoi(ctx.PostForm("start"))
	if err != nil {
		tool.RespErrorWithData(ctx, "解析失败")
		fmt.Println("Parase Is ERR ", err)
		return
	}
	rp := dao.NewRedisStore("Leaderboard"+strconv.Itoa(start), time.Hour*24, ctx)
	ML, err := service.GetMovieLeaderboard(start, rp)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("GetMovieLeaderboard Is ERR", err)
		return
	}
	tool.RespSuccessfulWithData(ctx, ML)
}
