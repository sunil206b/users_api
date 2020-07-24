package repo

import (
	"github.com/sunil206b/users_api/model"
	"github.com/sunil206b/users_api/utils/errors"
)

type IUserRepo interface {
	CreateUser(user *model.User) *errors.RestErr
	GetUser(userId int64) (*model.User, *errors.RestErr)
	UpdateUser(user *model.User) *errors.RestErr
	DeleteUser(userId int64) (bool, *errors.RestErr)
}
