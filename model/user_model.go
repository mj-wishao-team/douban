package model

import "time"

type User struct {
	Id         int64
	Username   string
	Password   string
	Email      string
	Phone      string
	Salt       string
	Avatar     string
	DomainName string
	Habitat    string
	Hometown   string
	Birthday   time.Time
	RegDate    time.Time
	Statement  string
	Followers  int64
	Followings int64
}
