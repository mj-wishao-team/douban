package service

import (
	"github.com/dgrijalva/jwt-go"
	"movie.douban/model"
	"movie.douban/tool"
	"time"
)

//解析AccessToken
func ParseAccessToken(tokenString string) (*model.MyClaims, error) {
	JwtCfg := tool.GetCfg().Jwt
	MySecret := []byte(JwtCfg.MySecret)
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	// 校验token并且判断token类型
	if clams, ok := token.Claims.(*model.MyClaims); ok && token.Valid {
		if clams.Type == "REFRESH_TOKEN" {
			//如果是 "REFRESH_TOKEN"则将token 空返回并将Type 改为ERR
			errClaims := new(model.MyClaims)
			errClaims.Type = "ERR"
			return errClaims, nil
		}
		return clams, nil
	} else {
		return nil, err
	}
}

//解析RefreshToken
func ParseRefreshToken(tokenString string) (*model.MyClaims, error) {
	JwtCfg := tool.GetCfg().Jwt
	MySecret := []byte(JwtCfg.MySecret)
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})

	if clams, ok := token.Claims.(*model.MyClaims); ok && token.Valid {
		if clams.Type == "TOKEN" {
			errClaims := new(model.MyClaims)
			errClaims.Type = "ERR"
			return errClaims, nil
		}
		return clams, nil
	} else {
		return nil, err
	}
}

//生成一个jwt
func GenToken(user model.User, ExpireTime int64, tokenType string) (string, error) {
	JwtCfg := tool.GetCfg().Jwt
	MySecret := []byte(JwtCfg.MySecret)

	c := model.MyClaims{
		User: user,
		Type: tokenType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + ExpireTime,
			Issuer:    "douban",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(MySecret)
}
