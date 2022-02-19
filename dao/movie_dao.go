package dao

import (
	"douban/model"
	"encoding/json"
	"strings"
)

//电影排行榜
func GetMovieLeaderboard(start int) (movieLists []model.MovieList, err error) {
	var movieList model.MovieList
	var Detail, Score string
	sqlStr := "SELECT name, score, avatar, mid, tags, detail FROM movie ORDER BY stars LIMIT 20 OFFSET ?"

	Stmt, err := DB.Prepare(sqlStr)

	defer Stmt.Close()

	if err != nil {
		return
	}
	rows, err := Stmt.Query(start)

	for rows.Next() {
		err = rows.Scan(&movieList.Name, &Score, &movieList.Avatar, &movieList.Mid, &movieList.Tags, &Detail)
		if err != nil {
			return
		}
		err = json.Unmarshal([]byte(Detail), &movieList.Detail)
		if err != nil {
			return
		}
		err = json.Unmarshal([]byte(Score), &movieList.Score)
		if err != nil {
			return
		}
		movieLists = append(movieLists, movieList)
	}
	return
}

//分类找电影
func SelectSubjectsByTag(tag, sort string, start int) (movieLists []model.MovieList, err error) {

	var movieList model.MovieList
	var Detail, Score string
	sqlStr := "SELECT name, score, avatar, mid, tags, detail FROM movie WHERE tags LIKE '%{tag}%'"
	sqlStr = strings.Replace(sqlStr, "{tag}", tag, -1)
	sqlStr = sqlStr + " ORDER BY " + sort + " LIMIT 20 OFFSET ?"

	Stmt, err := DB.Prepare(sqlStr)
	defer Stmt.Close()

	if err != nil {
		return nil, err
	}

	rows, err := Stmt.Query(start)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&movieList.Name, &Score, &movieList.Avatar, &movieList.Mid, &movieList.Tags, &Detail)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(Detail), &movieList.Detail)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(Score), &movieList.Score)
		if err != nil {
			return nil, err
		}
		movieLists = append(movieLists, movieList)
	}
	return movieLists, nil
}

//搜索电影
func SearchMovies(key string) (movieLists []model.MovieList, err error) {
	var movieList model.MovieList
	var Detail, Score string
	sqlStr := "SELECT name, score, avatar, mid, tags, detail FROM movie WHERE name LIKE '%{}%'"
	rows, err := DB.Query(strings.Replace(sqlStr, "{}", key, -1))
	if err != nil {
		return
	}

	for rows.Next() {
		err = rows.Scan(&movieList.Name, &Score, &movieList.Avatar, &movieList.Mid, &movieList.Tags, &Detail)
		if err != nil {
			return
		}

		err = json.Unmarshal([]byte(Detail), &movieList.Detail)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(Score), &movieList.Score)
		if err != nil {
			return nil, err
		}
		movieLists = append(movieLists, movieList)
	}
	return
}

//获取单个电影的信息
func GetMovieById(mid int64) (movie model.Movie, err error) {
	sqlStr := "SELECT mid, tags, date, stars, name, avatar, detail, score, plot, celebrities FROM movie WHERE mid = ?"

	Stmt, err := DB.Prepare(sqlStr)

	defer Stmt.Close()

	if err != nil {
		return
	}
	row := Stmt.QueryRow(mid)
	var detail, score, celebrities string
	err = row.Scan(
		&movie.Mid,
		&movie.Tags,
		&movie.Date,
		&movie.Stars,
		&movie.Name,
		&movie.Avatar,
		&detail,
		&score,
		&movie.Plot,
		&celebrities,
	)
	err = json.Unmarshal([]byte(detail), &movie.Detail)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(score), &movie.Score)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(celebrities), &movie.Celebrities)
	return
}

//跟新得分
func UpdateSubjectScore(mid int64, score model.MovieScore) (err error) {
	sqlStr := "UPDATE movie SET score = ? WHERE mid = ?"

	scoreB, err := json.Marshal(score)
	if err != nil {
		return
	}

	Stmt, err := DB.Prepare(sqlStr)

	_, err = Stmt.Exec(string(scoreB), mid)

	return err
}
