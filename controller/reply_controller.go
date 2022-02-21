package controller

import (
	"douban/model"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"html"
	"strconv"
	"time"
)

type ReplyController struct {
}

func (C *ReplyController) Router(engine *gin.Engine) {
	engine.GET("/api/movie/reply/:id", GetReply)
	engine.POST("/api/movie/reply/post", JWTAuthMiddleware(), ReplyPost)
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
	if kind != "review" && kind != "discussion" && kind != "reply" {
		tool.RespErrorWithData(ctx, "类型错误")
		return
	}
	//一般都是1
	start, err := strconv.Atoi(ctx.PostForm("start"))
	if err != nil {
		tool.RespErrorWithData(ctx, "start 错误")
		fmt.Println(err)
		return
	}
	Reply, err := service.GetReply(id, kind, start)

	if err != nil {
		fmt.Println(err)
		tool.RespErrorWithData(ctx, "获取评论失败")
		return
	}
	tool.RespSuccessfulWithData(ctx, Reply)
	//&& err.Error() != "sql: no rows in result set"
}

//发布回复
func ReplyPost(ctx *gin.Context) {

	pid, err := strconv.ParseInt(ctx.PostForm("pid"), 10, 64)
	if err != nil {
		tool.RespErrorWithData(ctx, "pid 格式错误")
		return
	}
	kind := ctx.PostForm("type")
	if kind != "review" && kind != "discussion" && kind != "reply" {
		tool.RespErrorWithData(ctx, "type 格式错误")
		return
	}

	Reply := model.Reply{
		Uid:     ctx.MustGet("id").(int64),
		Pid:     pid,
		Ptable:  kind,
		Date:    time.Now(),
		Content: html.UnescapeString(html.EscapeString(ctx.PostForm("value"))),
	}

	err = service.ReplyPost(Reply)
	if err != nil {
		fmt.Println(err)
		tool.RespErrorWithData(ctx, "回复失败")
		return
	}

	switch kind {
	case "review":
		err = service.UpdateReviewCNT(pid)
		if err != nil {
			fmt.Println(err)
			tool.RespErrorWithData(ctx, "增加回复人数失败")
			return
		}
	case "reply":
		err := service.UpdateReplyCNT(pid)
		if err != nil {
			fmt.Println(err)
			tool.RespErrorWithData(ctx, "增加回复人数失败")
			return
		}
	case "discussion":
		err := service.UpdateDiscussionCNT(pid)
		if err != nil {
			fmt.Println(err)
			tool.RespErrorWithData(ctx, "增加回复人数失败")
			return
		}

	}

	tool.RespSuccessfulWithData(ctx, "回复成功")

}

//删除回复
