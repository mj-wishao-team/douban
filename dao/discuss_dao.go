package dao

import (
	"douban/model"
	"gorm.io/gorm"
)

//获取讨论
//func GetDiscussion(id int64) (model.Discussion, error) {
//	var discussion model.Discussion
//
//	sqlStr := "SELECT d.id ,d.mid,d.uid,u.username,u.avatar,d.title,d.value,d.time,d.people FROM discussion d JOIN user u ON  d.uid=u.id AND d.id= ? "
//
//	Stmt, err := DB.Prepare(sqlStr)
//	defer Stmt.Close()
//
//	if err != nil {
//		return discussion, err
//	}
//
//	row := Stmt.QueryRow(id)
//
//	if row.Err() != nil {
//		return discussion, row.Err()
//	}
//
//	err = row.Scan(
//		&discussion.Id,
//		&discussion.Mid,
//		&discussion.Uid,
//		&discussion.UserName,
//		&discussion.Avatar,
//		&discussion.Title,
//		&discussion.Value,
//		&discussion.Date,
//		&discussion.ReplyNum)
//
//	if err != nil {
//		return discussion, err
//	}
//	return discussion, nil
//}

//获取讨论
func GetDiscussion(id int64) (model.Discussion, error) {
	var dicussions model.Discussion
	rs := Db.Where("id=?", id).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Model(&model.User{}).Find(&model.UserInfo{})
	}).Find(&dicussions)

	return dicussions, rs.Error
}

//插入讨论
//func InsertDiscussion(discussion model.Discussion) error {
//	sqlStr := "INSERT INTO discussion(uid,mid,title,value,time) VALUES (?,?,?,?,?)"
//	stmt, err := DB.Prepare(sqlStr)
//
//	if err != nil {
//		return err
//	}
//	_, err = stmt.Exec(discussion.Uid, discussion.Mid, discussion.Title, discussion.Value, discussion.Date)
//	stmt.Close()
//
//	return err
//
//}

//插入讨论
func InsertDiscussion(discussion model.Discussion) error {
	rs := Db.Create(&discussion)
	return rs.Error
}

//获取讨论列表
//func GetDiscussionList(sort string, mid int64) (DiscussionLists []model.DiscussionList, err error) {
//	var DiscussionList model.DiscussionList
//	sqlStr := "SELECT d.id ,d.mid,d.uid,u.username,d.title,d.time,d.people FROM discussion d JOIN user u ON  d.uid=u.id AND d.mid= ?"
//	//sqlStr = sqlStr + "ORDER BY " + sort
//	Stmt, err := DB.Prepare(sqlStr)
//
//	defer Stmt.Close()
//
//	if err != nil {
//		return nil, err
//	}
//
//	rows, err := Stmt.Query(mid)
//
//	if err != nil {
//		return nil, err
//	}
//
//	if rows.Err() != nil {
//		return nil, rows.Err()
//	}
//
//	for rows.Next() {
//
//		err = rows.Scan(&DiscussionList.Id, &DiscussionList.Mid, &DiscussionList.Uid, &DiscussionList.UserName, &DiscussionList.Title, &DiscussionList.Date, &DiscussionList.ReplyNum)
//		if err != nil {
//			return nil, err
//		}
//
//		DiscussionLists = append(DiscussionLists, DiscussionList)
//	}
//
//	if err != nil {
//		return nil, err
//	}
//
//	return DiscussionLists, nil
//}

//获取讨论列表
func GetDiscussionList(sort string, mid int64) (DiscussionLists []model.DiscussionList, err error) {
	rs := Db.Where("mid=?", mid).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Model(&model.User{}).Find(&model.UserInfo{})
	}).Find(&DiscussionLists)
	return DiscussionLists, rs.Error
}

//删除讨论
//func DeleteDiscussion(id int64) error {
//	sqlStr := "DELETE FROM discussion WHERE id= ?"
//	stmt, err := DB.Prepare(sqlStr)
//	if err != nil {
//		return err
//	}
//	defer stmt.Close()
//	_, err = stmt.Exec(id)
//	return err
//
//}

//删除讨论
func DeleteDiscussion(id int64) error {
	var discussion model.Discussion
	rs := Db.Where("id=?", id).Delete(&discussion)
	return rs.Error
}

//跟新讨论
//func UpdateDiscussion(discussion model.Discussion) error {
//	sqlStr := "UPDATE discussion SET title = ?, value = ?, time = ? WHERE id = ? AND uid = ?"
//	stmt, err := DB.Prepare(sqlStr)
//	if err != nil {
//		return err
//	}
//	_, err = stmt.Exec(discussion.Title, discussion.Value, discussion.Date, discussion.Id, discussion.Uid)
//	return err
//}
//跟新讨论
func UpdateDiscussion(discussion model.Discussion) error {
	rs := Db.Model(&model.Discussion{}).Updates(discussion)
	return rs.Error
}

//讨论点赞
//func DiscussLike(id int64) error {
//	sqlStr := "UPDATE discussion SET likes=likes+1 WHERE id=?"
//	stmt, err := DB.Prepare(sqlStr)
//	if err != nil {
//		return err
//	}
//	_, err = stmt.Exec(id)
//	return err
//}

func DiscussLike(id int64) error {
	rs := Db.Model(&model.Discussion{}).Where("id=?", id).Update("likes", gorm.Expr("likes + ?", 1))
	return rs.Error
}
