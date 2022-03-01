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

type MyLookController struct {
}

func (M *MyLookController) Router(engine *gin.Engine) {

	engine.GET("/api/movie/mine", GetMyLookMovieHome)
	engine.GET("/api/people/reviews", JWTAuthMiddleware(), GetSelfReviews)
}

//获取我的电影主页
func GetMyLookMovieHome(ctx *gin.Context) {
	accessToken := ctx.PostForm("access_token")
	refreshToken := ctx.PostForm("refresh_token")

	accessToken, refreshToken, uid := service.JudgeToken(accessToken, refreshToken, ctx)

	if uid != 0 && accessToken != "" && refreshToken != "" {

		var Online = make(map[string]interface{})
		rp := dao.NewRedisStore("MyLook_"+"_"+strconv.FormatInt(uid, 10), time.Hour*24, ctx)
		if flag := rp.GetRedisPages(rp.PreKey, &Online); flag == true {
			tool.RespSuccessfulWithData(ctx, Online)
			return
		}

		MSSeen, err := service.GetMyLook(uid)

		if err != nil && err.Error() != "sql: no rows in result set" {
			fmt.Println("GetMyLook Is ERR", err)
			tool.RespInternalError(ctx)
			return
		}

		Reviews, err := service.GetLargeCommentByUid(uid)
		if err != nil && err.Error() != "sql: no rows in result set" {
			tool.RespErrorWithData(ctx, "影评获取失败")
			fmt.Println("GetLargeCommentByUid err is", err)
			return
		}
		Online = map[string]interface{}{
			"Movies":        MSSeen,
			"影评":            Reviews,
			"status":        "true",
			"refresh_token": refreshToken,
			"access_token":  accessToken,
		}
		rp.SetRedisPages(rp.PreKey, &Online)
		tool.RespSuccessfulWithData(ctx, Online)
	} else {
		tool.RespErrorWithData(ctx, "未登录或token错误")
	}
}

//获取自己的影评
func GetSelfReviews(ctx *gin.Context) {
	uid := ctx.MustGet("id").(int64)
	Comment, err := service.GetLargeCommentByUid(uid)
	if err != nil {
		tool.RespErrorWithData(ctx, "获取评论失败")
		fmt.Println("GetSelfReviews_GetLargeCommentByUid is ERR", err)
		return
	}
	tool.RespSuccessfulWithData(ctx, Comment)
}
