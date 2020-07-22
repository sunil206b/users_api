package service

import (
	"database/sql"
	"github.com/sunil206b/users_api/dto"
	"github.com/sunil206b/users_api/model"
	"github.com/sunil206b/users_api/repo"
	"github.com/sunil206b/users_api/utils/errors"
)

type UserService struct {
	repo repo.IUserRepo
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		repo: repo.NewUserRepo(db),
	}
}

func (u *UserService) CreateUse(userDTO dto.UserDTO) (*dto.UserDTO, *errors.RestErr) {
	errMsg := userDTO.Validate()
	if errMsg != nil {
		 return nil, errMsg
	}
	var user model.User
	userDTO.CopyToUser(&user)
	if err := u.repo.CreateUser(&user); err != nil {
		errMsg = errors.NewBadRequest(err.Error())
		return nil, errMsg
	}
	userDTO.Id = user.Id
	return &userDTO, nil
}