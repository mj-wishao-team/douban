package service

import (
	"douban/dao"
	"douban/model"
)

//增加影评回复人数
func UpdateReviewCNT(id int64) error {
	err := dao.UpdateReviewCNT(id)
	return err
}

//增加讨论回复人数
func UpdateDiscussionCNT(id int64) error {
	err := dao.UpdateDiscussionCNT(id)
	return err

}

//增加回复的回复人数
func UpdateReplyCNT(id int64) error {
	err := dao.UpdateReplyCNT(id)
	return err

}

//短评点赞
func UpdateCommentLike(id int64, like int) error {
	err := dao.UpdateCommentLike(id, like)
	return err
}

//影评点赞
func UpdateReviewLike(id int64, like int) error {
	err := dao.UpdateReviewLike(id, like)
	return err
}

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

//获取单个影评
func GetReview(id int64) ([]model.LargeComment, error) {
	LC, err := dao.GetReview(id)
	return LC, err

}

//获取自己的影评
func GetLargeCommentByUid(Uid int64) ([]model.LargeComment, error) {
	Comment, err := dao.QueryLargeCommentByUid(Uid)
	return Comment, err
}
