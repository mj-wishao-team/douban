package dao

import "douban/model"

func GetReply(id int64, kind string, start int) (Replys []model.Reply, err error) {
	var reply model.Reply
	stmt, err := DB.Prepare(`SELECT r.id ,r.uid,r.ptable ,r.value ,r.pid,r.people,r.like,r.time,u.avatar,u.username FROM reply r JOIN user u ON r.uid=u.id AND s.pid=? AND r.ptable=?  ORDER BY like  LIMIT 20 OFFSET ?`)
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
		err = rows.Scan(&reply.Id, &reply.Uid, &reply.Petable, &reply.Content, &reply.Pid, &reply.RepCnt, &reply.Like, &reply.Date, &reply.Avatar, &reply.Username)
		if err != nil {
			return nil, err
		}
		Replys = append(Replys, reply)
	}
	return
}
