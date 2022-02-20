package controller

import (
	"douban/model"
	"douban/param"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
}

func (u *UserController) Router(engine *gin.Engine) {

	engine.GET("/api/user/get_user", getUerInfo)
	engine.POST("/api/user/get_user/:uid", JWTAuthMiddleware(), accountManagement)

	engine.POST("/api/verify/sms", sendSms)
	engine.POST("/api/user/register", register)
	engine.POST("/api/user/login/sms", loginBySms)
	engine.POST("/api/user/login/pw", login)
	engine.POST("/api/verify/emial", SendEmail)

	engine.PUT("/api/user/unbind_phone", unbindPhone)
	engine.PUT("/api/user/bind_phone", bindPhone)
	engine.PUT("/api/user/bind_email", bindEmail)
	engine.PUT("/api/user/unbind_email", JWTAuthMiddleware(), unbindEmail)

	engine.PUT("/api/user/change_habitat", changeHabitat)
	engine.PUT("/api/user/change_account", changeAccount)
	engine.PUT("/api/user/change_avatar", changeAvatar)

	engine.DELETE("/api/user/suicide", suicideAccount)
}

//User的账号设置
//获取个人界面信息

func getUerInfo(ctx *gin.Context) {
	//获取用户ID
	accessToken := ctx.PostForm("access_token")
	refreshToken := ctx.PostForm("refresh_token")

	if accessToken != "" && refreshToken != "" {
		Claims, flag, err := service.ParseToken(accessToken, refreshToken)

		if err != nil {
			tool.RespErrorWithData(ctx, "token错误")
			fmt.Println("err", err)
			return
		}
		if flag {
			accessToken, err := service.GenToken(Claims.User, 3600*24, "ACCESS_TOKEN")
			if err != nil {
				fmt.Println("JWTAuthMiddleware_CreateAccessTokenErr:", err)
				tool.RespInternalError(ctx)
				return
			}

			//refreshToken 一周
			refreshToken, err := service.GenToken(Claims.User, 604800, "REFRESH_TOKEN")
			if err != nil {
				fmt.Println("JWTAuthMiddleware_CreateRefreshTokenErr:", err)
				tool.RespInternalError(ctx)
				return
			}

			UserInfo, err := service.GetUserById(Claims.User.Id)
			if err != nil {
				fmt.Println("GetUserById Is err:", err)
				tool.RespInternalError(ctx)
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"userInfo":      UserInfo,
				"status":        "true",
				"refresh_token": refreshToken,
				"access_token":  accessToken,
			})

		} else {
			UserInfo, err := service.GetUserById(Claims.User.Id)
			if err != nil {
				fmt.Println("GetUserById Is err:", err)
				tool.RespInternalError(ctx)
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"userInfo":      UserInfo,
				"status":        "true",
				"refresh_token": refreshToken,
				"access_token":  accessToken,
			})
		}
	} else {
		tool.RespErrorWithData(ctx, "请重新登录")
	}
}

//账号管理
func accountManagement(ctx *gin.Context) {

}

