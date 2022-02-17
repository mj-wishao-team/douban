package service

import (
	"douban/dao"
	"douban/model"
)

func PutDiscussion(discussion model.Discussion) error {
	err := dao.InsertDiscussion(discussion)
	return err
}
