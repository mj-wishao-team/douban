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
	engine.GET("api/mine", JWTAuthMiddleware(), GetMyLookMovieHome)
}

func GetMyLookMovieHome(ctx *gin.Context) {
	uid := ctx.MustGet("id").(int64)
	MSWantLook, err := service.GetMyLook(uid, "想看")
	MSSeen, err := service.GetMyLook(uid, "看过")
	if err != nil && err.Error() != "sql: no rows in result set" {
		fmt.Println("GetMyLook Is ERR", err)
		tool.RespInternalError(ctx)
		return
	}

	Reviews, err := service.GetLargeCommentByUid(uid)
	if err != nil {
		tool.RespErrorWithData(ctx, "影评获取失败")
		fmt.Println("GetLargeCommentByUid err is", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"看过":     MSSeen,
		"想看":     MSWantLook,
		"影评":     Reviews,
		"status": "true",
	})
}
