package controller

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		var headerKeys []string
		for key, _ := range context.Request.Header {
			headerKeys = append(headerKeys, key)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		if origin != "" {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			context.Header("Access-Control-Max-Age", "172800")
			context.Header("Access-Control-Allow-Credentials", "ture")
			context.Set("content-type", "application/json") //// 设置返回格式是json
		}

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}

		//处理请求
		context.Next()
	}
}

//当access_token失效时，使用refresh_token来请求刷新access_token和refresh_token。
//如果refreshToken过期，需要用户重新登，并且重新刷新accessToken和refreshToken

func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 默认双Token放在请求头Authorization的Bearer中，并以空格隔开
		authHeader := ctx.Request.Header.Get("Authorization")
		fmt.Println(ctx.Request.Header)
		if authHeader == "" {
			tool.RespErrorWithData(ctx, "请求头为空")
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")

		if !(len(parts) == 3 && parts[0] == "Bearer") {
			tool.RespErrorWithData(ctx, "请求头内token格式有误")
			ctx.Abort()
			return
		}

		Clams, flag, err := service.ParseToken(parts[1], parts[2])
		if err != nil {
			fmt.Println("ParesTokenERR:", err)
			tool.RespErrorWithData(ctx, "你需要重新登录")
			ctx.Abort()
			return
		}
		// accessToken 已经失效，需要刷新双Token
		if flag {

			accessToken, err := service.GenToken(Clams.User, 300, "ACCESS_TOKEN")
			if err != nil {
				fmt.Println("JWTAuthMiddleware_CreateAccessTokenErr:", err)
				tool.RespInternalError(ctx)
				ctx.Abort()
				return
			}

			//refreshToken 一周
			refreshToken, err := service.GenToken(Clams.User, 604800, "REFRESH_TOKEN")
			if err != nil {
				fmt.Println("JWTAuthMiddleware_CreateRefreshTokenErr:", err)
				tool.RespInternalError(ctx)
				ctx.Abort()
				return
			}

			fmt.Println("token 更新")
			fmt.Println(accessToken + " " + refreshToken)

			// 如果需要刷新双Token时，返回双Token
			ctx.JSON(http.StatusOK, gin.H{
				"data":          "Token Refresh Success",
				"access_token":  accessToken,
				"refresh_token": refreshToken,
				"token":         accessToken + " " + refreshToken,
			})
			return
		}

		ctx.Set("id", Clams.User.Id)
		ctx.Set("username", Clams.User.Username)
		ctx.Next()
	}

}
