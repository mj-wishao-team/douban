package dao

import (
	"douban/model"
	"gorm.io/gorm"
)

//增加影评的回复人数  由于加入了事务所废弃
//func UpdateReviewCNT(id int64) error {
//	sqlStr := "UPDATE large_comment SET people=people+1 WHERE id = ?"
//	stmt, err := DB.Prepare(sqlStr)
//	defer stmt.Close()
//
//	if err != nil {
//		return err
//	}
//
//	_, err = stmt.Exec(id)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//增加讨论的回复人数
//
//func UpdateDiscussionCNT(id int64) error {
//	sqlStr := "UPDATE reply SET people=people+1 WHERE id = ?"
//	stmt, err := DB.Prepare(sqlStr)
//	defer stmt.Close()
//
//	if err != nil {
//		return err
//	}
//
//	_, err = stmt.Exec(id)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
////增加回复的回复人数
//func UpdateReplyCNT(id int64) error {
//	sqlStr := "UPDATE discussion SET people=people+1 WHERE id = ?"
//	stmt, err := DB.Prepare(sqlStr)
//	defer stmt.Close()
//
//	if err != nil {
//		return err
//	}
//
//	_, err = stmt.Exec(id)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//点赞短评
//func UpdateCommentLike(id int64, like int) error {
//	sqlStr := "UPDATE short_comment SET help=help+? WHERE id = ?"
//	stmt, err := DB.Prepare(sqlStr)
//	defer stmt.Close()
//
//	if err != nil {
//		return err
//	}
//
//	_, err = stmt.Exec(like, id)
//	if err != nil {
//		return err
//	}
//
//	return nil
//
//}

//点赞短评
func UpdateCommentLike(id int64, like int) error {
	rs := Db.Model(&model.Comment{}).Where(&model.Comment{Id: id}).Update("help", "help+1")
	return rs.Error
}

//点赞影评
//func UpdateReviewLike(id int64, like int) error {
//	if like == -1 {
//		sqlStr := "UPDATE large_comment SET unlike=unlike+1 WHERE id = ?"
//		stmt, err := DB.Prepare(sqlStr)
//		defer stmt.Close()
//
//		if err != nil {
//			return err
//		}
//
//		_, err = stmt.Exec(id)
//		if err != nil {
//			return err
//		}
//
//		return nil
//	} else {
//		sqlStr := "UPDATE large_comment SET likes=likes+1 WHERE id = ?"
//		stmt, err := DB.Prepare(sqlStr)
//		defer stmt.Close()
//
//		if err != nil {
//			return err
//		}
//
//		_, err = stmt.Exec(id)
//		if err != nil {
//			return err
//		}
//
//		return nil
//	}
//}

//点赞影评
func UpdateReviewLike(id int64, like int) error {
	if like == 1 {
		rs := Db.Model(&model.Review{}).Where("id=?", id).Update("likes", gorm.Expr("likes + ?", like))
		return rs.Error
	}
	rs := Db.Model(&model.Review{}).Where("id=?", id).Update("unlike", gorm.Expr("unlike + ?", like))
	return rs.Error
}

//获取影评

//func GetMovieReviews(mid int64) ([]model.Review, error) {
//	var commentSlice []model.Review
//
//	stmt, err := DB.Prepare(`SELECT l.id, l.mid, u.avatar,u.username,l.uid,l.title, l.comment, l.time,l.likes,l.unlike,l.star,l.people FROM large_comment l JOIN user u ON l.uid=u.id AND l.mid=? ORDER BY l.likes LIMIT 10 `)
//	if err != nil {
//		return nil, err
//	}
//	defer stmt.Close()
//
//	rows, err := stmt.Query(mid)
//	if err != nil {
//		return nil, err
//	}
//
//	defer rows.Close()
//	for rows.Next() {
//		var commentModel model.Review
//		err = rows.Scan(&commentModel.Id, &commentModel.Mid, &commentModel.Avatar, &commentModel.Username, &commentModel.Uid, &commentModel.Title, &commentModel.Comment, &commentModel.Time, &commentModel.Likes, &commentModel.Unlikes, &commentModel.Star, &commentModel.People)
//		if err != nil {
//			return nil, err
//		}
//		commentSlice = append(commentSlice, commentModel)
//	}
//
//	return commentSlice, nil
//}

//获取影评
func GetMovieReviews(mid int64) ([]model.Review, error) {
	var reviews []model.Review
	rs := Db.Where("mid=?", mid).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Model(&model.User{}).Find(&model.UserInfo{})
	}).Order("likes DESC").Limit(10).Find(&reviews)
	return reviews, rs.Error

}

