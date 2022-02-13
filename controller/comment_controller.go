package controller

import (
	"douban/model"
	"douban/param"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type CommentController struct {
}

func (C *CommentController) Router(engine *gin.Engine) {
	engine.POST("api/movie/comment/put_short", JWTAuthMiddleware(), putMovieShortComment)
	engine.POST("api/movie/comment/put_large", JWTAuthMiddleware(), putMovieLargeComment)
	engine.GET("api/movie/comment/get_short", JWTAuthMiddleware(), getShortComment)
	engine.GET("api/movie/comment/get_large", JWTAuthMiddleware(), getLargeComment)
}

//发表短评
func putMovieShortComment(ctx *gin.Context) {
	movieId, err := strconv.ParseInt(ctx.Param("mid"), 10, 64)

	uid := ctx.MustGet("id").(int64)
	star, err := strconv.Atoi(ctx.PostForm("star"))

	if err != nil {
		tool.RespErrorWithData(ctx, "解析失败")
		fmt.Println("putMovieShortComment is ERR", err)
		return
	}

	shortComment := ctx.PostForm("comment")

	if err != nil {
		fmt.Println("putMovieShortComment_ParseInt ERR is ", err)
		tool.RespInternalError(ctx)
		return
	}

	ShorComment := model.ShortComment{
		Mid:     movieId,
		Uid:     uid,
		Comment: shortComment,
		Time:    time.Now(),
		Star:    star,
	}

	err = service.PutMovieShortComment(ShorComment)
	if err != nil {
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, "评论成功")
}

//获取短评
func getShortComment(ctx *gin.Context) {
	Mid, err := strconv.ParseInt(ctx.Query("mid"), 10, 64)
	if err != nil {
		tool.RespErrorWithData(ctx, "解析失败")
		return
	}
	commentSlice, err := service.GetShortCommentSlice(Mid)

	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("getShortComment_GetShortCommentSlice Err is", err)
		return
	}

	if commentSlice == nil {
		commentSlice = []model.ShortComment{}
	}

	var newShortCommentSlice []param.ShortComment

	for _, commentModel := range commentSlice {
		var shortCommentParam param.ShortComment
		user, _ := service.GetUserById(commentModel.Uid)

		shortCommentParam.Time = commentModel.Time.Format("2006-01-02 15:04:05")
		shortCommentParam.Id = commentModel.Id
		shortCommentParam.Comment = commentModel.Comment
		shortCommentParam.User = user
		shortCommentParam.Help = commentModel.Help
		shortCommentParam.MId = commentModel.Mid

		newShortCommentSlice = append(newShortCommentSlice, shortCommentParam)
	}

	tool.RespSuccessfulWithData(ctx, newShortCommentSlice)

}

//发表影评
func putMovieLargeComment(ctx *gin.Context) {
	movieId, err := strconv.ParseInt(ctx.Param("mid"), 10, 64)

	title := ctx.PostForm("title")

	uid := ctx.MustGet("id").(int64)

	star, err := strconv.Atoi(ctx.PostForm("star"))

	if err != nil {
		tool.RespErrorWithData(ctx, "解析失败")
		fmt.Println("putMovieShortComment is ERR", err)
		return
	}

	largeComment := ctx.PostForm("comment")

	if err != nil {
		fmt.Println("putMovieShortComment_ParseInt ERR is ", err)
		tool.RespInternalError(ctx)
		return
	}

	LargeComment := model.LargeComment{
		Mid:     movieId,
		Uid:     uid,
		Title:   title,
		Comment: largeComment,
		Time:    time.Now(),
		Star:    star,
	}

	err = service.PutMovieLargeComment(LargeComment)
	if err != nil {
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, "评论成功")
}

//获取影评
func getLargeComment(ctx *gin.Context) {

}
