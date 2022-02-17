package controller

import (
	"douban/model"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type DiscussController struct {
}

func (D *DiscussController) Router(engine *gin.Engine) {
	engine.POST("api/movie/discussion/put_discuss", JWTAuthMiddleware(), putDiscuss)
	engine.DELETE("api/movie/dicussion/delele_discuss", deleteDiscuss)
	engine.PUT("api/movie/discussion/updata", updateDiscuss)
	engine.POST("api/movie/discussion/:id/like", discussLike)
	engine.GET("api/movie/discussion/:mid", GetDiscussion)
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
		Uid:      ctx.MustGet("id").(int64),
		Mid:      mid,
		UserName: ctx.MustGet("username").(string),
		Title:    ctx.PostForm("tile"),
		Value:    ctx.PostForm("value"),
		Date:     time.Now(),
		Stars:    0,
	}

	err = service.PutDiscussion(Discussions)
	if err != nil {
		tool.RespErrorWithData(ctx, "发表失败")
		fmt.Println("PutDiscusion is ERR ", err)
		return
	}
}

//获取讨论列表
func GetDiscussion(ctx *gin.Context) {
	//分类获取 默认为热度排名
	sort := ctx.PostForm("sort")

	mid, err := strconv.ParseInt(ctx.PostForm("mid"), 10, 64)
	if err != nil {
		tool.RespErrorWithData(ctx, "解析错误")
		fmt.Println("", err)
		return
	}
	dL, err := service.GetDiscussion(sort, mid)
	if err != nil {
		tool.RespErrorWithData(ctx, "")
		fmt.Println("GetDiscussion IS ERR", err)
		return
	}
	tool.RespSuccessfulWithData(ctx, dL)

}

//删除讨论
func deleteDiscuss(ctx *gin.Context) {

}

//跟新讨论
func updateDiscuss(ctx *gin.Context) {

}

//点赞or取消
func discussLike(ctx *gin.Context) {

}
