package controller

import (
	"demo03/com/west2/common"
	"github.com/gin-gonic/gin"
)

type TestController struct {
}

// 注册接口
func TestControllerRegister(testGrp *gin.RouterGroup) {
	testController := &TestController{}
	testGrp.Use().GET("/ping", testController.Ping)
}

//	@Summary		Ping!
//	@Description	测试
//	@Tags			TestController
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string
//	@Router			/test/ping [get]
func (controller TestController) Ping(c *gin.Context) {
	common.Success("pong", c)
}
