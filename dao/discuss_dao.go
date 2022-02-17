package dao

import (
	"douban/model"
	"fmt"
)

//插入讨论
func InsertDiscussion(discussion model.Discussion) error {
	sqlStr := "INSERT INTO discussion(uid,mid,title,value,time) VALUES (?,?,?,?,?)"
	stmt, err := DB.Prepare(sqlStr)

	if err != nil {
		fmt.Println("零零零零")
		return err
	}
	_, err = stmt.Exec(discussion.Uid, discussion.Mid, discussion.Title, discussion.Value, discussion.Date)
	stmt.Close()

	return err

}
