package service

import "demo03/com/west2/entity"

func (this *ServiceStruct) CreateUser(user entity.User) {
	this.DB.Create(&user)
}

func (this *ServiceStruct) LoginByUsername(username string) entity.User {
	user := entity.User{}
	this.DB.Where("username = (?)", username).Take(&user)
	return user
}