//获取自己短评
//func GetShortCommentByUidAndMid(uid, mid int64) ([]model.Comment, error) {
//	var commentSlice []model.Comment
//
//	stmt, err := DB.Prepare(`SELECT id,mid,uid, comment,post_time,help,star,static FROM short_comment  WHERE uid=? AND mid=?`)
//	if err != nil {
//		return nil, err
//	}
//	defer stmt.Close()
//
//	rows, err := stmt.Query(uid, mid)
//	if err != nil {
//		return nil, err
//	}
//
//	defer rows.Close()
//	for rows.Next() {
//		var commentModel model.Comment
//		err = rows.Scan(&commentModel.Id, &commentModel.Mid, &commentModel.Uid, &commentModel.Comment, &commentModel.Time, &commentModel.Help, &commentModel.Star, &commentModel.Static)
//		if err != nil {
//			return nil, err
//		}
//		commentSlice = append(commentSlice, commentModel)
//	}
//
//	return commentSlice, nil
//}

//获取自己短评
func GetShortCommentByUidAndMid(uid, mid int64) ([]model.Comment, error) {
	var comment []model.Comment
	rs := Db.Where("uid=? AND mid=?", uid, mid).Find(&comment)
	return comment, rs.Error

}

//插入短评
//func InsertShortComment(shortComment model.Comment) error {
//	sqlStr := "INSERT INTO short_comment(mid,uid, comment, post_time,star,static) values(?,?,?,?,?,?);"
//	Stmt, err := DB.Prepare(sqlStr)
//	_, err = Stmt.Exec(shortComment.Mid, shortComment.Uid, shortComment.Comment, shortComment.Time, shortComment.Star, shortComment.Static)
//	return err
//	return nil
//}
func InsertShortComment(comment model.Comment) error {
	rs := Db.Create(&comment)
	return rs.Error
}

//插入影评论
//func InsertLargeComment(largeComment model.Review) error {
//	sqlStr := "INSERT INTO large_comment(mid,uid,title,comment,time,star) values(?,?,?,?,?,?);"
//	Stmt, err := DB.Prepare(sqlStr)
//	_, err = Stmt.Exec(largeComment.Mid, largeComment.Uid, largeComment.Title, largeComment.Comment, largeComment.Time, largeComment.Star)
//	if err != nil {
//		return err
//	}
//	return nil
//}
func InsertLargeComment(review model.Review) error {
	rs := Db.Create(&review)
	return rs.Error
}

//获取短评
//func QueryShortCommentByMid(mid int64) ([]model.Comment, error) {
//	var commentSlice []model.Comment
//
//	stmt, err := DB.Prepare(`SELECT s.id, s.mid, s.uid,u.avatar,u.username, s.comment, s.post_time,s.help,s.star,s.static FROM short_comment s JOIN user u ON s.uid=u.id AND s.mid=?`)
//	if err != nil {
//		return nil, err
//	}
//	defer stmt.Close()
//
//	rows, err := stmt.Query(mid)
//	if err != nil {
//		return nil, err
//	}
//
//	defer rows.Close()
//	for rows.Next() {
//		var commentModel model.Comment
//		err = rows.Scan(&commentModel.Id, &commentModel.Mid, &commentModel.Uid, &commentModel.Avatar, &commentModel.Username, &commentModel.Comment, &commentModel.Time, &commentModel.Help, &commentModel.Star, &commentModel.Static)
//		if err != nil {
//			return nil, err
//		}
//		commentSlice = append(commentSlice, commentModel)
//	}
//
//	return commentSlice, nil
//}

func QueryShortCommentByMid(mid int64) ([]model.Comment, error) {
	var comments []model.Comment
	rs := Db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Model(&model.User{}).Find(&model.UserInfo{})
	}).Find(&comments)
	return comments, rs.Error
}

//
//func GetMovieComment(mid int64) ([]model.Comment, error) {
//
//	var commentSlice []model.Comment
//	stmt, err := DB.Prepare(`SELECT s.id, s.mid, s.uid,u.avatar,u.username, s.comment, s.post_time,s.help,s.star ,s.static FROM short_comment s JOIN user u ON s.uid=u.id AND s.mid=? ORDER BY s.help LIMIT 10 `)
//	if err != nil {
//		return nil, err
//	}
//	defer stmt.Close()
//
//	rows, err := stmt.Query(mid)
//	if err != nil {
//		return nil, err
//	}
//
//	defer rows.Close()
//	for rows.Next() {
//		var commentModel model.Comment
//		err = rows.Scan(&commentModel.Id, &commentModel.Mid, &commentModel.Uid, &commentModel.Avatar, &commentModel.Username, &commentModel.Comment, &commentModel.Time, &commentModel.Help, &commentModel.Star, &commentModel.Static)
//		if err != nil {
//			return nil, err
//		}
//		commentSlice = append(commentSlice, commentModel)
//	}
//
//	return commentSlice, nil
//}

