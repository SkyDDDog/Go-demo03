package entity

import entity "demo03/com/west2/entity/base"

type User struct {
	//gorm.Model
	entity.Model
	Username string `gorm:"comment:用户名;type:varchar(255)"`
	Password string `gorm:"comment:密码;type:varchar(255)"`
}

func NewUser() User {
	return User{
		Model: entity.InitModel(),
	}
}
