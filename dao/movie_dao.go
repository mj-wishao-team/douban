package dao

import (
	"douban/model"
)

func GetMovieById(id int64) ([]model.Movie, error) {
	var movies []model.Movie
	var movie model.Movie

	sqlStr := "SELECT id,name,poster,director,screenwriter,starring,type,tag,country,language,release_time,duration,alias,imdb,age, score,peoples,one_star,two_star,three_star,four_star,five_star FROM movie WHERE id= ? "
	Stmt, err := DB.Prepare(sqlStr)

	defer Stmt.Close()
	if err != nil {
		return nil, err
	}

	row := Stmt.QueryRow(id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err = row.Scan(
		&movie.Id,
		&movie.Name,
		&movie.Poster,
		&movie.Director,
		&movie.ScreenWriter,
		&movie.Starring,
		&movie.Type,
		&movie.Tag,
		&movie.Country,
		&movie.Language,
		&movie.ReleaseTime,
		&movie.Duration,
		&movie.Alias,
		&movie.Imdb,
		&movie.Age,
		&movie.Score,
		&movie.Peoples,
		&movie.OneStar,
		&movie.TwoStar,
		&movie.ThreeStar,
		&movie.FourStar,
		&movie.FiveStar,
	)
	if err != nil {
		return nil, err
	}

	movies = append(movies, movie)
	return movies, err
}

//查询电影评分
func QueryByMovie(id int64) (float64, float64, error) {
	var Score, Number float64

	sqlStr := "SELECT score,peoples FROM movie WHERE id= ? "
	Stmt, err := DB.Prepare(sqlStr)
	defer Stmt.Close()

	if err != nil {
		return 0, 0, err
	}

	row := Stmt.QueryRow(id)
	if row.Err() != nil {
		return 0, 0, row.Err()
	}

	err = row.Scan(&Score, &Number)
	if err != nil {
		return 0, 0, err
	}

	return Score, Number, nil
}

//电影排行榜
func GetMovieLeaderboard(limit int) ([]model.MovieList, error) {
	var MovieLists []model.MovieList

	sqlStr := "SELECT id,name, poster, score FROM movie ORDER BY score DESC LIMIT ?;"

	Stmt, err := DB.Prepare(sqlStr)

	defer Stmt.Close()

	if err != nil {
		return nil, err
	}

	rows, err := Stmt.Query(limit)

	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	for rows.Next() {
		var MovieList model.MovieList
		err = rows.Scan(&MovieList.Id, &MovieList.Name, &MovieList.Poster, &MovieList.Score)
		if err != nil {
			return nil, err
		}

		MovieLists = append(MovieLists, MovieList)
	}

	if err != nil {
		return nil, err
	}

	return MovieLists, nil
}

//选电影
func GetSortMovieByTags(tag string, sortWay string, limit int) ([]model.MovieList, error) {
	var MovieLists []model.MovieList

	sqlStr := "SELECT id,name, poster, score FROM movie WHERE tags LIKE '%?%'"
	sqlStr = sqlStr + "ORDER BY " + sortWay + "limit ?"
	Stmt, err := DB.Prepare(sqlStr)

	defer Stmt.Close()

	if err != nil {
		return nil, err
	}

	rows, err := Stmt.Query(tag, limit)

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	for rows.Next() {
		var MovieList model.MovieList
		err = rows.Scan(&MovieList.Id, &MovieList.Name, &MovieList.Poster, &MovieList.Score)
		if err != nil {
			return nil, err
		}

		MovieLists = append(MovieLists, MovieList)
	}

	if err != nil {
		return nil, err
	}

	return MovieLists, nil
}

//跟新电影评分
func UpdateMovieScore(Score float64, id int64) error {
	sqlStr := "update  movie  set  score=? where id = ? ;"
	Stmt, err := DB.Prepare(sqlStr)
	_, err = Stmt.Exec(Score, id)
	return err

}

//增加电影评论人数
func IncreaseMoviePeople(id int64) error {
	sqlStr := "update  movie  set  peoples=peoples+ 1  where id = ? ;"
	Stmt, err := DB.Prepare(sqlStr)
	_, err = Stmt.Exec(id)
	return err
}

//评价
func IncreaseOneStar(id int64) error {
	sqlStr := "update  movie  set  one_star=one_star+ 1  where id = ? ;"
	Stmt, err := DB.Prepare(sqlStr)
	_, err = Stmt.Exec(id)
	return err
}

func IncreaseTwoStar(id int64) error {
	sqlStr := "update  movie  set  two_star=two_star+ 1  where id = ? ;"
	Stmt, err := DB.Prepare(sqlStr)
	_, err = Stmt.Exec(id)
	return err
}

func IncreaseThreeStar(id int64) error {
	sqlStr := "update  movie  set  three_star=three_star+ 1  where id = ? ;"
	Stmt, err := DB.Prepare(sqlStr)
	_, err = Stmt.Exec(id)
	return err
}

func IncreaseFourStar(id int64) error {
	sqlStr := "update  movie  set  four_star=four_star+ 1  where id = ? ;"
	Stmt, err := DB.Prepare(sqlStr)
	_, err = Stmt.Exec(id)
	return err
}

func IncreaseFiveStar(id int64) error {
	sqlStr := "update  movie  set  five_star=five_star+ 1  where id = ? ;"
	Stmt, err := DB.Prepare(sqlStr)
	_, err = Stmt.Exec(id)
	return err
}
