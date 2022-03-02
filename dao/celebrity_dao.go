package dao

import (
	"douban/model"
)

func GetCelebrity(id int64) ([]model.Celebrity, error) {
	var Celebrities []model.Celebrity
	rs := Db.Where("id=?", id).Find(&Celebrities)
	return Celebrities, rs.Error

}

// 搜索 影人信息
func GetCelebrityByKeyWord(word string) (celebrities []model.Celebrity, err error) {
	var Celebrities []model.Celebrity
	rs := Db.Where("name LIKE '%?%' ", word).Or("english_name LIKE '%?%'", word).Find(&Celebrities)
	return Celebrities, rs.Error

}

//func GetCelebrityByKeyWord(word string) (celebrities []model.Celebrity, err error) {
//	sqlStr := "SELECT * FROM celebrity WHERE name LIKE '%?%' OR english_name LIKE '%?%'"
//	Stmt, err := DB.Prepare(sqlStr)
//
//	rows, err := Stmt.Query(word, word)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var celebrity model.Celebrity
//		err = rows.Scan(
//			&celebrity.Id,
//			&celebrity.Name,
//			&celebrity.Avatar,
//			&celebrity.EnglishName,
//			&celebrity.Gender,
//			&celebrity.Sign,
//			&celebrity.Birth,
//			&celebrity.Hometown,
//			&celebrity.Job,
//			&celebrity.IMDb,
//			&celebrity.Brief,
//		)
//		if err != nil {
//			return nil, err
//		}
//		celebrities = append(celebrities, celebrity)
//	}
//	return celebrities, err
//
//}
