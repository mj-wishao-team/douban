package param

type Phone struct {
	VerifyCode string `form:"verify_code" binding:"required"`
	Phone      string `form:"phone"`
}
