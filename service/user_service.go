package service

import (
	"douban/dao"
	"douban/model"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/smtp"
	"strings"
	"time"
)

//修改头像
func ChangeAvatar(Url string, id int64) error {
	err := dao.ChangeAvatar(Url, id)
	return err
}

//修改常驻地
func ChangeHabitat(habitat string, id int64) error {
	err := dao.ChangeHabitat(habitat, id)
	return err
}

//修改是否公开生日
func ChangeBirthdayPublic(birthdayPublic string, id int64) error {
	err := dao.ChangeBirthdayPublic(birthdayPublic, id)
	return err
}

//修改是否公开家乡
func ChangeHometownPubic(hometownPublic string, id int64) error {
	err := dao.ChangeHometownPublic(hometownPublic, id)
	return err
}

//修改生日
func ChangeBirthday(birthday time.Time, id int64) error {
	err := dao.ChangeBirthday(birthday, id)
	return err
}

//修改家乡地址
func ChangeHometown(hometown string, id int64) error {
	err := dao.ChangeHometown(hometown, id)
	return err
}

//修改用户名
func ChangeUserName(name string, id int64) error {
	err := dao.ChangUserName(name, id)
	return err
}

//删除账户
func DeleteAccount(id int64) error {
	err := dao.DeleteAccount(id)
	return err
}

//判断ID是否正确，并获取User信息
//ture 正确 flase 错误
func JudgeAndQueryUserByUserID(Id int64) (model.User, bool, error) {
	User, err := dao.QueryUserByID(Id)
	if err != nil {
		//判断错误类型
		if err.Error() == "sql: no rows in result set" {
			return model.User{}, false, nil
		}
		return model.User{}, false, err
	}

	return User, true, nil

}

//根据Id查询用户
func GetUserById(id int64) (model.User, error) {
	User, err := dao.QueryUserByID(id)
	return User, err
}

//判断电话号码是否注册
//true 注册 flase 未注册
func JudgeAndQueryUserByPhone(phone string) (model.User, bool, error) {
	User, err := dao.QueryByPhone(phone)
	if err != nil {
		//判断错误类型
		if err.Error() == "sql: no rows in result set" {
			return model.User{}, false, nil
		}
		return model.User{}, false, err
	}

	return User, true, nil
}

//判断用户名是否注册
//true 注册 flase 未注册
func JudgeAndQueryUserByUserName(username string) (model.User, bool, error) {
	User, err := dao.QueryByUserName(username)
	if err != nil {
		//判断错误类型
		if err.Error() == "sql: no rows in result set" {
			return model.User{}, false, nil
		}
		return model.User{}, false, err
	}

	return User, true, nil
}

//判断邮箱是否注册
//true 注册 flase 未注册
func JudgeAndQueryUserByEmail(email string) (model.User, bool, error) {
	User, err := dao.QueryByEmail(email)
	if err != nil {
		//判断错误类型
		if err.Error() == "sql: no rows in result set" {
			return model.User{}, false, nil
		}
		return model.User{}, false, err
	}

	return User, true, nil
}

//判断密码
func JudgePasswordCorrect(loginAccount, password string) (model.User, bool, error) {
	// 判断登录类型
	flag := tool.VerifyEmailFormat(strings.ToLower(loginAccount))
	if flag {
		//邮箱登录
		user, err := dao.QueryByEmail(loginAccount)
		//判断错误类型
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				return model.User{}, false, nil
			}
			return model.User{}, false, err
		}
		//判断密码
		if tool.Match(user.Password, password, user.Salt) {
			return user, true, nil
		} else {
			return model.User{}, false, nil
		}
	} else {
		//手机号登录
		user, err := dao.QueryByPhone(loginAccount)
		//判断错误类型
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				return model.User{}, false, nil
			}
			return model.User{}, false, err
		}
		//判断密码
		if tool.Match(user.Password, password, user.Salt) {
			return user, true, nil
		} else {
			return model.User{}, false, nil
		}
	}
}

//发送短信
func SendSms(phone string) (string, error) {
	//生成六位数验证码
	code := tool.RandCode()
	err := tool.SendSms(phone, code)
	if err != nil {
		return "", err
	}
	return code, err
}

//发送邮箱

//以下使用网易邮箱SMTP
//各类邮箱配置https://blog.csdn.net/weixin_45604257/article/details/102963591?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522164264601116780265413399%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=164264601116780265413399&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~rank_v31_ecpm-1-102963591.first_rank_v2_pc_rank_v29&utm_term=golang%E5%AE%9E%E7%8E%B0%E5%8F%91%E9%80%81%E9%82%AE%E7%AE%B1%E9%AA%8C%E8%AF%81%E7%A0%81&spm=1018.2226.3001.4187
func SendCodeByEmail(email string) (string, error) {
	emailCfg := tool.GetCfg().Email
	code := tool.RandCode()
	auth := smtp.PlainAuth("", emailCfg.ServiceEmail, emailCfg.ServicePwd, emailCfg.SmtpHost)
	to := []string{email}

	fmt.Println("SEND EMAIL TO :", email)

	str := fmt.Sprintf("From:%v\r\nTo:%v\r\nSubject:豆瓣验证码\r\n\r\n您的验证码为：%s\r\n请在5分钟内完成验证", emailCfg.ServiceEmail, email, code)
	msg := []byte(str)
	err := smtp.SendMail(emailCfg.SmtpHost+":"+emailCfg.SmtpPort, auth, emailCfg.ServiceEmail, to, msg)
	if err != nil {
		return "", err
	}
	return code, err
}

//修改邮箱or绑定邮箱
func ChangeEmail(email string, id int64) error {
	err := dao.ChangeEmail(email, id)
	return err
}

//修改电话号码
func ChangePhone(phone string, id int64) error {
	err := dao.ChangePhone(phone, id)
	return err
}

//redis 校验

//将验证码存储到redis
func PutCodeInRedis(ctx *gin.Context, key, value string) error {
	err := dao.SetRedisValue(ctx, key, value)
	return err
}

//校验验证码
func JudgeVerifyCode(ctx *gin.Context, key string, givenValue string) (bool, error) {
	value, err := dao.GetRedisValue(ctx, key)
	if err != nil {
		return false, err
	}

	if value != givenValue {
		return false, nil
	}

	return true, nil
}

// USER 插入

//将user信息插入MySQL
func InsertUser(user model.User) error {
	err := dao.InsertUser(user)
	return err
}
