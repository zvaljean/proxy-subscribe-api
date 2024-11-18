package entity

import (
	"gorm.io/gorm"
)

type UserDto struct {
	Type  int
	Token string
	Path  string
}

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
