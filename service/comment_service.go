package service

import (
	"douban/dao"
	"douban/model"
)

//获取短评
func GetShortCommentByUidAndMid(uid, mid int64) ([]model.ShortComment, error) {
	SC, err := dao.GetShortCommentByUidAndMid(uid, mid)
	return SC, err
}

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

//获取短评
func GetShortCommentSlice(mid int64) ([]model.ShortComment, error) {
	commentSlice, err := dao.QueryShortCommentByMid(mid)
	return commentSlice, err
}

func GetMovieComment(mid int64) ([]model.ShortComment, error) {
	commentSlice, err := dao.GetMovieComment(mid)
	return commentSlice, err
}

//获取影评
func GetLargeCommentSlice(mid int64) ([]model.LargeComment, error) {
	commentSlice, err := dao.QueryLargeCommentByMid(mid)
	return commentSlice, err
}

func GetMovieReviews(mid int64) ([]model.LargeComment, error) {
	commentSlice, err := dao.GetMovieReviews(mid)
	return commentSlice, err
}

//获取自己的影评
func GetLargeCommentByUid(Uid int64) ([]model.LargeComment, error) {
	Comment, err := dao.QueryLargeCommentByUid(Uid)
	if err.Error() == "sql: no rows in result set" {
		return nil, nil
	}
	return Comment, err
}
