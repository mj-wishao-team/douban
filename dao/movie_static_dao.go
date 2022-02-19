package dao

import (
	"douban/model"
)

//插入我看内容
func InsertMovieStatic(MovieType string, uid, mid int64) error {
	sqlStr := "INSERT INTO movie_static(uid,mid,movie_type) VALUES (?,?,?)"
	stmt, err := DB.Prepare(sqlStr)

	if err != nil {
		return err
	}
	_, err = stmt.Exec(uid, mid, MovieType)

	stmt.Close()

	return err
}

//获取我看内容
func GetMyLook(uid int64, str string) (MyLooks []model.MovieStatic, err error) {

	var myLook model.MovieStatic
	sqlStr := "SELECT s.uid,s.mid,s.movie_type,m.avatar,m.name form movie_static s JOIN movie m ON s.mid=m.mid AND s.uid=? AND s.movie_type=?"

	Stmt, err := DB.Prepare(sqlStr)

	defer Stmt.Close()

	if err != nil {
		return
	}
	rows, err := Stmt.Query(uid, str)

	for rows.Next() {
		err = rows.Scan(&myLook.Uid, &myLook.Mid, &myLook.Type, &myLook.MovieAvatar, &myLook.MovieName)
		if err != nil {
			return
		}

		MyLooks = append(MyLooks, myLook)
	}
	return
}