func GetMovieComment(mid int64) ([]model.Comment, error) {
	var comments []model.Comment
	rs := Db.Where("mid=?", mid).Preload("users", func(db *gorm.DB) *gorm.DB {
		return db.Select("avatar", "username", "uid").Where("mid=?", mid)
	}).Order("help").Limit(10).Find(&comments)
	return comments, rs.Error
}

//获取影评论
//func QueryLargeCommentByMid(mid int64) ([]model.Review, error) {
//	var commentSlice []model.Review
//
//	stmt, err := DB.Prepare(`SELECT l.id, l.mid, u.avatar,u.username,l.uid,l.title, l.comment, l.time,l.likes,l.unlike,l.star,l.people FROM large_comment l JOIN user u ON l.uid=u.id AND l.mid=?`)
//	if err != nil {
//		return nil, err
//	}
//	defer stmt.Close()
//
//	rows, err := stmt.Query(mid)
//	if err != nil {
//		return nil, err
//	}
//
//	defer rows.Close()
//	for rows.Next() {
//		var commentModel model.Review
//		err = rows.Scan(&commentModel.Id, &commentModel.Mid, &commentModel.Avatar, &commentModel.Username, &commentModel.Uid, &commentModel.Title, &commentModel.Comment, &commentModel.Time, &commentModel.Likes, &commentModel.Unlikes, &commentModel.Star, &commentModel.People)
//		if err != nil {
//			return nil, err
//		}
//		commentSlice = append(commentSlice, commentModel)
//	}
//
//	return commentSlice, nil
//}
func QueryLargeCommentByMid(mid int64) ([]model.Review, error) {
	var review []model.Review
	rs := Db.Where(&model.Review{Mid: mid}).Find(&review)
	return review, rs.Error
}

//获取单个影评
//func GetReview(id int64) (commentModels []model.Review, err error) {
//
//	stmt, err := DB.Prepare(`SELECT l.id, l.mid, u.avatar,u.username,l.uid,l.title, l.comment, l.time,l.likes,l.unlike,l.star,l.people FROM large_comment l JOIN user u ON l.uid=u.id AND l.id=?`)
//	if err != nil {
//		return nil, err
//	}
//	defer stmt.Close()
//
//	rows, err := stmt.Query(id)
//	if err != nil {
//		return nil, err
//	}
//
//	for rows.Next() {
//		var commentModel model.Review
//		err = rows.Scan(&commentModel.Id, &commentModel.Mid, &commentModel.Avatar, &commentModel.Username, &commentModel.Uid, &commentModel.Title, &commentModel.Comment, &commentModel.Time, &commentModel.Likes, &commentModel.Unlikes, &commentModel.Star, &commentModel.People)
//		if err != nil {
//			return nil, err
//		}
//		commentModels = append(commentModels, commentModel)
//	}
//	return commentModels, nil
//
//}

func GetReview(id int64) (review []model.Review, err error) {
	rs := Db.Where(&model.Review{Id: id}).Find(&review)
	return review, rs.Error
}

func QueryLargeCommentByUid(uid int64) ([]model.Review, error) {
	var review []model.Review
	rs := Db.Where(&model.Review{Id: uid}).Find(&review)
	return review, rs.Error
}

//func QueryLargeCommentByUid(uid int64) ([]model.Review, error) {
//	var commentSlice []model.Review
//
//	stmt, err := DB.Prepare(`SELECT id, mid, uid,title, comment, time,likes,unlike,star ,people FROM large_comment WHERE uid= ?`)
//	if err != nil {
//		return nil, err
//	}
//	defer stmt.Close()
//
//	rows, err := stmt.Query(uid)
//	if err != nil {
//		return nil, err
//	}
//
//	defer rows.Close()
//	for rows.Next() {
//		var commentModel model.Review
//		err = rows.Scan(&commentModel.Id, &commentModel.Mid, &commentModel.Uid, &commentModel.Title, &commentModel.Comment, &commentModel.Time, &commentModel.Likes, &commentModel.Unlikes, &commentModel.Star, &commentModel.People)
//		if err != nil {
//			return nil, err
//		}
//		commentSlice = append(commentSlice, commentModel)
//	}
//
//	return commentSlice, nil
//}
