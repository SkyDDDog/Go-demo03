package route

import (
	"demo03/com/west2/controller"
	_ "demo03/com/west2/docs"
	"demo03/com/west2/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func PathRoute(r *gin.Engine) *gin.Engine {

	rootPath := r.Group("/api")
	{
		testPath := rootPath.Group("/test")
		{
			controller.TestControllerRegister(testPath)
		}
		userPath := rootPath.Group("/user")
		{
			controller.UserControllerRegister(userPath)
		}

		notePath := rootPath.Group("/note")
		notePath.Use(handler.TokenHanlder())
		{
			controller.NoteControllerRegister(notePath)
		}
	}

	return r
}

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// 载入跨域配置
	r.Use(Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	PathRoute(r)
	return r
}
