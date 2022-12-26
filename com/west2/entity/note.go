package entity

import (
	entity "demo03/com/west2/entity/base"
	"time"
)

type Note struct {
	//gorm.Model
	entity.Model
	Title   string    `gorm:"comment:标题;type:varchar(255)"`
	Content string    `gorm:"comment:内容;type:varchar(255)"`
	State   string    `gorm:"comment:完成状态(1已完成|0未完成);type:char"`
	EndTime time.Time `gorm:"comment:截止时间;datetime"`
}

func NewNote() Note {
	return Note{
		Model: entity.InitModel(),
	}
}
