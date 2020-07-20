package service

import (
	"database/sql"
	"github.com/sunil206b/users_api/model"
	"github.com/sunil206b/users_api/repo"
)

type UserService struct {
	repo *repo.UserRepo
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		repo: repo.NewUserRepo(db),
	}
}

func (u *UserService) CreateUser(user model.User) (*model.User, error) {
	return &user, nil
}