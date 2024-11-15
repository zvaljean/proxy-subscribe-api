package model

import (
	"gorm.io/gorm"
	"valjean/proxy/subscribe/pkg/common/errno"
	. "valjean/proxy/subscribe/pkg/entity"
)

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db: db}
}

func (u *UserModel) CreateItem(item *User) (bool, *errno.BizCode) {

	result := u.db.Create(item)
	if err := result.Error; err != nil {
		return false, errno.ErrDBQuery
	}
	return true, errno.OK

}

func (u *UserModel) FindItemByToken(token string) (*User, error) {

	dst := &User{}
	result := u.db.Model(&User{}).Where("token = ?", token).Find(&dst)

	if err := result.Error; err != nil {
		return nil, err
	}
	return dst, nil
}

func (u *UserModel) FindAll() ([]User, error) {

	var users []User
	result := u.db.Find(&users)
	if err := result.Error; err != nil {
		return nil, err
	}

	return users, nil
}
