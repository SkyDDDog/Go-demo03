package handler

import (
	"demo03/com/west2/common"
	"demo03/com/west2/config/log"
	"demo03/com/west2/util"
	"github.com/gin-gonic/gin"
)

func TokenHanlder() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		log.Logger.Info("access-token: " + token)
		if util.TokenValid(token) {
			c.Next()
		} else {
			c.Abort()
			common.Fail("token校验失败", c)
		}
	}
}
