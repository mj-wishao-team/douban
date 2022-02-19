package controller

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MyLookController struct {
}

func (M *MyLookController) Router(engine *gin.Engine) {

	engine.GET("/api/movie/mine", GetMyLookMovieHome)
	engine.GET("/api/people/reviews", GetSelfReviews)
}

//获取我的电影主页
func GetMyLookMovieHome(ctx *gin.Context) {
	accessToken := ctx.PostForm("access_token")
	refreshToken := ctx.PostForm("refresh_token")

	if accessToken != "" && refreshToken != "" {
		Claims, flag, err := service.ParseToken(accessToken, refreshToken)

		if err != nil {
			tool.RespErrorWithData(ctx, "token错误")
			fmt.Println("err", err)
			return
		}
		if flag {
			accessToken, err := service.GenToken(Claims.User, 3600*24, "ACCESS_TOKEN")
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

			MSWantLook, err := service.GetMyLook(Claims.User.Id, "想看")
			MSSeen, err := service.GetMyLook(Claims.User.Id, "看过")

			if err != nil {
				fmt.Println("GetMyLook Is ERR", err)
				tool.RespInternalError(ctx)
				return
			}

			Reviews, err := service.GetLargeCommentByUid(Claims.User.Id)
			if err != nil {
				tool.RespErrorWithData(ctx, "影评获取失败")
				fmt.Println("GetLargeCommentByUid err is", err)
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"看过":            MSSeen,
				"想看":            MSWantLook,
				"影评":            Reviews,
				"status":        "true",
				"refresh_token": refreshToken,
				"access_token":  accessToken,
			})

		} else {

			MSWantLook, err := service.GetMyLook(Claims.User.Id, "想看")
			fmt.Println("111111")

			MSSeen, err := service.GetMyLook(Claims.User.Id, "看过")
			fmt.Println("111111")

			if err != nil {
				fmt.Println("GetMyLook Is ERR", err)
				tool.RespInternalError(ctx)
				return
			}

			Reviews, err := service.GetLargeCommentByUid(Claims.User.Id)
			fmt.Println("111111")

			if err != nil {
				tool.RespErrorWithData(ctx, "影评获取失败")
				fmt.Println("GetLargeCommentByUid err is", err)
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"看过":            MSSeen,
				"想看":            MSWantLook,
				"影评":            Reviews,
				"status":        "true",
				"refresh_token": refreshToken,
				"access_token":  accessToken,
			})
		}
	} else {
		tool.RespErrorWithData(ctx, "请重新登录")
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
