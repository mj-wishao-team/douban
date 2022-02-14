package service

import (
	"douban/dao"
	"douban/model"
	"fmt"
	"strconv"
)

//获取单个电影信息
func GetMovieById(id int64) ([]model.Movie, error) {
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
	"latest": "release_time DESC",
	"hotest": "score DESC",
}

//选电影
func GetMovieListByTag(tag string, sort string, limit int) ([]model.MovieList, error) {
	ML, err := dao.GetSortMovieByTags(tag, orderWay[sort], limit)
	return ML, err
}

//影片评价
func ChangeMovieScoreById(id int64, star int) error {

	number, score, err := dao.QueryByMovie(id)
	starStr := strconv.Itoa(star)
	Star, err := strconv.ParseFloat(starStr, 64)
	NewScore := (score*number + Star) / (number + 1)

	err = dao.UpdateMovieScore(NewScore, id)
	if err != nil {
		fmt.Println("Score is ERR ", err)
		return err
	}

	err = dao.IncreaseMoviePeople(id)
	if err != nil {
		fmt.Println("People is ERR", err)
		return err
	}

	switch star {
	case 1:
		err = dao.IncreaseOneStar(id)
	case 2:
		err = dao.IncreaseTwoStar(id)
	case 3:
		err = dao.IncreaseThreeStar(id)
	case 4:
		err = dao.IncreaseFourStar(id)
	case 5:
		err = dao.IncreaseFiveStar(id)
	}
	if err != nil {
		fmt.Println("Star is err", err)
		return err
	}

	return err
}
