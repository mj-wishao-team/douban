package param

type Email struct {
	VerifyCode string `form:"verify_code" binding:"required"`
	Email      string `form:"email"`
	Pwd        string `form:"password"`
}
