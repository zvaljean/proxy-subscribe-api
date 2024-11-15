package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string
	Type   int
	Token  string `gorm:"index"`
	Path   string
	Data   string
	Remark string
}

func (u *User) TableName() string {
	return "user"
}
