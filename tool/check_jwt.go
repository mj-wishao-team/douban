package tool

import (
	"douban/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

//检查token

func CheckToken(ctx *gin.Context, claims *model.MyClaims, err error) bool {

	if err == nil && claims.Type == "ERR" {
		fmt.Println("TokenTypeIsERR:", err)
		RespErrorWithData(ctx, "PARSE_TOKEN_ERROR")
		return false
	}

	if err != nil {
		if err.Error()[:16] == "token is expired" {
			fmt.Println("err is :", err)
			RespErrorWithData(ctx, "TOKEN_EXPIRED")
			return false
		}

		fmt.Println("ParesTokenERR:", err)
		RespErrorWithData(ctx, "PARSE_TOKEN_ERROR")
		return false
	}

	return true
}
