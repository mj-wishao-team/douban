package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	User User
	Type string // REFRESH_TOKEN and TOKEN 用于更新token的标识
	Time time.Time
	jwt.StandardClaims
}
