package model

import "time"

type UserInfo struct {
	Id         int64
	Username   string
	Email      string
	Phone      string
	Avatar     string
	DomainName string
	Habitat    string //栖息地
	Hometown   string //家乡
	Birthday   time.Time
	RegDate    time.Time
	Statement  string //个人称述
	Followers  int64
	Followings int64
}
