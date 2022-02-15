package service

import (
	"douban/dao"
	"douban/model"
)

func GetCelebrity(id int64) ([]model.Celebrity, error) {
	celebrity, err := dao.GetCelebrity(id)
	return celebrity, err
}
