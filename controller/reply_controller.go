package controller

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ReplyController struct {
}

func (C *ReplyController) Router(engine *gin.Engine) {

}

//TODO
//获取回复
func GetReply(ctx *gin.Context) {
	//pid
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		tool.RespErrorWithData(ctx, "id错误")
		return
	}
	kind := ctx.Query("type")
	if kind != "review" && kind != "discussion" && kind != "comment" && kind != "reply" {
		tool.RespErrorWithData(ctx, "类型错误")
		return
	}
	//一般都是1
	start, err := strconv.Atoi(ctx.Query("start"))
	if err != nil {
		start = 0

	}
	Reply, err := service.GetReply(id, kind, start)

	if err != nil && err.Error() != "sql: no rows in result set" {
		fmt.Println(err)
		tool.RespErrorWithData(ctx, "获取评论失败")
		return
	}
	tool.RespSuccessfulWithData(ctx, Reply)

}

//发布回复

//删除回复
