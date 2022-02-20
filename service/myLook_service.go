package service

import (
	"douban/dao"
	"douban/model"
)

func GetMyLook(uid int64) ([]model.MovieStatic, error) {
	MS, err := dao.GetMyLook(uid)

	return MS, err
}
