package controller

import (
	"douban/model"
	"douban/param"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
}

func (u *UserController) Router(engine *gin.Engine) {

	engine.GET("/api/user/get_user/:uid", JWTAuthMiddleware(), getUerInfo)
	engine.POST("/api/user/get_user/:uid", JWTAuthMiddleware(), accountManagement)

	engine.POST("/api/verify/sms", sendSms)
	engine.POST("/api/user/register", register)
	engine.POST("/api/user/login/sms", loginBySms)
	engine.POST("/api/user/login/pw", login)
	engine.POST("/api/verify/emial", SendEmail)

	engine.PUT("/api/user/unbind_phone", JWTAuthMiddleware(), unbindPhone)
	engine.PUT("/api/user/bind_phone", JWTAuthMiddleware(), bindPhone)
	engine.PUT("/api/user/bind_email", JWTAuthMiddleware(), bindEmail)
	engine.PUT("/api/user/unbind_email", JWTAuthMiddleware(), unbindEmail)
	//engine.PUT("/api/user/change_name",JWTAuthMiddleware(),changeName)

	engine.DELETE("/api/user/suicide", JWTAuthMiddleware(), suicideAccount)
}

//User的账号设置
//获取个人界面信息

func getUerInfo(ctx *gin.Context) {
	//获取用户ID
	Id := ctx.MustGet("id").(int64)

	tool.CatchPanic(ctx, "getUerInfo")

	_, flag, err := service.JudgeAndQueryUserByUserID(Id)

	if err != nil {
		fmt.Println("getUerInfo_JudgeUserID is Err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	if !flag {
		tool.RespErrorWithData(ctx, "UID无效")
		return
	}

}

//账号管理
func accountManagement(ctx *gin.Context) {

}

//改变昵称
//30天才能改一次
//func  changeName(ctx *gin.Context){
//	id := ctx.MustGet("id").(int64)
//	tool.CatchPanic(ctx,"suicideAccount")
//	service.changeUserName(id)
//}

//注销账号
func suicideAccount(ctx *gin.Context) {
	id := ctx.MustGet("id").(int64)
	tool.CatchPanic(ctx, "suicideAccount")
	err := service.DeleteAccount(id)
	if err != nil {
		tool.RespErrorWithData(ctx, "注销失败")
		fmt.Println("suicideAccount_DeleteAccount  is ERR", err)
		return
	}
	tool.RespSuccessfulWithData(ctx, "注销成功")
}

//解绑手机号
func unbindPhone(ctx *gin.Context) {
	var phoneParam param.Phone
	id := ctx.MustGet("id").(int64)
	//var id int64=4
	tool.CatchPanic(ctx, "unbindPhone")

	err := ctx.ShouldBind(&phoneParam)

	if err != nil {
		if err.Error()[:12] == "Key: 'Phone." {
			tool.RespErrorWithData(ctx, "参数解析失败")
			fmt.Println("unbindPhone_ShouldBind ERR is", err)
			return
		}
		tool.RespInternalError(ctx)
		fmt.Println("unbindPhone_ParesePara ERR is :", err)
		return
	}

	User, bool, err := service.JudgeAndQueryUserByUserID(id)

	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("unbind_JudgeAndQueryUserByUserID ERR is :", err)
		return
	}
	if !bool {
		tool.RespErrorWithData(ctx, "用户id不存在")
		return
	}

	//校验验证码
	flag, err := service.JudgeVerifyCode(ctx, User.Phone, phoneParam.VerifyCode)

	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("unbindPhone_JudgeVerifyCode ERR is :", err)
		return
	}
	if !flag {
		tool.RespErrorWithData(ctx, "验证码或密码错误")
		return
	}

	err = service.ChangePhone("", id)
	if err != nil {
		tool.RespErrorWithData(ctx, "解绑失败")
		fmt.Println("unbindPhone_ChangePhone ERR is ", err)
		return
	}

	tool.RespSuccessfulWithData(ctx, "解绑成功")
}

