package dao

import (
	"douban/model"
)

//插入短评
func InsertShortComment(shortComment model.ShortComment) error {
	sqlStr := "INSERT INTO short_comment(mid,uid, comment, post_time,star) values(?,?,?,?,?);"
	Stmt, err := DB.Prepare(sqlStr)
	_, err = Stmt.Exec(shortComment.Mid, shortComment.Uid, shortComment.Comment, shortComment.Time, shortComment.Star)
	return err
	return nil
}

//插入影评论
func InsertLargeComment(largeComment model.LargeComment) error {
	sqlStr := "INSERT INTO short_comment(mid,uid,title,comment, post_time,star) values(?,?,?,?,?,?);"
	Stmt, err := DB.Prepare(sqlStr)
	_, err = Stmt.Exec(largeComment.Mid, largeComment.Uid, largeComment.Title, largeComment.Comment, largeComment.Time, largeComment.Star)
	return err
	return nil
}

//获取短评
func QueryShortCommentByMid(mid int64) ([]model.ShortComment, error) {
	var commentSlice []model.ShortComment

	stmt, err := DB.Prepare(`SELECT id, mid, uid, comment, post_time,help,star,report FROM short_comment WHERE mid= ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(mid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var commentModel model.ShortComment
		err = rows.Scan(&commentModel.Id, &commentModel.Mid, &commentModel.Uid, &commentModel.Comment, &commentModel.Time, &commentModel.Help, &commentModel.Star, &commentModel.Report)
		if err != nil {
			return nil, err
		}
		commentSlice = append(commentSlice, commentModel)
	}

	return commentSlice, nil
}

//获取影评论
func QueryLargeCommentByMid(mid int64) ([]model.LargeComment, error) {
	var commentSlice []model.LargeComment

	stmt, err := DB.Prepare(`SELECT id, mid, uid,title, comment, time,likes,unlike,star,report FROM large_comment WHERE mid= ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(mid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var commentModel model.LargeComment
		err = rows.Scan(&commentModel.Id, &commentModel.Mid, &commentModel.Uid, &commentModel.Title, &commentModel.Comment, &commentModel.Time, &commentModel.Likes, &commentModel.Unlikes, &commentModel.Star, &commentModel.Report)
		if err != nil {
			return nil, err
		}
		commentSlice = append(commentSlice, commentModel)
	}

	return commentSlice, nil
}
