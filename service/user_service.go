package service

import (
	"database/sql"
	"github.com/sunil206b/users_api/dto"
	"github.com/sunil206b/users_api/model"
	"github.com/sunil206b/users_api/repo"
	"github.com/sunil206b/users_api/utils/errors"
	"time"
)

type userService struct {
	repo repo.IUserRepo
}

func NewUserService(db *sql.DB) IUserService {
	return &userService{
		repo: repo.NewUserRepo(db),
	}
}

func (us *userService) CreateUse(userDTO dto.UserDTO) (*dto.UserDTO, *errors.RestErr) {
	var user model.User
	if errMsg := userDTO.CopyToUser(&user); errMsg != nil {
		return nil, errMsg
	}
	user.DateCreated = time.Now().UTC()
	if errMsg := us.repo.CreateUser(&user); errMsg != nil {
		return nil, errMsg
	}
	userDTO.Id = user.Id
	userDTO.Password = user.Password
	userDTO.Status = user.Status
	return &userDTO, nil
}

func (us *userService) GetUser(userId int64) (*dto.UserDTO, *errors.RestErr) {
	user, err := us.repo.GetUser(userId)
	if err != nil {
		return nil, err
	}
	var userDto dto.UserDTO
	userDto.CopyToDTO(user)
	return &userDto, nil
}

func (us *userService) UpdateUser(isPartial bool, userDTO dto.UserDTO) (*dto.UserDTO, *errors.RestErr) {
	user, errMsg := us.repo.GetUser(userDTO.Id)
	if errMsg != nil {
		return nil, errMsg
	}
	if isPartial {
		errMsg = userDTO.PartialUpdate(user)
		if errMsg != nil {
			return nil, errMsg
		}
		userDTO.CopyToDTO(user)
	} else {
		if errMsg := userDTO.CopyToUser(user); errMsg != nil {
			return nil, errMsg
		}
		userDTO.Status = user.Status
		userDTO.Password = user.Password
	}

	if errMsg := us.repo.UpdateUser(user); errMsg != nil {
		return nil, errMsg
	}
	return &userDTO, nil
}

func (us *userService) DeleteUser(userId int64) *errors.RestErr {
	if errMsg := us.repo.DeleteUser(userId); errMsg != nil {
		return errMsg
	}
	return nil
}

func (us *userService) Search(status string) ([]dto.UserDTO, *errors.RestErr) {
	users, errMsg := us.repo.Search(status)
	if errMsg != nil {
		return nil, errMsg
	}
	userDTOs := make([]dto.UserDTO, 0)
	for _, user := range users {
		userDTO := dto.UserDTO{}
		userDTO.CopyToDTO(&user)
		userDTOs = append(userDTOs, userDTO)
	}
	return userDTOs, nil
}

func (us *userService) FindUserByEmail(email string) (*dto.LoginUserDTO, *errors.RestErr) {
	user, err := us.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	var userDto dto.LoginUserDTO
	userDto.CopyToLoginDTO(user)
	return &userDto, nil
}
