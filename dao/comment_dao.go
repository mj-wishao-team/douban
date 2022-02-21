package dao

import (
	"douban/model"
)

//增加影评的回复人数
func UpdateReviewCNT(id int64) error {
	sqlStr := "UPDATE large_comment SET people=people+1 WHERE id = ?"
	stmt, err := DB.Prepare(sqlStr)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

//增加讨论的回复人数

func UpdateDiscussionCNT(id int64) error {
	sqlStr := "UPDATE reply SET people=people+1 WHERE id = ?"
	stmt, err := DB.Prepare(sqlStr)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

//增加回复的回复人数
func UpdateReplyCNT(id int64) error {
	sqlStr := "UPDATE discussion SET people=people+1 WHERE id = ?"
	stmt, err := DB.Prepare(sqlStr)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

//点赞短评
func UpdateCommentLike(id int64, like int) error {
	sqlStr := "UPDATE short_comment SET help=help+? WHERE id = ?"
	stmt, err := DB.Prepare(sqlStr)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(like, id)
	if err != nil {
		return err
	}

	return nil

}

//点赞影评
func UpdateReviewLike(id int64, like int) error {
	if like == -1 {
		sqlStr := "UPDATE large_comment SET unlike=unlike+1 WHERE id = ?"
		stmt, err := DB.Prepare(sqlStr)
		defer stmt.Close()

		if err != nil {
			return err
		}

		_, err = stmt.Exec(id)
		if err != nil {
			return err
		}

		return nil
	} else {
		sqlStr := "UPDATE large_comment SET likes=likes+1 WHERE id = ?"
		stmt, err := DB.Prepare(sqlStr)
		defer stmt.Close()

		if err != nil {
			return err
		}

		_, err = stmt.Exec(id)
		if err != nil {
			return err
		}

		return nil
	}
}

//获取影评
func GetMovieReviews(mid int64) ([]model.LargeComment, error) {
	var commentSlice []model.LargeComment

	stmt, err := DB.Prepare(`SELECT l.id, l.mid, u.avatar,u.username,l.uid,l.title, l.comment, l.time,l.likes,l.unlike,l.star,l.people FROM large_comment l JOIN user u ON l.uid=u.id AND l.mid=? ORDER BY l.likes LIMIT 10 `)
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
		err = rows.Scan(&commentModel.Id, &commentModel.Mid, &commentModel.Avatar, &commentModel.Username, &commentModel.Uid, &commentModel.Title, &commentModel.Comment, &commentModel.Time, &commentModel.Likes, &commentModel.Unlikes, &commentModel.Star, &commentModel.People)
		if err != nil {
			return nil, err
		}
		commentSlice = append(commentSlice, commentModel)
	}

	return commentSlice, nil
}

//获取自己短评
func GetShortCommentByUidAndMid(uid, mid int64) ([]model.ShortComment, error) {
	var commentSlice []model.ShortComment

	stmt, err := DB.Prepare(`SELECT id,mid,uid, comment,post_time,help,star,static FROM short_comment  WHERE uid=? AND mid=?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(uid, mid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var commentModel model.ShortComment
		err = rows.Scan(&commentModel.Id, &commentModel.Mid, &commentModel.Uid, &commentModel.Comment, &commentModel.Time, &commentModel.Help, &commentModel.Star, &commentModel.Static)
		if err != nil {
			return nil, err
		}
		commentSlice = append(commentSlice, commentModel)
	}

	return commentSlice, nil
}

//插入短评
func InsertShortComment(shortComment model.ShortComment) error {
	sqlStr := "INSERT INTO short_comment(mid,uid, comment, post_time,star,static) values(?,?,?,?,?,?);"
	Stmt, err := DB.Prepare(sqlStr)
	_, err = Stmt.Exec(shortComment.Mid, shortComment.Uid, shortComment.Comment, shortComment.Time, shortComment.Star, shortComment.Static)
	return err
	return nil
}

//插入影评论
func InsertLargeComment(largeComment model.LargeComment) error {
	sqlStr := "INSERT INTO large_comment(mid,uid,title,comment,time,star) values(?,?,?,?,?,?);"
	Stmt, err := DB.Prepare(sqlStr)
	_, err = Stmt.Exec(largeComment.Mid, largeComment.Uid, largeComment.Title, largeComment.Comment, largeComment.Time, largeComment.Star)
	if err != nil {
		return err
	}
	return nil
}

//获取短评
func QueryShortCommentByMid(mid int64) ([]model.ShortComment, error) {
	var commentSlice []model.ShortComment

	stmt, err := DB.Prepare(`SELECT s.id, s.mid, s.uid,u.avatar,u.username, s.comment, s.post_time,s.help,s.star,s.static FROM short_comment s JOIN user u ON s.uid=u.id AND s.mid=?`)
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
		err = rows.Scan(&commentModel.Id, &commentModel.Mid, &commentModel.Uid, &commentModel.Avatar, &commentModel.Username, &commentModel.Comment, &commentModel.Time, &commentModel.Help, &commentModel.Star, &commentModel.Static)
		if err != nil {
			return nil, err
		}
		commentSlice = append(commentSlice, commentModel)
	}

	return commentSlice, nil
}

func GetMovieComment(mid int64) ([]model.ShortComment, error) {
	var commentSlice []model.ShortComment
	stmt, err := DB.Prepare(`SELECT s.id, s.mid, s.uid,u.avatar,u.username, s.comment, s.post_time,s.help,s.star ,s.static FROM short_comment s JOIN user u ON s.uid=u.id AND s.mid=? ORDER BY s.help LIMIT 10 `)
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
		err = rows.Scan(&commentModel.Id, &commentModel.Mid, &commentModel.Uid, &commentModel.Avatar, &commentModel.Username, &commentModel.Comment, &commentModel.Time, &commentModel.Help, &commentModel.Star, &commentModel.Static)
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

	stmt, err := DB.Prepare(`SELECT l.id, l.mid, u.avatar,u.username,l.uid,l.title, l.comment, l.time,l.likes,l.unlike,l.star,l.people FROM large_comment l JOIN user u ON l.uid=u.id AND l.mid=?`)
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
		err = rows.Scan(&commentModel.Id, &commentModel.Mid, &commentModel.Avatar, &commentModel.Username, &commentModel.Uid, &commentModel.Title, &commentModel.Comment, &commentModel.Time, &commentModel.Likes, &commentModel.Unlikes, &commentModel.Star, &commentModel.People)
		if err != nil {
			return nil, err
		}
		commentSlice = append(commentSlice, commentModel)
	}

	return commentSlice, nil
}

//获取单个影评
func GetReview(id int64) (commentModels []model.LargeComment, err error) {

	stmt, err := DB.Prepare(`SELECT l.id, l.mid, u.avatar,u.username,l.uid,l.title, l.comment, l.time,l.likes,l.unlike,l.star,l.people FROM large_comment l JOIN user u ON l.uid=u.id AND l.id=?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var commentModel model.LargeComment
		err = rows.Scan(&commentModel.Id, &commentModel.Mid, &commentModel.Avatar, &commentModel.Username, &commentModel.Uid, &commentModel.Title, &commentModel.Comment, &commentModel.Time, &commentModel.Likes, &commentModel.Unlikes, &commentModel.Star, &commentModel.People)
		if err != nil {
			return nil, err
		}
		commentModels = append(commentModels, commentModel)
	}
	return commentModels, nil

}

func QueryLargeCommentByUid(uid int64) ([]model.LargeComment, error) {
	var commentSlice []model.LargeComment

	stmt, err := DB.Prepare(`SELECT id, mid, uid,title, comment, time,likes,unlike,star ,people FROM large_comment WHERE uid= ?`)
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
		err = rows.Scan(&commentModel.Id, &commentModel.Mid, &commentModel.Uid, &commentModel.Title, &commentModel.Comment, &commentModel.Time, &commentModel.Likes, &commentModel.Unlikes, &commentModel.Star, &commentModel.People)
		if err != nil {
			return nil, err
		}
		commentSlice = append(commentSlice, commentModel)
	}

	return commentSlice, nil
}
