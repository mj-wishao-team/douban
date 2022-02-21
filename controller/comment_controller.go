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

type CommentController struct {
}

func (C *CommentController) Router(engine *gin.Engine) {
	engine.POST("api/movie/comment/:mid", JWTAuthMiddleware(), putMovieShortComment)
	engine.POST("api/movie/review/put", JWTAuthMiddleware(), putMovieLargeComment)

	engine.GET("api/movie/comment/:mid", getShortComment)
	engine.GET("api/movie/review/:mid", getLargeComment)

	engine.PUT("api/movie/review/like/:mid", updateLikeReview)
	engine.PUT("api/movie/comment/like/:mid", updateLikeComment)
}

//短评点赞
func updateLikeComment(ctx *gin.Context) {
	CommentId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		tool.RespErrorWithData(ctx, "解析1失败")
		fmt.Println("updateLikeComment is ERR", err)
		return
	}
	like, err := strconv.Atoi(ctx.PostForm("like"))
	if err != nil {
		tool.RespErrorWithData(ctx, "解析1失败")
		fmt.Println("updateLikeComment is ERR", err)
		return
	}

	err = service.UpdateCommentLike(CommentId, like)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println(err)
		return
	}
	tool.RespSuccessfulWithData(ctx, "点赞成功")
}

//影评点赞
func updateLikeReview(ctx *gin.Context) {
	ReviewId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		tool.RespErrorWithData(ctx, "解析1失败")
		fmt.Println("updateLikeReview is ERR", err)
		return
	}
	like, err := strconv.Atoi(ctx.PostForm("like"))
	if err != nil {
		tool.RespErrorWithData(ctx, "解析1失败")
		fmt.Println(" updateLikeReview is ERR", err)
		return
	}
	err = service.UpdateReviewLike(ReviewId, like)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println(err)
		return
	}
	tool.RespSuccessfulWithData(ctx, "点赞成功")
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

	shortComment := html.UnescapeString(html.EscapeString(ctx.PostForm("comment")))

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
	tool.RespSuccessfulWithData(ctx, commentSlice)

}

//发表影评
func putMovieLargeComment(ctx *gin.Context) {
	movieId, err := strconv.ParseInt(ctx.PostForm("mid"), 10, 64)

	title := ctx.PostForm("title")

	uid := ctx.MustGet("id").(int64)

	star, err := strconv.Atoi(ctx.PostForm("star"))

	if err != nil {
		tool.RespErrorWithData(ctx, "解析失败")
		fmt.Println("putMovieShortComment is ERR", err)
		return
	}

	//防止xxs 注入
	largeComment := html.UnescapeString(html.EscapeString(ctx.PostForm("comment")))

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
		fmt.Println(err)
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
	Mid, err := strconv.ParseInt(ctx.Param("mid"), 10, 64)
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

	tool.RespSuccessfulWithData(ctx, commentSlice)
}
