package dao

import (
	"douban/model"
)

func GetCelebrity(id int64) ([]model.Celebrity, error) {
	var Celebrities []model.Celebrity
	sqlStr := "SELECT * FROM celebrity Where id=?;"
	Stmt, err := DB.Prepare(sqlStr)

	rows, err := Stmt.Query(id)
	defer Stmt.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var Celebrity model.Celebrity
		err := rows.Scan(&Celebrity.Id, &Celebrity.Name,
			&Celebrity.Avatar,
			&Celebrity.EnglishName, &Celebrity.Gender,
			&Celebrity.Sign, &Celebrity.Birth,
			&Celebrity.Hometown, &Celebrity.Job,
			&Celebrity.IMDb, &Celebrity.Brief)

		if err != nil {
			return nil, err
		}
		Celebrities = append(Celebrities, Celebrity)
	}
	return Celebrities, err
}

// 搜索 影人信息
func GetCelebrityByKeyWord(word string) (celebrities []model.Celebrity, err error) {
	sqlStr := "SELECT * FROM celebrity WHERE name LIKE '%?%' OR english_name LIKE '%?%'"
	Stmt, err := DB.Prepare(sqlStr)

	rows, err := Stmt.Query(word, word)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var celebrity model.Celebrity
		err = rows.Scan(
			&celebrity.Id,
			&celebrity.Name,
			&celebrity.Avatar,
			&celebrity.EnglishName,
			&celebrity.Gender,
			&celebrity.Sign,
			&celebrity.Birth,
			&celebrity.Hometown,
			&celebrity.Job,
			&celebrity.IMDb,
			&celebrity.Brief,
		)
		if err != nil {
			return nil, err
		}
		celebrities = append(celebrities, celebrity)
	}
	return celebrities, err

}
