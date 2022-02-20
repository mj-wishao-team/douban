package dao

import (
	"douban/model"
	"fmt"
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
func GetMyLook(uid int64) (MyLooks []model.MovieStatic, err error) {

	var myLook model.MovieStatic
	sqlStr := "SELECT s.uid,s.mid,s.movie_type,m.avatar,m.name FROM movie_static s JOIN movie m ON s.mid=m.mid AND s.uid=? "

	Stmt, err := DB.Prepare(sqlStr)
	defer Stmt.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	rows := Stmt.QueryRow(uid)
	err = rows.Scan(&myLook.Uid, &myLook.Mid, &myLook.Type, &myLook.MovieAvatar, &myLook.MovieName)
	if err != nil {
		return
	}
	MyLooks = append(MyLooks, myLook)
	return
}
