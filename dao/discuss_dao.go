package dao

import (
	"douban/model"
	"fmt"
)

//插入讨论
func InsertDiscussion(discussion model.Discussion) error {
	sqlStr := "INSERT INTO discussion(uid,mid,username,title,value,time) VALUES (?,?,?,?,?)"
	stmt, err := DB.Prepare(sqlStr)

	if err != nil {
		fmt.Println("零零零零")
		return err
	}
	_, err = stmt.Exec(discussion.Uid, discussion.Mid, discussion.UserName, discussion.Title, discussion.Value, discussion.Date)
	stmt.Close()

	return err

}

//获取讨论列表
func GetDiscussionList(sort string, mid int64) (DiscussionLists []model.DiscussionList, err error) {
	var DiscussionList model.DiscussionList
	sqlStr := "SELECT id,mid,username,title, time ,reply FROM discussion WHERE mid=? ORDER BY " + sort

	Stmt, err := DB.Prepare(sqlStr)

	defer Stmt.Close()

	if err != nil {
		return nil, err
	}

	rows, err := Stmt.Query(mid)

	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	for rows.Next() {

		err = rows.Scan(&DiscussionList.Id, &DiscussionList.Mid, &DiscussionList.UserName, &DiscussionList.Title, &DiscussionList.Date, &DiscussionList.ReplyNum)
		if err != nil {
			return nil, err
		}

		DiscussionLists = append(DiscussionLists, DiscussionList)
	}

	if err != nil {
		return nil, err
	}

	return DiscussionLists, nil

}
