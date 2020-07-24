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
	user.Gender = userDTO.Gender
	user.Phone = userDTO.Phone
	user.Email = userDTO.Email
	user.Password = userDTO.Password
	user.DateUpdated = time.Now()
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
	userDTO.DateCreated = date.GetFmtDate(user.DateCreated)
	userDTO.DateUpdated = date.GetFmtDate(user.DateUpdated)
}
