package service

import (
	"douban/dao"
	"douban/model"
)

//发布短评
func PutMovieShortComment(shortComment model.ShortComment) error {
	err := dao.InsertShortComment(shortComment)
	return err
}

//发表影评
func PutMovieLargeComment(Comment model.LargeComment) error {
	err := dao.InsertLargeComment(Comment)
	return err
}
