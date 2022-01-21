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
	engine.GET("/api/user/get_user/:uid", getUerInfo)
	engine.POST("/api/user/get_user/:uid", accountManagement)

	engine.POST("/api/verify/sms", sendSms)
	engine.POST("/api/user/register", register)
	engine.POST("/api/user/login/sms", loginBySms)
	engine.POST("/api/user/login/pw", login)
	engine.POST("/api/verify/emial", SendEmail)
	engine.PUT("/api/user/bind_email", bindEmail)
	engine.PUT("/api/user/unbind_email", unbindEmail)
}

//User的账号设置

//获取个人界面信息
func getUerInfo(ctx *gin.Context) {
	//获取用户ID
	Id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("getUerInfo_ParseInt is Err: ", err)
		tool.RespErrorWithData(ctx, "UID无效")
		return
	}

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

//发送邮箱
func SendEmail(ctx *gin.Context) {
	//将大写转化成小写提升用户体验
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

	if emailParam.Email == "" {
		tool.RespErrorWithData(ctx, "邮箱不能为空")
		return
	}

	if err != nil {
		if err.Error()[:12] == "Key: 'Email." {
			tool.RespErrorWithData(ctx, "缺少必要参数")
			fmt.Println("changeEmail_ParesePara ERR is :", err)
			return
		}
		tool.RespInternalError(ctx)
		fmt.Println("binEmail_ParesePara ERR is :", err)
		return
	}

	//获取token来得到用户信息
	//token只能来获取不可修改的用户内容
	claims, err := service.ParseAccessToken(emailParam.Token)
	flag := tool.CheckToken(ctx, claims, err)
	if !flag {
		return
	}

	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("bindEmail_ParseAccessToken ERR is :", err)
		return
	}

	//校验验证码
	flag, err = service.JudgeVerifyCode(ctx, emailParam.Email, emailParam.VerifyCode)

	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("ChangeEmail_JudgeVerifyCode ERR is :", err)
		return
	}

	if flag {
		err := service.ChangeEmail(emailParam.Email, claims.User.Id)
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

	//获取token来得到用户id 如果是直接token来获取用户信息的话 token也要跟新
	claims, err := service.ParseAccessToken(emailParam.Token)
	flag := tool.CheckToken(ctx, claims, err)

	if !flag {
		return
	}

	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("unbindEmail_ParseAccessToken ERR is :", err)
		return
	}
	//通过id来得到用户信息
	User, bool, err := service.JudgeAndQueryUserByUserID(claims.User.Id)

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
	flag, err = service.JudgeVerifyCode(ctx, User.Email, emailParam.VerifyCode)

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
		err := service.ChangeEmail("", claims.User.Id)
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
