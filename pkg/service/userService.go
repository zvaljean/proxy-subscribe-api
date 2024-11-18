package service

import (
	"zvaljean/proxy-subscribe-api/pkg/common/errno"
	"zvaljean/proxy-subscribe-api/pkg/entity"
	"zvaljean/proxy-subscribe-api/pkg/log"
	"zvaljean/proxy-subscribe-api/pkg/model"
)

type UserService struct {
	userModel *model.UserModel
}

func NewUserService(userModel *model.UserModel) *UserService {
	return &UserService{userModel: userModel}
}

func (svc *UserService) FindUserByTokenTypePath(param *entity.UserDto) (*entity.User, *errno.BizCode) {

	user, err := svc.userModel.FindUserByTokenTypePath(param)
	if log.ErrorCheck(err, "userModel.FindUserByToken") {
		return nil, errno.ErrDBQuery
	}

	if user == nil {
		return nil, errno.ErrDataNotExist
	}

	return user, nil

}
func (svc *UserService) FindUserByToken(token string) (*entity.User, *errno.BizCode) {

	user, err := svc.userModel.FindUserByToken(token)
	if log.ErrorCheck(err, "userModel.FindUserByToken") {
		return nil, errno.ErrDBQuery
	}

	if user == nil {
		return nil, errno.ErrDataNotExist
	}

	return user, nil

}
