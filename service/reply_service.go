package service

import (
	"douban/dao"
	"douban/model"
)

func GetReply(id int64, kind string, start int) ([]model.Reply, error) {
	Rp, err := dao.GetReply(id, kind, start)
	return Rp, err
}

func ReplyPost(Reply model.Reply) error {
	err := dao.ReplyPost(Reply)
	return err
}
