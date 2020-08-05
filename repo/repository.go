package repo

import (
	"github.com/sunil206b/store_utils_go/errors"
	"github.com/sunil206b/users_api/model"
)

type IUserRepo interface {
	CreateUser(user *model.User) *errors.RestErr
	GetUser(userId int64) (*model.User, *errors.RestErr)
	UpdateUser(user *model.User) *errors.RestErr
	DeleteUser(userId int64) *errors.RestErr
	Search(status string) ([]model.User, *errors.RestErr)
	FindByEmail(email string) (*model.LoginUser, *errors.RestErr)
}
