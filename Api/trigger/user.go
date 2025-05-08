package trigger

import (
	"Api/handler"
	"Api/request"
	"Api/response"
	"github.com/gin-gonic/gin"
	"stu_srv/proto_stu/student"
)

func StuRegister(c *gin.Context) {
	var data request.StuRegister
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	register, err := handler.StuRegister(c, &student.StuRegisterRequest{
		Username: data.Username,
		Password: data.Password,
		Gender:   data.Gender,
		Phone:    data.Phone,
		Email:    data.Email,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, register)
}
