package dao

import (
	"context"
	"douban/tool"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

var redisConn = tool.RedisClient

func GetRedisValue(ctx *gin.Context, key string) (string, error) {
	GetKey := redisConn.Get(ctx, key)
	if GetKey.Err() != nil {
		return "", GetKey.Err()
	}

	return GetKey.Val(), nil
}

//不管是邮箱还是电话号码都是同5minute
func SetRedisValue(ctx *gin.Context, key string, value string) error {
	SetKV := redisConn.Set(ctx, key, value, time.Minute*2)
	return SetKV.Err()
}

//用于页面缓存

type RedisStore struct {
	PreKey     string        //页面类型key
	Expiration time.Duration //过期时间
	Context    context.Context
}

func NewRedisStore(preKey string, expire time.Duration, context context.Context) *RedisStore {
	return &RedisStore{
		PreKey:     preKey,
		Expiration: expire,
		Context:    context,
	}
}

func (rs *RedisStore) SetRedisPages(key string, val interface{}) error {
	bytes, _ := json.Marshal(val)
	SetKV := redisConn.Set(rs.Context, rs.PreKey, bytes, rs.Expiration)
	return SetKV.Err()
}

func (rs *RedisStore) GetRedisPages(key string, obj interface{}) bool {
	val, err := redisConn.Get(rs.Context, key).Result()
	if err != nil {
		return false
	} else {
		json.Unmarshal([]byte(val), &obj)
		return true
	}
}
