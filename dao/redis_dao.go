package dao

import (
	"github.com/gin-gonic/gin"
	"movie.douban/tool"
	"time"
)

func GetRedisValue(ctx *gin.Context, key string) (string, error) {
	redisConn := tool.RedisClient
	GetKey := redisConn.Get(ctx, key)
	if GetKey.Err() != nil {
		return "", GetKey.Err()
	}

	return GetKey.Val(), nil
}

//不管是邮箱还是电话号码都是同5minute
func SetRedisValue(ctx *gin.Context, key string, value string) error {
	redisConn := tool.RedisClient
	SetKV := redisConn.Set(ctx, key, value, time.Minute*5)
	return SetKV.Err()
}
