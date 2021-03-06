package controller

import (
	"douban/model"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type DiscussController struct {
}

func (D *DiscussController) Router(engine *gin.Engine) {
	engine.POST("/api/movie/discussion/put_discuss", JWTAuthMiddleware(), putDiscuss)
	engine.DELETE("api/movie/dicussion/delele_discuss/:id", JWTAuthMiddleware(), deleteDiscuss)
	engine.PUT("api/movie/discussion/updata", JWTAuthMiddleware(), updateDiscuss)
	engine.POST("api/movie/discussion/like/:id", JWTAuthMiddleware(), discussLike)

	engine.GET("api/movie/discussions/:mid", GetDiscussionList)

	engine.GET("api/movie/discussion/:id", GetDiscussion)
}

//发表讨论
func putDiscuss(ctx *gin.Context) {
	mid, err := strconv.ParseInt(ctx.PostForm("mid"), 10, 64)
	if err != nil {
		tool.RespErrorWithData(ctx, "解析错误")
		fmt.Println("", err)
		return
	}

	Discussions := model.Discussion{
		Uid:   ctx.MustGet("id").(int64),
		Mid:   mid,
		Title: ctx.PostForm("title"),
		Value: ctx.PostForm("value"),
		Date:  time.Now(),
		Stars: 0,
	}

	err = service.PutDiscussion(Discussions)
	if err != nil {
		tool.RespErrorWithData(ctx, "发表失败")
		fmt.Println("PutDiscusion is ERR ", err)
		return
	}
	tool.RespSuccessfulWithData(ctx, "发布成功")
}

//获取讨论列表
func GetDiscussionList(ctx *gin.Context) {
	//分类获取 默认为热度排名
	sort := ctx.PostForm("sort")

	mid, err := strconv.ParseInt(ctx.Param("mid"), 10, 64)
	if err != nil {
		tool.RespErrorWithData(ctx, "解析错误")
		fmt.Println("", err)
		return
	}
	dL, err := service.GetDiscussionList(sort, mid)
	if err != nil {
		tool.RespErrorWithData(ctx, "")
		fmt.Println("GetDiscussion IS ERR", err)
		return
	}
	tool.RespSuccessfulWithData(ctx, dL)

}

//获取讨论
func GetDiscussion(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("Getdiscussion is ERR ", err)
		return
	}
	Discussion, err := service.GetDiscussion(id)

	if err != nil && err.Error() != "sql: no rows in result set" {
		tool.RespInternalError(ctx)
		fmt.Println("GetDiscussion is ERR ", err)
		return
	}
	reply, err := service.GetReply(id, "discussion", 1)

	ctx.JSON(http.StatusOK, gin.H{
		"Discussion": Discussion,
		"Reply":      reply,
		"data":       "true",
	})
}

//删除讨论
func deleteDiscuss(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("deleteDiscuss is ERR ", err)
		return
	}
	Discussion, err := service.GetDiscussion(id)

	if Discussion.Uid != ctx.MustGet("id").(int64) {
		tool.RespErrorWithData(ctx, "不是本用户")
		return
	}
	err = service.DeleteDisucuss(id)

	if err != nil {
		tool.RespErrorWithData(ctx, "删除失败")
		fmt.Println("deleteDiscussion is ERR", err)
		return
	}
	tool.RespSuccessfulWithData(ctx, "删除成功")

}

//跟新讨论
func updateDiscuss(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.PostForm("id"), 10, 64)
	if err != nil {
		tool.RespErrorWithData(ctx, "解析错误")
		fmt.Println("", err)
		return
	}

	Discussions := model.Discussion{
		Id:       id,
		Uid:      ctx.MustGet("id").(int64),
		UserName: ctx.MustGet("username").(string),
		Title:    ctx.PostForm("tile"),
		Value:    ctx.PostForm("value"),
		Date:     time.Now(),
		Stars:    0,
	}

	err = service.UpdateDiscussion(Discussions)
	if err != nil {
		tool.RespErrorWithData(ctx, "发表失败")
		fmt.Println("PutDiscusion is ERR ", err)
		return
	}
	tool.RespSuccessfulWithData(ctx, "跟新成功")
}

//点赞or取消
func discussLike(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		tool.RespErrorWithData(ctx, "解析错误")
		fmt.Println("", err)
		return
	}
	err = service.DiscussLike(id)
	if err != nil {
		fmt.Println(err)
		tool.RespErrorWithData(ctx, "点赞失败")
		return
	}
	tool.RespSuccessfulWithData(ctx, "点赞成功")
}
