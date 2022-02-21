package model

import "time"

type MovieTag struct {
	Name   string
	Score  string
	Avatar string
	Mid    int64
}

type Movie struct {
	Mid         int64
	Name        string
	Stars       int64
	Date        time.Time
	Tags        string
	Avatar      string
	Detail      MovieDetail
	Score       MovieScore
	Plot        string
	Celebrities []int64
}

type MovieScore struct {
	Score    string
	TotalCnt int
	Five     string
	Four     string
	Three    string
	Two      string
	One      string
}

type MovieDetail struct {
	Nicknames  []string
	Director   string
	Writers    []string
	Characters []string
	Type       []string
	Website    string
	Region     string
	Language   string
	Release    string
	Period     int
	IMDb       string
}

type MovieList struct {
	Mid    int64       `json:"mid"`
	Name   string      `json:"name"`
	Tags   string      `json:"tags"`
	Avatar string      `json:"avatar"`
	Detail MovieDetail `json:"detail"`
	Score  MovieScore  `json:"score"`
}

type MovieStatic struct {
	Mid         int64
	Uid         int64
	Type        string
	MovieName   string
	MovieAvatar string
}
