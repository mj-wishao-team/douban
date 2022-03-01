package service

import (
	"douban/dao"
	"douban/model"
	"fmt"
	"strconv"
	"strings"
)

//插入电影的类别(想看OR看过)
func InsertMovieStatic(MovieType string, uid int64, mid int64) error {
	err := dao.InsertMovieStatic(MovieType, uid, mid)
	return err
}

//搜索电影
func SearchMovies(word string) ([]model.MovieList, error) {
	Movies, err := dao.SearchMovies(word)
	return Movies, err
}

//获取单个电影信息
func GetMovieById(id int64) (model.Movie, error) {
	movies, err := dao.GetMovieById(id)
	return movies, err
}

//判断Mid是否存在，存在则返回true
func JudgeMovie(Mid int64) (bool, error) {

	_, err := dao.GetMovieById(Mid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

var orderWay = map[string]string{
	"latest": "date DESC",
	"hotest": "stars DESC",
	"host":   "people DESC",
	"time":   "time DESC",
}

//选电影
func GetMovieListByTag(tag string, sort string, start int) ([]model.MovieList, error) {
	ML, err := dao.SelectSubjectsByTag(tag, orderWay[sort], start)
	return ML, err
}

//电影排行榜
func GetMovieLeaderboard(start int, rp *dao.RedisStore) (MovieList []model.MovieList, err error) {

	if flag := rp.GetRedisPages(rp.PreKey, &MovieList); flag == true {
		return
	}

	MovieList, err = dao.GetMovieLeaderboard(start)

	if err := rp.SetRedisPages(rp.PreKey, &MovieList); err != nil {
		return nil, err
	}
	return
}

//影片评价
func UpdateSubjectScore(mid int64, score int) (err error) {

	movie, err := GetMovieById(mid)
	if err != nil {
		fmt.Println("001")
		return err
	}
	Score, err := strconv.ParseFloat(movie.Score.Score, 64)
	if err != nil {
		fmt.Println("002")
		return err
	}
	Totalcnt := float64(movie.Score.TotalCnt)
	NewScore := (Totalcnt*Score + float64(score)*2) / (Totalcnt + 1)
	NewScoreStr := strconv.FormatFloat(NewScore, 'f', 2, 64)

	var NewMovieScore = model.MovieScore{
		Score:    NewScoreStr,
		TotalCnt: int(Totalcnt + 1),
		Five:     movie.Score.Five,
		Four:     movie.Score.Four,
		Three:    movie.Score.Three,
		Two:      movie.Score.Two,
		One:      movie.Score.One,
	}

	switch score {
	case 1:
		NewMovieScore.One, err = parsePctToNewPct(movie.Score.One, Totalcnt)
		if err != nil {
			fmt.Println("003")
			return err
		}
	case 2:
		NewMovieScore.Two, err = parsePctToNewPct(movie.Score.Two, Totalcnt)
		if err != nil {
			return err
		}
	case 3:
		NewMovieScore.Three, err = parsePctToNewPct(movie.Score.Three, Totalcnt)
		if err != nil {
			return err
		}
	case 4:
		NewMovieScore.Four, err = parsePctToNewPct(movie.Score.Four, Totalcnt)

		if err != nil {
			fmt.Println("004")
			return err
		}
	case 5:
		NewMovieScore.Five, err = parsePctToNewPct(movie.Score.Five, Totalcnt)
		if err != nil {
			return err
		}
	}

	err = dao.UpdateSubjectScore(mid, NewMovieScore)

	return
}

func parsePctToNewPct(v string, Totalcnt float64) (per string, err error) {
	v = strings.Replace(v, "%", "", -1)
	ret, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return v, err
	}
	return (strings.TrimLeft(strconv.FormatFloat((ret*Totalcnt+100)/(Totalcnt+1)*100, 'f', 2, 64), "0.") + "%"), err
}
