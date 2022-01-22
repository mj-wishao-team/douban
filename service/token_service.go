package service

import (
	"douban/model"
	"douban/tool"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//生成一个jwt
func GenToken(user model.User, ExpireTime int64, tokenType string) (string, error) {
	JwtCfg := tool.GetCfg().Jwt

	c := model.MyClaims{
		User: user,
		Type: tokenType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + ExpireTime,
			Issuer:    "douban",
		},
	}

	if tokenType == "ACCESS_TOKEN" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
		return token.SignedString([]byte(JwtCfg.AccessSecret))
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
		return token.SignedString([]byte(JwtCfg.RefreshSecret))
	}

}

//解析双token
func ParseToken(accessTokenString, refreshTokenString string) (*model.MyClaims, bool, error) {
	JwtCfg := tool.GetCfg().Jwt
	accessToken, err := jwt.ParseWithClaims(accessTokenString, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtCfg.AccessSecret), nil
	})

	//access_token 没有过期
	if claims, ok := accessToken.Claims.(*model.MyClaims); ok && accessToken.Valid {
		return claims, false, nil
	}

	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtCfg.RefreshSecret), nil
	})

	//若果refresh_token 也过期了者要重新登录
	if err != nil {
		return nil, false, err
	}
	//若果access_token过期判断refresh_token
	if claims, ok := refreshToken.Claims.(*model.MyClaims); ok && refreshToken.Valid {
		return claims, true, nil
	}

	return nil, false, errors.New("invalid token")
}
