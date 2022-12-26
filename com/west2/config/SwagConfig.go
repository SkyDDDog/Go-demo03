package config

import "demo03/com/west2/docs"

func InitSwag() {
	docs.SwaggerInfo.Title = "小黄老师备忘录"
	docs.SwaggerInfo.Description = "West2-Golang-Demo03"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:1234"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
