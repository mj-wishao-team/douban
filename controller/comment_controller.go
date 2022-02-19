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
	engine.POST("api/movie/comment/put_short/:mid", JWTAuthMiddleware(), putMovieShortComment)
	engine.POST("api/movie/comment/put_large", JWTAuthMiddleware(), putMovieLargeComment)
	engine.GET("api/movie/comment/get_short", getShortComment)
	engine.GET("api/movie/comment/get_large", getLargeComment)
	engine.GET("api/movie/short_comment/:id/add_like", JWTAuthMiddleware(), addShortCommentLike)
}

//短评点赞
func addShortCommentLike(ctx *gin.Context) {

}

//发表短评
func putMovieShortComment(ctx *gin.Context) {
	movieId, err := strconv.ParseInt(ctx.Param("mid"), 10, 64)
	uid := ctx.MustGet("id").(int64)
	//fmt.Println(uid)

	if err != nil {
		tool.RespErrorWithData(ctx, "解析1失败")
		fmt.Println("putMovieShortComment is ERR", err)
		return
	}
	MovieType := ctx.PostForm("type")

	star, err := strconv.Atoi(ctx.PostForm("star"))

	if err != nil {
		fmt.Println("putMovieShortComment is ERR", err)
		tool.RespErrorWithData(ctx, "解析2失败")
		return
	}
	//想看or看过

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
		Static:  MovieType,
	}

	err = service.PutMovieShortComment(ShorComment)
	if err != nil {
		tool.RespInternalError(ctx)
		return
	}

	err = service.UpdateSubjectScore(movieId, star)
	if err != nil {
		fmt.Println("ChangeMovieScore Is err", err)
		tool.RespInternalError(ctx)
		return
	}

	err = service.InsertMovieStatic(MovieType, uid, movieId)
	if err != nil {
		fmt.Println("InsertMovieStatic Is err", err)
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

	ctx.JSON(200, gin.H{
		"short_comment": newShortCommentSlice,
	})

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

	err = service.UpdateSubjectScore(movieId, star)
	if err != nil {
		fmt.Println("ChangeMovieScore Is err", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, "评论成功")
}

//获取影评
func getLargeComment(ctx *gin.Context) {
	Mid, err := strconv.ParseInt(ctx.Query("mid"), 10, 64)
	if err != nil {
		tool.RespErrorWithData(ctx, "解析失败")
		return
	}
	commentSlice, err := service.GetLargeCommentSlice(Mid)

	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("getLargeComment_GetLargeCommentSlice Err is", err)
		return
	}

	if commentSlice == nil {
		commentSlice = []model.LargeComment{}
	}

	var newLargeCommentSlice []param.LargeComment

	for _, commentModel := range commentSlice {
		var largeCommentParam param.LargeComment
		user, _ := service.GetUserById(commentModel.Uid)

		largeCommentParam.Time = commentModel.Time.Format("2006-01-02 15:04:05")
		largeCommentParam.Id = commentModel.Id
		largeCommentParam.Comment = commentModel.Comment
		largeCommentParam.Title = commentModel.Title
		largeCommentParam.User = user
		largeCommentParam.Likes = commentModel.Likes
		largeCommentParam.Unlikes = commentModel.Unlikes
		largeCommentParam.MId = commentModel.Mid

		newLargeCommentSlice = append(newLargeCommentSlice, largeCommentParam)
	}

	ctx.JSON(200, gin.H{
		"large_comment": newLargeCommentSlice,
	})
}
