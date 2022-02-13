package model

import (
	"time"
)

type ShortComment struct {
	Id      int64
	Mid     int64
	Uid     int64
	Comment string
	Time    time.Time
	Help    int64
	Report  int64
	Star    int
}

type LargeComment struct {
	Id      int64
	Mid     int64
	Uid     int64
	Title   string
	Comment string
	Time    time.Time
	likes   int64
	unlikes int64
	Report  int64
	Star    int
}
