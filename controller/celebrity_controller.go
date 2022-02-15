package controller

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CelebrityController struct {
}

func (C *CelebrityController) Router(engine *gin.Engine) {
	engine.GET("api/movie/celebrity/:id", GetCelebrity)
}
func GetCelebrity(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Query("id"), 10, 64)
	Celebrity, err := service.GetCelebrity(id)
	if err != nil {
		fmt.Println("GetCelebrity_GetCelebrity is ERR", err)
		tool.RespErrorWithData(ctx, "未有相关信息")
		return
	}
	tool.RespSuccessfulWithData(ctx, Celebrity)
}
