package dao

import "douban/model"

func GetMovieById(id int64) ([]model.Movie, error) {
	var movies []model.Movie
	var movie model.Movie

	sqlStr := "SELECT id,name,data,star,type ,release_time,picture,language,country,length,statement,followers,followings FROM movie WHERE id= ? "
	Stmt, err := DB.Prepare(sqlStr)
	defer Stmt.Close()

	if err != nil {
		return nil, err
	}

	row := Stmt.QueryRow(id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err = row.Scan(&movie.Id)
	if err != nil {
		return nil, err
	}

	movies = append(movies, movie)
	return movies, err
}
