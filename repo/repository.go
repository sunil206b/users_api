package repo

import (
	"github.com/sunil206b/users_api/model"
)

type IUserRepo interface {
	CreateUser(user *model.User) error
	GetUser(userId int) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(userId int) (bool, error)
}