//修改头像
func changeAvatar(ctx *gin.Context) {

	img, err := ctx.FormFile("avatar")

	Id, err := strconv.ParseInt(ctx.PostForm("id"), 10, 64)

	if err != nil {
		fmt.Println("FormFileErr: ", err)
		tool.RespErrorWithData(ctx, "上传失败")
		return
	}

	//大小限制2Mb
	if img.Size > (2 << 20) {
		tool.RespErrorWithData(ctx, "头像文件过大")
		return
	}
	file, err := img.Open()
	//判断文件格式
	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println(err)
	}

	fileFormat := tool.GetFileType(fileByte)

	if fileFormat == "" {
		tool.RespErrorWithData(ctx, "头像格式无效")
		return
	}

	filePath := "/avatar/" + strconv.FormatInt(Id, 10) + "." + fileFormat

	//上传头像
	err = service.UploadAvatar(file, filePath)
	if err != nil {
		fmt.Println("UploadAvatarErr: ", err)
		tool.RespInternalError(ctx)
		return
	}

	cfg := tool.GetCfg().Cos
	url := cfg.AvatarUrl + filePath

	//头像入数据库
	err = service.ChangeAvatar(url, Id)
	if err != nil {
		fmt.Println("changeAvatar_ChangeAvatar Err is: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, "修改成功")

}

//修改常常住地
func changeHabitat(ctx *gin.Context) {

	Id := ctx.MustGet("id").(int64)
	tool.CatchPanic(ctx, "changeHabitat")
	Habitat := ctx.PostForm("habitat")

	err := service.ChangeHabitat(Habitat, Id)
	if err != nil {
		tool.RespErrorWithData(ctx, "修改常住地失败")
		fmt.Println("changeHabitat_ChangeHabitat ERR is", err)
		return
	}
	tool.RespSuccessfulWithData(ctx, "修改成功")
}

//修改账户信息
func changeAccount(ctx *gin.Context) {
	var accountParam param.Account
	err := ctx.ShouldBind(&accountParam)
	if err != nil {
		tool.RespErrorWithData(ctx, "参数解析失败")
		return
	}
	accessToken := ctx.PostForm("access_token")
	refreshToken := ctx.PostForm("refresh_token")

	if accessToken != "" && refreshToken != "" {
		Claims, flag, err := service.ParseToken(accessToken, refreshToken)

		if err != nil {
			tool.RespErrorWithData(ctx, "token错误")
			fmt.Println("err", err)
			return
		}
		if flag {
			accessToken, err := service.GenToken(Claims.User, 3600*24, "ACCESS_TOKEN")
			if err != nil {
				fmt.Println("JWTAuthMiddleware_CreateAccessTokenErr:", err)
				tool.RespInternalError(ctx)
				return
			}

			//refreshToken 一周
			refreshToken, err := service.GenToken(Claims.User, 604800, "REFRESH_TOKEN")
			if err != nil {
				fmt.Println("JWTAuthMiddleware_CreateRefreshTokenErr:", err)
				tool.RespInternalError(ctx)
				return
			}
			if accountParam.NewName != "" {
				err := service.ChangeUserName(accountParam.NewName, Claims.User.Id)
				if err != nil {
					tool.RespErrorWithData(ctx, "修改失败")
					fmt.Println("changeAccount_ChangeUserName is ERR :", err)
					return
				}
			}

			if accountParam.Hometown != "" {
				err := service.ChangeHometown(accountParam.Hometown, Claims.User.Id)
				if err != nil {
					tool.RespErrorWithData(ctx, "修改失败")
					return
				}
			}

			if accountParam.Birthday != "" {
				Birthday, err := time.ParseInLocation("2006-01-02", accountParam.Birthday, time.Local)
				if err != nil {
					fmt.Println("changeAccount_ParseInLocationErr: ", err)
					tool.RespErrorWithData(ctx, "日期格式错误")
					return
				}
				err = service.ChangeBirthday(Birthday, Claims.User.Id)
				if err != nil {
					tool.RespErrorWithData(ctx, "修改失败")
					return
				}
			}

			if accountParam.HometownPublic != "" {
				err := service.ChangeHometownPubic(accountParam.HometownPublic, Claims.User.Id)
				if err != nil {
					fmt.Println("changeAccount_ChangeHometownPubic ERR", err)
					tool.RespErrorWithData(ctx, "修改失败")
					return
				}

			}

			if accountParam.BirthdayPublic != "" {
				err := service.ChangeBirthdayPublic(accountParam.BirthdayPublic, Claims.User.Id)
				if err != nil {
					fmt.Println("changeAccount_ChangeBirthdayPublic ERR", err)
					tool.RespErrorWithData(ctx, "修改失败")
					return
				}
			}

			ctx.JSON(http.StatusOK, gin.H{
				"data":          "修改成功",
				"status":        "true",
				"refresh_token": refreshToken,
				"access_token":  accessToken,
			})

		} else {
			if accountParam.NewName != "" {
				err := service.ChangeUserName(accountParam.NewName, Claims.User.Id)
				if err != nil {
					tool.RespErrorWithData(ctx, "修改失败")
					fmt.Println("changeAccount_ChangeUserName is ERR :", err)
					return
				}
			}

			if accountParam.Hometown != "" {
				err := service.ChangeHometown(accountParam.Hometown, Claims.User.Id)
				if err != nil {
					tool.RespErrorWithData(ctx, "修改失败")
					return
				}
			}

			if accountParam.Birthday != "" {
				Birthday, err := time.ParseInLocation("2006-01-02", accountParam.Birthday, time.Local)
				if err != nil {
					fmt.Println("changeAccount_ParseInLocationErr: ", err)
					tool.RespErrorWithData(ctx, "日期格式错误")
					return
				}
				err = service.ChangeBirthday(Birthday, Claims.User.Id)
				if err != nil {
					tool.RespErrorWithData(ctx, "修改失败")
					return
				}
			}

			if accountParam.HometownPublic != "" {
				err := service.ChangeHometownPubic(accountParam.HometownPublic, Claims.User.Id)
				if err != nil {
					fmt.Println("changeAccount_ChangeHometownPubic ERR", err)
					tool.RespErrorWithData(ctx, "修改失败")
					return
				}

			}

			if accountParam.BirthdayPublic != "" {
				err := service.ChangeBirthdayPublic(accountParam.BirthdayPublic, Claims.User.Id)
				if err != nil {
					fmt.Println("changeAccount_ChangeBirthdayPublic ERR", err)
					tool.RespErrorWithData(ctx, "修改失败")
					return
				}
			}
			ctx.JSON(http.StatusOK, gin.H{
				"data":          "修改成功",
				"status":        "true",
				"refresh_token": refreshToken,
				"access_token":  accessToken,
			})
		}
	} else {
		tool.RespErrorWithData(ctx, "请重新登录")
	}

}

//注销账号
func suicideAccount(ctx *gin.Context) {

	accessToken := ctx.PostForm("access_token")
	refreshToken := ctx.PostForm("refresh_token")

	if accessToken != "" && refreshToken != "" {
		Claims, flag, err := service.ParseToken(accessToken, refreshToken)

		if err != nil {
			tool.RespErrorWithData(ctx, "token错误")
			fmt.Println("err", err)
			return
		}
		if flag {
			accessToken, err := service.GenToken(Claims.User, 3600*24, "ACCESS_TOKEN")
			if err != nil {
				fmt.Println("JWTAuthMiddleware_CreateAccessTokenErr:", err)
				tool.RespInternalError(ctx)
				return
			}

			//refreshToken 一周
			refreshToken, err := service.GenToken(Claims.User, 604800, "REFRESH_TOKEN")
			if err != nil {
				fmt.Println("JWTAuthMiddleware_CreateRefreshTokenErr:", err)
				tool.RespInternalError(ctx)
				return
			}

			err = service.DeleteAccount(Claims.User.Id)
			if err != nil {
				tool.RespErrorWithData(ctx, "注销失败")
				fmt.Println("suicideAccount_DeleteAccount  is ERR", err)
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"data":          "注销成功",
				"status":        "true",
				"refresh_token": refreshToken,
				"access_token":  accessToken,
			})

		} else {
			err := service.DeleteAccount(Claims.User.Id)
			if err != nil {
				tool.RespErrorWithData(ctx, "注销失败")
				fmt.Println("suicideAccount_DeleteAccount  is ERR", err)
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"data":          "注销成功",
				"status":        "true",
				"refresh_token": refreshToken,
				"access_token":  accessToken,
			})
		}
		return
	} else {
		tool.RespErrorWithData(ctx, "请重新登录")
	}
	tool.CatchPanic(ctx, "suicideAccount")

}

