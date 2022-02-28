package controller

import (
	"douban/dao"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
	//service.JudgeTokenAndMessage(accessToken,refreshToken,Id)
	if err != nil {
		tool.RespErrorWithData(ctx, err)
		fmt.Println("getMovie_ParseInt ERR is", err)
		return
	}

	if accessToken != "" && refreshToken != "" {
		Claims, flag, err := service.ParseToken(accessToken, refreshToken)
		if err != nil {
			tool.RespErrorWithData(ctx, "token错误")
			fmt.Println("err", err)
			return
		}
		if flag {
			accessToken, err := service.GenToken(Claims.User, 300, "ACCESS_TOKEN")
			if err != nil {
				fmt.Println("JWTAuthMiddleware_CreateAccessTokenErr:", err)
				tool.RespInternalError(ctx)
				ctx.Abort()
				return
			}

			//refreshToken 一周
			refreshToken, err := service.GenToken(Claims.User, 604800, "REFRESH_TOKEN")
			if err != nil {
				fmt.Println("JWTAuthMiddleware_CreateRefreshTokenErr:", err)
				tool.RespInternalError(ctx)
				ctx.Abort()
				return
			}

			SC, err := service.GetShortCommentByUidAndMid(Claims.User.Id, Id)
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
			ctx.JSON(http.StatusOK, gin.H{
				"Movies":         movies,
				"MyShortComment": SC,
				"acess_token":    accessToken,
				"refresh_token":  refreshToken,
				"status":         "true",
				"discussion":     Discussion,
				"comment":        sC,
				"reviews":        LC,
			})
			return
		} else {
			SC, err := service.GetShortCommentByUidAndMid(Claims.User.Id, Id)

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
			ctx.JSON(http.StatusOK, gin.H{
				"Movies":         movies,
				"MyShortComment": SC,
				"acess_token":    accessToken,
				"refresh_token":  refreshToken,
				"discussion":     Discussion,
				"comment":        sC,
				"reviews":        LC,
				"status":         "true",
			})
			return
		}

	} else {

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

		ctx.JSON(http.StatusOK, gin.H{
			"Movies":     movies,
			"status":     "true",
			"discussion": Discussion,
			"comment":    sC,
			"reviews":    LC,
		})
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