//绑定手机号
func bindPhone(ctx *gin.Context) {
	var phoneParam param.Phone
	err := ctx.ShouldBind(&phoneParam)
	id := ctx.MustGet("id").(int64)
	tool.CatchPanic(ctx, "bindPhone")

	if phoneParam.Phone == "" {
		tool.RespErrorWithData(ctx, "手机号不能为空")
		return
	}

	if err != nil {
		if err.Error()[:12] == "Key: 'email." {
			tool.RespErrorWithData(ctx, "缺少必要参数")
			fmt.Println("bindPhone_ParesePara ERR is :", err)
			return
		}
		tool.RespInternalError(ctx)
		fmt.Println("binEmail_ParesePara ERR is :", err)
		return
	}

	//校验验证码
	flag, err := service.JudgeVerifyCode(ctx, phoneParam.Phone, phoneParam.VerifyCode)

	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("bindPhone_JudgeVerifyCode ERR is :", err)
		return
	}

	if flag {
		//判断phone 是否被注册过
		_, bool, err := service.JudgeAndQueryUserByPhone(phoneParam.Phone)
		if err != nil {
			fmt.Println("bindPhone_JudgeAndQueryUserByPhone ERR is :", err)
			tool.RespInternalError(ctx)
			return
		}
		if bool {
			tool.RespErrorWithData(ctx, "手机号已经被其他账号绑定")
			return
		}
		err = service.ChangePhone(phoneParam.Phone, id)
		if err != nil {
			fmt.Println("bindPhone_ChangePhone ERR is :", err)
			tool.RespInternalError(ctx)
			return
		}
		tool.RespSuccessfulWithData(ctx, "绑定成功")
		return
	} else {
		tool.RespErrorWithData(ctx, "验证码错误或者过期")
		return
	}

}

//发送邮箱
func SendEmail(ctx *gin.Context) {
	email := strings.ToLower(ctx.PostForm("email"))

	if !tool.VerifyEmailFormat(email) {
		tool.RespErrorWithData(ctx, "邮箱格式错误")
		return
	}
	code, err := service.SendCodeByEmail(email)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("SendEmail_SenCodeByEmail is ERR :", err)
		return
	}
	//将验证码储存到redis上
	err = service.PutCodeInRedis(ctx, email, code)
	if err != nil {
		tool.RespErrorWithData(ctx, err)
		fmt.Println("SendEmail_PutCodeInRedis is ERR :", err)
		return
	}

	tool.RespSuccessfulWithData(ctx, "发送成功")
}

//绑定邮箱
func bindEmail(ctx *gin.Context) {
	var emailParam param.Email
	err := ctx.ShouldBind(&emailParam)
	id := ctx.MustGet("id").(int64)
	tool.CatchPanic(ctx, "bindEmail")

	if emailParam.Email == "" {
		tool.RespErrorWithData(ctx, "邮箱不能为空")
		return
	}

	if err != nil {
		if err.Error()[:12] == "Key: 'Email." {
			tool.RespErrorWithData(ctx, "缺少必要参数")
			fmt.Println("bindEmail_ParesePara ERR is :", err)
			return
		}
		tool.RespInternalError(ctx)
		fmt.Println("binEmail_ParesePara ERR is :", err)
		return
	}

	//校验验证码
	flag, err := service.JudgeVerifyCode(ctx, emailParam.Email, emailParam.VerifyCode)

	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("bindEmail_JudgeVerifyCode ERR is :", err)
		return
	}

	if flag {
		_, bool, err := service.JudgeAndQueryUserByEmail(emailParam.Email)
		if err != nil {
			fmt.Println("bindEmail_JudgeAndQueryUserByEmail ERR is :", err)
			tool.RespInternalError(ctx)
			return
		}
		if bool {
			tool.RespErrorWithData(ctx, "邮箱已经被其他账号绑定")
			return
		}
		err = service.ChangeEmail(emailParam.Email, id)
		if err != nil {
			fmt.Println("bindEmail_ChangeEmail ERR is :", err)
			tool.RespInternalError(ctx)
			return
		}
		tool.RespSuccessfulWithData(ctx, "绑定成功")
		return
	} else {
		tool.RespErrorWithData(ctx, "验证码错误或者过期")
		return
	}
}

