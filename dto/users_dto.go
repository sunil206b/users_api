package dto

import (
	"github.com/sunil206b/users_api/model"
	"github.com/sunil206b/users_api/utils/errors"
	"strings"
)

type UserDTO struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}

func (userDTO UserDTO) Validate() *errors.RestErr {
	userDTO.Email = strings.TrimSpace(strings.ToLower(userDTO.Email))
	if userDTO.Email == "" {
		return errors.NewBadRequest("invalid email address")
	}
	return nil
}

func (userDTO UserDTO) CopyToUser(user *model.User) {
	user.Id = userDTO.Id
	user.FirstName = userDTO.FirstName
	user.LastName = userDTO.LastName
	user.Email = userDTO.Email
	user.Password = userDTO.Password
	user.DateCreated = userDTO.DateCreated
	user.DateUpdated = userDTO.DateUpdated
}