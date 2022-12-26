package controller

import (
	"demo03/com/west2/common"
	"demo03/com/west2/entity"
	"demo03/com/west2/service"
	"demo03/com/west2/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// 注册接口
func UserControllerRegister(userGrp *gin.RouterGroup) {
	userController := &UserController{}
	userGrp.Use().POST("/register", userController.createUser)
	userGrp.Use().POST("/login", userController.login)
}

// @Summary		创建用户
// @Description	用户注册
// @Tags			UserController
// @Accept			json
// @Produce		json
// @param			string	body		string	false	"用户vo"
// @Success		200		{object}	string
// @Router			/user/register [post]
func (controller UserController) createUser(c *gin.Context) {
	user := entity.NewUser()
	c.ShouldBind(&user)
	hashPassword, _ := util.PasswordHash(user.Password)
	user.Password = hashPassword
	service.Service.CreateUser(user)
	common.CommonSuccess(c)
}

// @Summary		用户登录
// @Description	用户登录
// @Tags			UserController
// @Accept			json
// @Produce		json
// @param			string	body		string	false	"用户vo"
// @Success		200		{object}	string
// @Router			/user/login [POST]
func (controller UserController) login(c *gin.Context) {
	data, _ := c.GetRawData()
	// 定义map或结构体
	var m map[string]interface{}

	// 反序列化
	_ = json.Unmarshal(data, &m)
	username := util.Strval(m["username"])
	rawPassword := util.Strval(m["password"])
	token, refreshToken, _ := util.GenToken(username, rawPassword)
	var result map[string]string
	result = make(map[string]string)
	result["access-token"] = token
	result["refresh-token"] = refreshToken
	loginUser := service.Service.LoginByUsername(username)
	if loginUser.ID == "" || !util.PasswordVerify(rawPassword, loginUser.Password) {
		common.Fail("用户名或密码错误", c)
	} else {
		common.Success(result, c)
	}
}
