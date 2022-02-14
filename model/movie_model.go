package model

type Movie struct {
	Id           int64
	Name         string
	Poster       string
	Director     string
	ScreenWriter string
	Starring     string
	Type         string
	Tag          string
	Country      string
	Language     string
	ReleaseTime  string
	Duration     string
	Alias        string
	Imdb         string
	Score        string
	Age          string
	Peoples      int
	OneStar      int
	TwoStar      int
	ThreeStar    int
	FourStar     int
	FiveStar     int
}

type MovieList struct {
	Id     int64
	Name   string
	Poster string
	Score  string
}
