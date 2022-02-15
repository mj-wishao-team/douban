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

	stmt, err := DB.Prepare(`SELECT s.id, s.mid, s.uid,u.avatar,u.username, s.comment, s.post_time,s.help,s.star,s.report FROM short_comment s JOIN user u ON s.uid=u.id AND s.mid=?`)
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
		err = rows.Scan(&commentModel.Id, &commentModel.Mid, &commentModel.Uid, &commentModel.Avatar, &commentModel.Username, &commentModel.Comment, &commentModel.Time, &commentModel.Help, &commentModel.Star, &commentModel.Report)
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

	stmt, err := DB.Prepare(`SELECT l.id, l.mid, u.avatar,u.username,l.uid,l.title, l.comment, l.time,l.likes,l.unlike,l.star,l.report FROM large_comment l JOIN user u ON l.uid=u.id AND l.mid=?`)
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
		err = rows.Scan(&commentModel.Id, &commentModel.Mid, &commentModel.Avatar, &commentModel.Username, &commentModel.Uid, &commentModel.Title, &commentModel.Comment, &commentModel.Time, &commentModel.Likes, &commentModel.Unlikes, &commentModel.Star, &commentModel.Report)
		if err != nil {
			return nil, err
		}
		commentSlice = append(commentSlice, commentModel)
	}

	return commentSlice, nil
}

func QueryLargeCommentByUid(uid int64) ([]model.LargeComment, error) {
	var commentSlice []model.LargeComment

	stmt, err := DB.Prepare(`SELECT id, mid, uid,title, comment, time,likes,unlike,star,report FROM large_comment WHERE uid= ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(uid)
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
