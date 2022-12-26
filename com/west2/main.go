package main

import (
	route "demo03/com/west2/common/route"
	"demo03/com/west2/config"
	log "demo03/com/west2/config/log"
	"demo03/com/west2/config/toml"
	_ "demo03/com/west2/docs"
	"demo03/com/west2/entity"
	"demo03/com/west2/service"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

//	@title			小黄老师备忘录
//	@version		1.0
//	@description	West2-Golang-Demo03
//	@termsOfService	https://github.com/SkyDDDog/demo03

//	@contact.name	二火
//	@contact.url	https://github.com/SkyDDDog
//	@contact.email	lear_yd@qq.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		127.0.0.1:1234
// @BasePath	/api
func main() {
	// 初始化日志配置
	log.InitLogger(toml.GetConfig().Log.Path, toml.GetConfig().Log.Level)
	// 初始化mysql配置
	db := config.InitDb()
	db.AutoMigrate(&entity.User{})
	// 初始化路由配置
	router := route.NewRouter()
	// 初始化service层
	service.InitService(db)
	// 启动微服务
	router.Run(":1234")
}
