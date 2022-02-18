package service

import (
	"douban/dao"
	"douban/model"
)

//发表讨论
func PutDiscussion(discussion model.Discussion) error {
	err := dao.InsertDiscussion(discussion)
	return err
}

//获取讨论列表
func GetDiscussionList(sort string, mid int64) ([]model.DiscussionList, error) {
	discussionList, err := dao.GetDiscussionList(orderWay[sort], mid)
	return discussionList, err
}

//获取讨论
func GetDiscussion(id int64) ([]model.Discussion, error) {
	Discussion, err := dao.GetDiscussion(id)
	return Discussion, err
}

//删除讨论
func DeleteDisucuss(id int64) error {
	err := dao.DeleteDiscussion(id)
	return err
}
