package model

import (
	"gorm.io/gorm"
	"zvaljean/proxy-subscribe-api/pkg/common/errno"
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

func (u *UserModel) FindUserByTokenTypePath(param *UserDto) (*User, error) {

	dst := &User{}
	result := u.db.Where(param).Find(&dst)
	//"token = ? and type = ? and path = ?", param.Type, param.Token, param.Path).Find(&dst)

	if result.Error != nil {
		return nil, result.Error
	}
	return dst, nil
}
func (u *UserModel) FindUserByToken(token string) (*User, error) {

	dst := &User{}
	result := u.db.Model(&User{}).Where("token = ?", token).Find(&dst)

	if result.Error != nil {
		return nil, result.Error
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