//解绑手机号
func unbindPhone(ctx *gin.Context) {
	var phoneParam param.Phone
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

	accessToken := ctx.PostForm("access_token")
	refreshToken := ctx.PostForm("refresh_token")

	if accessToken != "" && refreshToken != "" {
		Claims, flag, err := service.ParseToken(accessToken, refreshToken)

		if err != nil {
			tool.RespErrorWithData(ctx, "token错误")
			fmt.Println("err", err)
			return
		}
		if flag {
			accessToken, err := service.GenToken(Claims.User, 3600*24, "ACCESS_TOKEN")
			if err != nil {
				fmt.Println("JWTAuthMiddleware_CreateAccessTokenErr:", err)
				tool.RespInternalError(ctx)
				return
			}

			//refreshToken 一周
			refreshToken, err := service.GenToken(Claims.User, 604800, "REFRESH_TOKEN")
			if err != nil {
				fmt.Println("JWTAuthMiddleware_CreateRefreshTokenErr:", err)
				tool.RespInternalError(ctx)
				return
			}
			User, bool, err := service.JudgeAndQueryUserByUserID(Claims.User.Id)

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

			err = service.ChangePhone("", Claims.User.Id)
			if err != nil {
				tool.RespErrorWithData(ctx, "解绑失败")
				fmt.Println("unbindPhone_ChangePhone ERR is ", err)
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"data":          "解绑成功",
				"status":        "true",
				"refresh_token": refreshToken,
				"access_token":  accessToken,
			})

		} else {
			User, bool, err := service.JudgeAndQueryUserByUserID(Claims.User.Id)

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

			err = service.ChangePhone("", Claims.User.Id)
			if err != nil {
				tool.RespErrorWithData(ctx, "解绑失败")
				fmt.Println("unbindPhone_ChangePhone ERR is ", err)
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"data":          "解绑成功",
				"status":        "true",
				"refresh_token": refreshToken,
				"access_token":  accessToken,
			})
		}
	} else {
		tool.RespErrorWithData(ctx, "请重新登录")
	}

}