//解绑邮箱
//若该用户没有绑定手机号则不能解绑
//若解绑则默认为有手机号
func unbindEmail(ctx *gin.Context) {
	var emailParam param.Email
	id := ctx.MustGet("id").(int64)

	err := ctx.ShouldBind(&emailParam)

	if err != nil {
		if err.Error()[:12] == "Key: 'Email." {
			tool.RespErrorWithData(ctx, "缺少必要参数")
			fmt.Println("unbbindEmail_ParesePara ERR is :", err)
			return
		}
		tool.RespInternalError(ctx)
		fmt.Println("unbindEmail_ParesePara ERR is :", err)
		return
	}

	if emailParam.Pwd == "" {
		tool.RespErrorWithData(ctx, "密码不能为空")
		return
	}

	User, bool, err := service.JudgeAndQueryUserByUserID(id)

	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("unbind_JudgeAndQueryUserByUserID ERR is :", err)
		return
	}
	if !bool {
		tool.RespErrorWithData(ctx, "用户id不存在")
		return
	}

	//校验验证码
	flag, err := service.JudgeVerifyCode(ctx, User.Email, emailParam.VerifyCode)

	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("unbindEmail_JudgeVerifyCode ERR is :", err)
		return
	}
	if !flag {
		tool.RespErrorWithData(ctx, "验证码或密码错误")
		return
	}

	//根据手机号判断密码是否正确
	_, bool, err = service.JudgePasswordCorrect(User.Phone, emailParam.Pwd)

	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("unbindEmail_JudgePasswordCorrect ERR is :", err)
		return
	}

	if !bool {
		tool.RespErrorWithData(ctx, "验证码或密码错误")
		return
	} else {
		//
		err := service.ChangeEmail("", id)
		if err != nil {
			fmt.Println("unbindEmail_ChangeEmail ERR is :", err)
			tool.RespInternalError(ctx)
			return
		}
		tool.RespSuccessfulWithData(ctx, "解绑成功")
		return
	}
}

//用户注册 登入功能

//密码登录（ps:官网几乎废弃这个，典中典安全隐患）
func login(ctx *gin.Context) {
	loginAccount := ctx.PostForm("loginAccount")
	password := ctx.PostForm("password")

	if loginAccount == "" {
		tool.RespErrorWithData(ctx, "请输入注册时用的邮箱或者手机号")
		return
	}

	if password == "" {
		tool.RespErrorWithData(ctx, "请输入密码")
		return
	}

	User, flag, err := service.JudgePasswordCorrect(loginAccount, password)

	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("UserLogin is", err)
		return
	}
	if !flag {
		tool.RespErrorWithData(ctx, "账号或密码错误")
		return
	}

	if flag {
		// 生成AceessToken 和RefreshToken
		//accessToken 5分钟
		accessToken, err := service.GenToken(User, 300, "ACCESS_TOKEN")
		if err != nil {
			fmt.Println("CreateAccessTokenErr:", err)
			tool.RespInternalError(ctx)
			return
		}

		//refreshToken 一周
		refreshToken, err := service.GenToken(User, 604800, "REFRESH_TOKEN")
		if err != nil {
			fmt.Println("CreateRefreshTokenErr:", err)
			tool.RespInternalError(ctx)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"status":        "ture",
			"data":          User.Id,
		})
		fmt.Println(accessToken + " " + refreshToken)
		return
	}
}

