package tool

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CatchPanic(ctx *gin.Context, errWhere string) {
	defer func() {
		if err := recover(); err != nil {
			RespInternalError(ctx)
			fmt.Println(errWhere+" panic is", err)
			return
		}
	}()
}
