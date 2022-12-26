package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ERROR   = -1
	SUCCESS = 0
)

// 返回的对象
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Result{
		Code: SUCCESS,
		Msg:  "success",
		Data: data,
	})
}
func CommonSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, Result{
		Code: SUCCESS,
		Msg:  "success",
		Data: "",
	})
}

func Fail(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Result{
		Code: ERROR,
		Msg:  msg,
		Data: nil,
	})
}

func TokenInvalid() string {
	return "{code: 0, msg: \"token鉴权失败\", data: }"
}
