package request

type StuRegister struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Phone    string `form:"phone" binding:"required"`
	Gender   string `form:"gender" binding:"required"`
}
