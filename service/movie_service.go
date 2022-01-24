package service

import (
	"douban/dao"
	"douban/model"
)

func GetMovieById(id int64) ([]model.Movie, error) {
	movies, err := dao.GetMovieById(id)
	return movies, err
}

func PutMovieShortComment(shortComment string, Mid int64) error {
	dao.InsertShortComment(shortComment, Mid)
}
