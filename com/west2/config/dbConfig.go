package config

import (
	"demo03/com/west2/config/toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDb() *gorm.DB {
	//myLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	//	logger.Config{
	//		SlowThreshold: time.Second, // 慢 SQL 阈值
	//		LogLevel:                  logger.Silent, // 日志级别
	//		IgnoreRecordNotFoundError: true,  // 忽略ErrRecordNotFound（记录未找到）错误
	//		Colorful:                  false, // 禁用彩色打印
	//	},
	//)
	mysqlConfig := toml.GetConfig().Mysql
	dsn := mysqlConfig.Username + ":" + mysqlConfig.Password + "@tcp(" + mysqlConfig.Host + ":" + mysqlConfig.Port + ")" +
		"/" + mysqlConfig.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: myLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true}},
	)

	return db
}
