package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: 500,
		Msg:  "请求失败",
		Data: data,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: 200,
		Msg:  "请求成功",
		Data: data,
	})
}
