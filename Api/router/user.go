package router

import (
	"Api/trigger"
	"github.com/gin-gonic/gin"
)

func LoadUserRouter(r *gin.Engine) {
	stu := r.Group("./stu")
	{
		stu.POST("/register", trigger.StuRegister)
	}
}
