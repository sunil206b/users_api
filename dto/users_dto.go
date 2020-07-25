package dto

import (
	"github.com/sunil206b/users_api/model"
	"github.com/sunil206b/users_api/utils/date"
	"github.com/sunil206b/users_api/utils/errors"
	"strings"
	"time"
)

type UserDTO struct {
	Id          int64    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Birth       string `json:"birth"`
	Gender      string `json:"gender"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func (userDTO UserDTO) Validate() *errors.RestErr {
	userDTO.FirstName = strings.TrimSpace(userDTO.FirstName)
	userDTO.LastName = strings.TrimSpace(userDTO.LastName)
	userDTO.Birth = strings.TrimSpace(userDTO.Birth)
	userDTO.Phone = strings.TrimSpace(userDTO.Phone)
	userDTO.Gender = strings.TrimSpace(userDTO.Gender)
	userDTO.Email = strings.TrimSpace(strings.ToLower(userDTO.Email))
	userDTO.Password = strings.TrimSpace(userDTO.Password)
	if userDTO.Email == "" {
		return errors.NewBadRequest("invalid email address")
	}
	return nil
}

func (userDTO UserDTO) CopyToUser(user *model.User) *errors.RestErr{
	user.Id = userDTO.Id
	user.FirstName = userDTO.FirstName
	user.LastName = userDTO.LastName
	if errMsg := userDTO.SetBirthToUser(user); errMsg != nil {
		return errMsg
	}
	user.Gender = userDTO.Gender
	user.Phone = userDTO.Phone
	user.Email = userDTO.Email
	user.Password = userDTO.Password
	user.DateUpdated = time.Now()
	return nil
}

func (userDTO *UserDTO) CopyToDTO(user *model.User) {
	userDTO.Id = user.Id
	userDTO.FirstName = user.FirstName
	userDTO.LastName = user.LastName
	userDTO.Birth = date.GetFmtDate(user.Birth)
	userDTO.Gender = user.Gender
	userDTO.Phone = user.Phone
	userDTO.Email = user.Email
	userDTO.Password = user.Password
}

func (userDTO *UserDTO) SetBirthToUser(user *model.User) *errors.RestErr {
	age, err := date.ConvertToDate(userDTO.Birth)
	if err != nil {
		return errors.NewBadRequest(err.Error())
	}
	user.Birth = age
	return nil
}

func (userDTO *UserDTO) PartialUpdate(user *model.User) *errors.RestErr {
	if userDTO.FirstName != "" {
		user.FirstName = userDTO.FirstName
	}
	if userDTO.LastName != "" {
		user.LastName = userDTO.LastName
	}
	if userDTO.Birth != "" {
		if errMsg := userDTO.SetBirthToUser(user); errMsg != nil {
			return errMsg
		}
	}
	if userDTO.Gender != "" {
		user.Gender = userDTO.Gender
	}
	if userDTO.Phone != "" {
		user.Phone = userDTO.Phone
	}
	if userDTO.Email != "" {
		user.Email = userDTO.Email
	}
	if userDTO.Password != "" {
		user.Password = userDTO.Password
	}
	user.DateUpdated = time.Now()
	return nil
}