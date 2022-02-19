package controller

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchController struct {
}

func (D *SearchController) Router(engine *gin.Engine) {
	engine.GET("/api/movie/search", search)
}

func search(ctx *gin.Context) {
	words := ctx.Query("key")
	if words == "" {
		tool.RespErrorWithData(ctx, "key 不能为空")
		return
	}
	Celebrities, err := service.SearchCelebrity(words)

	if err != nil && err.Error() == "sql: no rows in result set" {
		tool.RespInternalError(ctx)
		fmt.Println(" Search Celebrity_err is", err)
		return
	}

	Movies, err := service.SearchMovies(words)

	if err != nil && err.Error() == "sql: no rows in result set" {
		tool.RespInternalError(ctx)
		fmt.Println(" Search Movie_err is", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"movie":     Movies,
		"celebrity": Celebrities,
	})

}
