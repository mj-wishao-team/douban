package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"movie.douban/service"
	"movie.douban/tool"
	"net/http"
)

//当access_token失效时，使用refresh_token来请求刷新access_token。
//如果refreshToken过期，需要用户重新登，并且重新刷新accessToken和refreshToken
//若忘：过程图详见https://blog.csdn.net/qq_33600019/article/details/80855759

//通过refresh_token获取access_token

func getAccessToken(ctx *gin.Context) {
	refreshToken := ctx.Query("refreshToken")

	//判断refreshToken状态
	Claims, err := service.ParseRefreshToken(refreshToken)
	if err != nil {
		if err.Error()[:16] == "token is expired" {
			tool.RespErrorWithData(ctx, "refreshToke 过期了")
			return
		}

		fmt.Println("getAccessToken_ParseRefreshToken is Err:", err)
		tool.RespErrorWithData(ctx, "refreshToken不正确或系统错误")
		return
	}

	if Claims.Type == "ERR" {
		tool.RespErrorWithData(ctx, "token is ERR")
		return
	}

	//根据id更新user
	newUser, _, err := service.JudgeAndQueryUserByUserID(Claims.User.Id)

	//创建新access_token
	newAccessToken, err := service.GenToken(newUser, 120, "TOKEN")
	if err != nil {
		fmt.Println("GenToken_getTokenCreateErr:", err)
		tool.RespInternalError(ctx)
		return
	}

	//返回access_token
	ctx.JSON(http.StatusOK, gin.H{
		"status":       "ture",
		"access_token": newAccessToken,
	})
}
