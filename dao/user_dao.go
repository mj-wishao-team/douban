package dao

import (
	"douban/model"
	"fmt"
	"time"
)

//修改头像
func ChangeAvatar(url string, id int64) error {
	sqlStr := "UPDATE user SET avatar = ? WHERE id = ?"
	stmt, err := DB.Prepare(sqlStr)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(url, id)
	if err != nil {
		return err
	}

	return nil
}

//修改常驻地
func ChangeHabitat(habitat string, id int64) error {
	sqlStr := "UPDATE user SET habitat = ? WHERE id = ?"
	stmt, err := DB.Prepare(sqlStr)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(habitat, id)
	if err != nil {
		return err
	}

	return nil
}

//修改家乡公开
func ChangeHometownPublic(hometownPublic string, id int64) error {
	sqlStr := "UPDATE user SET hometown_public = ? WHERE id = ?"
	stmt, err := DB.Prepare(sqlStr)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(hometownPublic, id)
	if err != nil {
		return err
	}

	return nil
}

//修改生日公开
func ChangeBirthdayPublic(birthdayPublic string, id int64) error {
	sqlStr := "UPDATE user SET birthday_public = ? WHERE id = ?"
	stmt, err := DB.Prepare(sqlStr)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(birthdayPublic, id)
	if err != nil {
		return err
	}

	return nil

}

//修改生日
func ChangeBirthday(birthday time.Time, id int64) error {
	sqlStr := "UPDATE user SET birthday = ? WHERE id = ?"
	stmt, err := DB.Prepare(sqlStr)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(birthday, id)
	if err != nil {
		return err
	}

	return nil
}

//修改家乡
func ChangeHometown(hometown string, id int64) error {
	sqlStr := "UPDATE user SET hometown = ? WHERE id = ?"
	stmt, err := DB.Prepare(sqlStr)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(hometown, id)
	if err != nil {
		return err
	}

	return nil

}

//修改名字
func ChangUserName(name string, id int64) error {
	sqlStr := "UPDATE user SET username = ? WHERE id = ?"
	stmt, err := DB.Prepare(sqlStr)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, id)
	if err != nil {
		return err
	}

	return nil
}

//删除账号
func DeleteAccount(id int64) error {
	sqlStr := "DELETE FROM user WHERE( id= ?)"
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err

}

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

//根据ID查询UserInfo信息
func QueryUserInfoByID(id int64) (UserInfo model.UserInfo, err error) {
	user := model.UserInfo{}

	sqlStr := "SELECT id,username,email,phone,avatar,domain_name,habitat,hometown,birthday,statement,followers,followings,regdate FROM user WHERE id= ? "
	Stmt, err := DB.Prepare(sqlStr)
	defer Stmt.Close()

	if err != nil {
		return user, err
	}

	row := Stmt.QueryRow(id)
	if row.Err() != nil {
		return user, row.Err()
	}

	err = row.Scan(&user.Id, &user.Username, &user.Email, &user.Phone, &user.Avatar, &user.DomainName, &user.Habitat, &user.Hometown, &user.Birthday, &user.Statement, &user.Followers, &user.Followings, &user.RegDate)
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
