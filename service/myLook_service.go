package service

import (
	"douban/dao"
	"douban/model"
)

func GetMyLook(uid int64, str string) ([]model.MovieStatic, error) {
	MS, err := dao.GetMyLook(uid, str)
	return MS, err
}
