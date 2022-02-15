package dao

import "douban/model"

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
