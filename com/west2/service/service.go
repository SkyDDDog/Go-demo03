package service

import "gorm.io/gorm"

type ServiceStruct struct {
	*gorm.DB //ecgo为框架的包名
}

var Service *ServiceStruct

func NewService(db *gorm.DB) *ServiceStruct {
	return &ServiceStruct{db}
}

func InitService(db *gorm.DB) {
	Service = NewService(db)
}
