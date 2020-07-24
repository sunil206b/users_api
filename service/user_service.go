package service

import (
	"database/sql"
	"github.com/sunil206b/users_api/dto"
	"github.com/sunil206b/users_api/model"
	"github.com/sunil206b/users_api/repo"
	"github.com/sunil206b/users_api/utils/date"
	"github.com/sunil206b/users_api/utils/errors"
	"time"
)

type UserService struct {
	repo repo.IUserRepo
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		repo: repo.NewUserRepo(db),
	}
}

func (us *UserService) CreateUse(userDTO dto.UserDTO) (*dto.UserDTO, *errors.RestErr) {
	errMsg := userDTO.Validate()
	if errMsg != nil {
		 return nil, errMsg
	}
	var user model.User
	userDTO.CopyToUser(&user)
	age, err := date.ConvertToDate(userDTO.Birth)
	if err != nil {
		errMsg = errors.NewBadRequest(err.Error())
	}
	user.Birth = age
	user.DateCreated = time.Now()
	if err := us.repo.CreateUser(&user); err != nil {
		return nil, err
	}
	userDTO.Id = user.Id
	userDTO.DateCreated = date.GetFmtDate(user.DateCreated)
	userDTO.DateUpdated = date.GetFmtDate(user.DateUpdated)
	return &userDTO, nil
}

func (us *UserService) GetUser(userId int64) (*dto.UserDTO, *errors.RestErr) {
	user, err := us.repo.GetUser(userId)
	if err != nil {
		return nil, err
	}
	var userDto dto.UserDTO
	userDto.CopyToDTO(user)
	return &userDto, nil
}