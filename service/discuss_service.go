package service

import (
	"douban/dao"
	"douban/model"
)

func PutDiscussion(discussion model.Discussion) error {
	err := dao.InsertDiscussion(discussion)
	return err
}

func GetDiscussion(sort string, mid int64) ([]model.DiscussionList, error) {
	discussionList, err := dao.GetDiscussionList(orderWay[sort], mid)
	return discussionList, err
}
