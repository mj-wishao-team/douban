package dao

import (
	"douban/model"
	"errors"
)

func GetReply(id int64, kind string, start int) (Replys []model.Reply, err error) {
	var reply model.Reply
	stmt, err := DB.Prepare(`SELECT r.id ,r.uid,r.ptable ,r.value,r.pid,r.people,r.likes,r.time,u.avatar,u.username FROM reply r JOIN user u ON r.uid=u.id AND r.pid=? AND r.ptable=?  ORDER BY likes LIMIT 20 OFFSET ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id, kind, start)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&reply.Id, &reply.Uid, &reply.Ptable, &reply.Content, &reply.Pid, &reply.RepCnt, &reply.Like, &reply.Date, &reply.Avatar, &reply.Username)
		if err != nil {
			return nil, err
		}
		Replys = append(Replys, reply)
	}
	return
}

func ReplyPost(reply model.Reply) error {
	sqlStr := "INSERT INTO reply(uid, pid, ptable, time , value) VALUES(?, ?, ?, ?, ?)"
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(reply.Uid, reply.Pid, reply.Ptable, reply.Date, reply.Content)
	return err
}

//采用事务处理
func PostReply(reply model.Reply) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	var sqlStr2 string

	sqlStr1 := "INSERT INTO reply(uid, pid, ptable, time , value) VALUES(?, ?, ?, ?, ?)"

	switch reply.Ptable {
	case "review":
		sqlStr2 = "UPDATE large_comment SET people=people+1 WHERE id=?"
	case "comment":
		sqlStr2 = "UPDATE short_comment SET people=people+1 WHERE id=?"
	case "discussion":
		sqlStr2 = "UPDATE discussion SET people=people+1 WHERE id=?"
	default:
		return errors.New("ptable类型错误")
	}

	stmt1, err := tx.Prepare(sqlStr1)
	if err != nil {
		return err
	}
	_, err = stmt1.Exec(reply.Uid, reply.Pid, reply.Ptable, reply.Date, reply.Content)

	//回滚事务
	if err != nil {
		tx.Rollback()
		return err
	}

	stmt2, err := tx.Prepare(sqlStr2)
	if err != nil {
		return err
	}
	_, err = stmt2.Exec(reply.Pid)
	//回滚事务
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}
