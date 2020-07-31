package service

import (
	"github.com/sunil206b/users_api/dto"
	"github.com/sunil206b/users_api/utils/errors"
)

type IUserService interface {
	CreateUse(userDTO dto.UserDTO) (*dto.UserDTO, *errors.RestErr)
	GetUser(userId int64) (*dto.UserDTO, *errors.RestErr)
	UpdateUser(isPartial bool, userDTO dto.UserDTO) (*dto.UserDTO, *errors.RestErr)
	DeleteUser(userId int64) *errors.RestErr
	Search(status string) ([]dto.UserDTO, *errors.RestErr)
}