//短信发送
func sendSms(ctx *gin.Context) {
	phone := ctx.PostForm("phone")

	if phone == "" {
		tool.RespSuccessfulWithData(ctx, "电话号码不能为空")
		return
	}
	if !tool.VerifyMobileFormat(phone) {
		tool.RespErrorWithData(ctx, "手机号格式错误")
		return
	}
	//发送短信且返回验证码
	code, err := service.SendSms(phone)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("sendSms is err :", err)
		return
	}
	//将验证码放入redis中
	err = service.PutCodeInRedis(ctx, phone, code)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("PutCodeInRedis is err :", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": phone,
		"info": "短信发送成功",
	})
}

//短信登录
func loginBySms(ctx *gin.Context) {
	phone := ctx.PostForm("phone")
	verifyCode := ctx.PostForm("verify_code")

	if phone == "" {
		tool.RespSuccessfulWithData(ctx, "电话号码不能为空")
		return
	}
	if !tool.VerifyMobileFormat(phone) {
		tool.RespErrorWithData(ctx, "手机号格式错误")
		return
	}

	//校验二维码
	flag, err := service.JudgeVerifyCode(ctx, phone, verifyCode)
	if err != nil {
		if err.Error() == "redis: nil" {
			tool.RespErrorWithData(ctx, "未发送验证码")
			return
		}
		tool.RespInternalError(ctx)
		fmt.Println("JudgeCodeErr: ", err)
		return
	}

	if flag == false {
		tool.RespErrorWithData(ctx, "验证码错误或者过期")
		return
	}
	//判断是否为注册  ture-注册 flase-未注册
	User, flag, err := service.JudgeAndQueryUserByPhone(phone)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println(ctx)
		return
	}

	//登录
	if flag {

		// 生成AceessToken 和RefreshToken
		//accessToken 5分钟
		accessToken, err := service.GenToken(User, 300, "ACCESS_TOKEN")
		if err != nil {
			fmt.Println("CreateAccessTokenErr:", err)
			tool.RespInternalError(ctx)
			return
		}

		//refreshToken 一周
		refreshToken, err := service.GenToken(User, 604800, "REFRESH_TOKEN")
		if err != nil {
			fmt.Println("CreateRefreshTokenErr:", err)
			tool.RespInternalError(ctx)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"status":        "ture",
			"data":          User.Id,
			"info":          "登录成功",
		})
		fmt.Println(accessToken + " " + refreshToken)
	} else {
		//新用户跳转到注册界面
		ctx.JSON(http.StatusOK, gin.H{
			"info":   "新用户",
			"status": "true",
			"data":   phone,
		})
	}
}

//注册
func register(ctx *gin.Context) {
	//解析表单
	var registerParam param.Register
	err := ctx.ShouldBind(&registerParam)

	if err != nil {
		//若前端没有把phone填入则返回解析失败
		tool.RespErrorWithData(ctx, "参数解析失败")
		return
	}

	//判断用户名否注册
	_, flag, err := service.JudgeAndQueryUserByUserName(registerParam.Username)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("register is err: ", err)
		return
	}
	if flag {
		tool.RespErrorWithData(ctx, "该用户名已经注册")
		return
	}

	//同豆瓣标准一致
	if len(registerParam.Username) > 14 {
		tool.RespErrorWithData(ctx, "用户名不能大于14个字符")
		return
	}

	if len(registerParam.Username) == 0 {
		tool.RespErrorWithData(ctx, "用户名不能为空")
		return
	}

	//判断密码是否正确
	if len(registerParam.Pwd) < 6 {
		tool.RespErrorWithData(ctx, " 密码不能小于6个字符")
		return
	}
	//同豆瓣标准一致
	if len(registerParam.Pwd) > 20 {
		tool.RespErrorWithData(ctx, " 密码不能大于20个字符")
		return
	}
	//实例化
	salt := strconv.FormatInt(time.Now().Unix(), 10)
	user := model.User{
		Phone:    registerParam.Phone,
		Username: registerParam.Username,
		Password: tool.HashWithSalted(registerParam.Pwd, salt),
		RegDate:  time.Now(),
		Salt:     salt,
	}

	err = service.InsertUser(user)
	if err != nil {
		fmt.Println("RegisterInsertUserERR: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, "注册成功")

}
