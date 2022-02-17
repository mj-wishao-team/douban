package dao

import (
	"douban/model"
)

//获取讨论
func GetDiscussion(id int64) (Discussion []model.Discussion, err error) {
	var discussion model.Discussion

	sqlStr := "SELECT d.id ,d.mid,d.uid,u.username,u.avatar,d.tiltle,d.value,d.time,d.reply FROM discussion d JOIN user u ON  d.uid=u.id AND d.id= ? "

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
		&discussion.Id,
		&discussion.Mid,
		&discussion.Uid,
		&discussion.UserName,
		&discussion.Avatar,
		&discussion.Title,
		&discussion.Value,
		&discussion.Date,
		&discussion.ReplyNum)

	if err != nil {
		return nil, err
	}
	Discussion = append(Discussion, discussion)
	return Discussion, nil
}

//插入讨论
func InsertDiscussion(discussion model.Discussion) error {
	sqlStr := "INSERT INTO discussion(uid,mid,username,title,value,time) VALUES (?,?,?,?,?)"
	stmt, err := DB.Prepare(sqlStr)

	if err != nil {
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
