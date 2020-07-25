package service

import (
	"database/sql"
	"github.com/sunil206b/users_api/dto"
	"github.com/sunil206b/users_api/model"
	"github.com/sunil206b/users_api/repo"
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
	if errMsg = userDTO.CopyToUser(&user); errMsg != nil {
		return nil, errMsg
	}
	user.DateCreated = time.Now()
	if errMsg = us.repo.CreateUser(&user); errMsg != nil {
		return nil, errMsg
	}
	userDTO.Id = user.Id
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

func (us *UserService) UpdateUser(isPartial bool, userDTO dto.UserDTO) (*dto.UserDTO, *errors.RestErr) {
	errMsg := userDTO.Validate()
	if errMsg != nil {
		return nil, errMsg
	}

	var user model.User
	if isPartial {
		user1, errMsg := us.repo.GetUser(userDTO.Id)
		if errMsg != nil {
			return nil, errMsg
		}
		user = *user1
		errMsg = userDTO.PartialUpdate(&user)
		if errMsg != nil {
			return nil, errMsg
		}
		userDTO.CopyToDTO(&user)
	} else {
		if errMsg := userDTO.CopyToUser(&user); errMsg != nil {
			return nil, errMsg
		}
	}

	if errMsg := us.repo.UpdateUser(&user); errMsg != nil {
		return nil, errMsg
	}
	return &userDTO, nil
}

func (us *UserService) DeleteUser(userId int64) *errors.RestErr {
	if errMsg := us.repo.DeleteUser(userId); errMsg != nil {
		return errMsg
	}
	return nil
}