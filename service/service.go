package service

import (
	"github.com/sunil206b/store_utils_go/errors"
	"github.com/sunil206b/users_api/dto"
)

type IUserService interface {
	CreateUse(userDTO dto.UserDTO) (*dto.UserDTO, *errors.RestErr)
	GetUser(userId int64) (*dto.UserDTO, *errors.RestErr)
	UpdateUser(isPartial bool, userDTO dto.UserDTO) (*dto.UserDTO, *errors.RestErr)
	DeleteUser(userId int64) *errors.RestErr
	Search(status string) ([]dto.UserDTO, *errors.RestErr)
	FindUserByEmail(email string) (*dto.LoginUserDTO, *errors.RestErr)
}
