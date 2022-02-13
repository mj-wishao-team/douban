package param

import "douban/model"

type ShortComment struct {
	Id      int64
	User    model.UserInfo
	MId     int64
	Comment string
	Time    string
	Help    int64
	Report  int64 //投诉
	Star    int
}

type LargeComment struct {
	Id      int64
	User    model.UserInfo
	Title   string
	MId     int64
	Comment string
	Time    string
	Likes   int64
	Unlikes int64
	Report  int64
	Star    int
}
