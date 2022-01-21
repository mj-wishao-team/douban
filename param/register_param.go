package param

type Register struct {
	Username string `form:"username"`
	Pwd      string `form:"password"`
	Phone    string `form:"phone" binding:"required"`
}