//绑定手机号
func bindPhone(ctx *gin.Context) {
	var phoneParam param.Phone
	err := ctx.ShouldBind(&phoneParam)

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
	accessToken := ctx.PostForm("access_token")
	refreshToken := ctx.PostForm("refresh_token")

	if accessToken != "" && refreshToken != "" {
		Claims, flag, err := service.ParseToken(accessToken, refreshToken)

		if err != nil {
			tool.RespErrorWithData(ctx, "token错误")
			fmt.Println("err", err)
			return
		}
		if flag {
			accessToken, err := service.GenToken(Claims.User, 3600*24, "ACCESS_TOKEN")
			if err != nil {
				fmt.Println("JWTAuthMiddleware_CreateAccessTokenErr:", err)
				tool.RespInternalError(ctx)
				return
			}

			//refreshToken 一周
			refreshToken, err := service.GenToken(Claims.User, 604800, "REFRESH_TOKEN")
			if err != nil {
				fmt.Println("JWTAuthMiddleware_CreateRefreshTokenErr:", err)
				tool.RespInternalError(ctx)
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
				err = service.ChangePhone(phoneParam.Phone, Claims.User.Id)
				if err != nil {
					fmt.Println("bindPhone_ChangePhone ERR is :", err)
					tool.RespInternalError(ctx)
					return
				}
				ctx.JSON(http.StatusOK, gin.H{
					"data":          "绑定成功",
					"status":        "true",
					"refresh_token": refreshToken,
					"access_token":  accessToken,
				})
				return
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"data":          "验证码错误或过期",
					"status":        "true",
					"refresh_token": refreshToken,
					"access_token":  accessToken,
				})
				return
			}

		} else {
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
				err = service.ChangePhone(phoneParam.Phone, Claims.User.Id)
				if err != nil {
					fmt.Println("bindPhone_ChangePhone ERR is :", err)
					tool.RespInternalError(ctx)
					return
				}
				ctx.JSON(http.StatusOK, gin.H{
					"data":          "绑定成功",
					"status":        "true",
					"refresh_token": refreshToken,
					"access_token":  accessToken,
				})
				return
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"data":          "验证码错误或过期",
					"status":        "true",
					"refresh_token": refreshToken,
					"access_token":  accessToken,
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"data":          "验证码错误或过期",
				"status":        "true",
				"refresh_token": refreshToken,
				"access_token":  accessToken,
			})
			return
		}
	} else {
		tool.RespErrorWithData(ctx, "请重新登录")
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

	accessToken := ctx.PostForm("access_token")
	refreshToken := ctx.PostForm("refresh_token")

	if accessToken != "" && refreshToken != "" {
		Claims, flag, err := service.ParseToken(accessToken, refreshToken)

		if err != nil {
			tool.RespErrorWithData(ctx, "token错误")
			fmt.Println("err", err)
			return
		}
		if flag {
			accessToken, err := service.GenToken(Claims.User, 3600*24, "ACCESS_TOKEN")
			if err != nil {
				fmt.Println("JWTAuthMiddleware_CreateAccessTokenErr:", err)
				tool.RespInternalError(ctx)
				return
			}

			//refreshToken 一周
			refreshToken, err := service.GenToken(Claims.User, 604800, "REFRESH_TOKEN")
			if err != nil {
				fmt.Println("JWTAuthMiddleware_CreateRefreshTokenErr:", err)
				tool.RespInternalError(ctx)
				return
			}

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
					ctx.JSON(http.StatusOK, gin.H{
						"data":          "已经被其他账户绑定",
						"status":        "true",
						"refresh_token": refreshToken,
						"access_token":  accessToken,
					})
					return
				}
				err = service.ChangeEmail(emailParam.Email, id)
				if err != nil {
					fmt.Println("bindEmail_ChangeEmail ERR is :", err)
					tool.RespInternalError(ctx)
					return
				}
				ctx.JSON(http.StatusOK, gin.H{
					"data":          "绑定成功",
					"status":        "true",
					"refresh_token": refreshToken,
					"access_token":  accessToken,
				})
				return
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"data":          "绑定失败",
					"status":        "true",
					"refresh_token": refreshToken,
					"access_token":  accessToken,
				})
				return
			}
		} else {
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
					ctx.JSON(http.StatusOK, gin.H{
						"data":          "已经被其他账户绑定",
						"status":        "true",
						"refresh_token": refreshToken,
						"access_token":  accessToken,
					})
					return
				}
				err = service.ChangeEmail(emailParam.Email, id)
				if err != nil {
					fmt.Println("bindEmail_ChangeEmail ERR is :", err)
					tool.RespInternalError(ctx)
					return
				}
				ctx.JSON(http.StatusOK, gin.H{
					"data":          "绑定成功",
					"status":        "true",
					"refresh_token": refreshToken,
					"access_token":  accessToken,
				})
				return
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"data":          "绑定失败",
					"status":        "true",
					"refresh_token": refreshToken,
					"access_token":  accessToken,
				})
				return
			}
		}

	} else {
		tool.RespErrorWithData(ctx, "请重新登录")
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
		accessToken, err := service.GenToken(User, 60*24, "ACCESS_TOKEN")
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
			"token":         accessToken + " " + refreshToken,
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
		accessToken, err := service.GenToken(User, 60*24, "ACCESS_TOKEN")
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
			"token":         accessToken + " " + refreshToken,
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
