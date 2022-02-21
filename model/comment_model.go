package model

import (
	"time"
)

type ShortComment struct {
	Id       int64
	Mid      int64
	Uid      int64
	Username string
	Avatar   string
	Static   string
	Comment  string
	Time     time.Time
	Help     int64
	Report   int64
	Star     int
}

type LargeComment struct {
	Id       int64
	Mid      int64
	Uid      int64
	Username string
	Avatar   string
	Title    string
	Comment  string
	Time     time.Time
	People   int
	Likes    int64
	Unlikes  int64
	Report   int64
	Star     int
}
