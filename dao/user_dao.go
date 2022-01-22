package dao

import (
	"douban/model"
	"fmt"
)

//修改电话号码
func ChangePhone(phone string, id int64) error {
	sqlStr := "UPDATE user SET phone = ? WHERE id = ?"
	stmt, err := DB.Prepare(sqlStr)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(phone, id)
	if err != nil {
		return err
	}

	return nil

}

//判断电话是否注测
//修改邮箱
func ChangeEmail(email string, id int64) error {
	sqlStr := "UPDATE user SET email = ? WHERE id = ?"
	stmt, err := DB.Prepare(sqlStr)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(email, id)
	if err != nil {
		return err
	}

	return nil
}

//根据ID查询
func QueryUserByID(id int64) (model.User, error) {
	user := model.User{}

	sqlStr := "SELECT id,username,password,email,phone,salt,avatar,domain_name,habitat,hometown,birthday,statement,followers,followings FROM user WHERE id= ? "
	Stmt, err := DB.Prepare(sqlStr)
	defer Stmt.Close()

	if err != nil {
		return user, err
	}

	row := Stmt.QueryRow(id)
	if row.Err() != nil {
		return user, row.Err()
	}

	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Salt, &user.Avatar, &user.DomainName, &user.Habitat, &user.Username, &user.Birthday, &user.Statement, &user.Followers, &user.Followings)
	if err != nil {
		return user, err
	}

	return user, nil
}

//根据用户名查询
func QueryByUserName(username string) (model.User, error) {
	user := model.User{}

	sqlStr := "SELECT id,username,password,email,phone,salt,avatar,domain_name,habitat,hometown,birthday,statement,followers,followings FROM user WHERE username= ? "
	Stmt, err := DB.Prepare(sqlStr)
	defer Stmt.Close()

	if err != nil {
		return user, err
	}

	row := Stmt.QueryRow(username)
	if row.Err() != nil {
		return user, row.Err()
	}

	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Salt, &user.Avatar, &user.DomainName, &user.Habitat, &user.Username, &user.Birthday, &user.Statement, &user.Followers, &user.Followings)
	if err != nil {
		return user, err
	}

	return user, nil
}

//根据邮箱查询
func QueryByEmail(email string) (model.User, error) {
	user := model.User{}

	sqlStr := "SELECT id,username,password,email,phone,salt,avatar,domain_name,habitat,hometown,birthday,statement,followers,followings FROM user WHERE emial= ? "
	Stmt, err := DB.Prepare(sqlStr)
	defer Stmt.Close()

	if err != nil {
		return user, err
	}

	row := Stmt.QueryRow(email)
	if row.Err() != nil {
		return user, row.Err()
	}

	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Salt, &user.Avatar, &user.DomainName, &user.Habitat, &user.Username, &user.Birthday, &user.Statement, &user.Followers, &user.Followings)
	if err != nil {
		return user, err
	}

	return user, nil
}

//根据电话查询
func QueryByPhone(phone string) (model.User, error) {
	user := model.User{}

	sqlStr := "SELECT id,username,password,email,phone,salt,avatar,domain_name,habitat,hometown,birthday,statement,followers,followings FROM user WHERE phone= ? "
	Stmt, err := DB.Prepare(sqlStr)
	defer Stmt.Close()

	if err != nil {
		return user, err
	}

	row := Stmt.QueryRow(phone)
	if row.Err() != nil {
		return user, row.Err()
	}

	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Salt, &user.Avatar, &user.DomainName, &user.Habitat, &user.Username, &user.Birthday, &user.Statement, &user.Followers, &user.Followings)
	if err != nil {
		return user, err
	}

	return user, nil
}

//插入User信息
func InsertUser(user model.User) error {
	sqlStr := "INSERT INTO user(username, password, regdate, phone, salt) VALUES (?,?,?,?,?)"
	stmt, err := DB.Prepare(sqlStr)

	if err != nil {
		fmt.Println("零零零零")
		return err
	}
	_, err = stmt.Exec(user.Username, user.Password, user.RegDate, user.Phone, user.Salt)

	stmt.Close()

	return err
}
