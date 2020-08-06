package dto

import (
	"github.com/sunil206b/store_utils_go/date"
	"github.com/sunil206b/store_utils_go/errors"
	"github.com/sunil206b/users_api/model"
	"github.com/sunil206b/store_utils_go/crypto"
	"strings"
	"time"
)

type UserDTO struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Birth     string `json:"birth"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    string `json:"status"`
}

type LoginUserDTO struct {
	Id        int64  `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (loginUser *LoginUserDTO) CopyToLoginDTO(user *model.LoginUser) {
	loginUser.Id = user.Id
	loginUser.Email = user.Email
	loginUser.Password = user.Password
}

func (userDTO UserDTO) Validate() []*errors.RestErr {
	userDTO.FirstName = strings.TrimSpace(userDTO.FirstName)
	userDTO.LastName = strings.TrimSpace(userDTO.LastName)
	userDTO.Birth = strings.TrimSpace(userDTO.Birth)
	userDTO.Phone = strings.TrimSpace(userDTO.Phone)
	userDTO.Gender = strings.TrimSpace(userDTO.Gender)
	userDTO.Email = strings.TrimSpace(strings.ToLower(userDTO.Email))
	userDTO.Password = strings.TrimSpace(userDTO.Password)
	errMsg := make([]*errors.RestErr, 0)
	if userDTO.FirstName == "" {
		errMsg = append(errMsg, errors.NewBadRequest("First Name is required"))
	}
	if userDTO.LastName == "" {
		errMsg = append(errMsg, errors.NewBadRequest("Last Name is required"))
	}
	if userDTO.Birth == "" {
		errMsg = append(errMsg, errors.NewBadRequest("Date of Birth is required"))
	}
	if userDTO.Phone == "" {
		errMsg = append(errMsg, errors.NewBadRequest("Phone number is required"))
	}
	if userDTO.Gender == "" {
		errMsg = append(errMsg, errors.NewBadRequest("Gender is required"))
	}
	if userDTO.Email == "" {
		errMsg = append(errMsg, errors.NewBadRequest("invalid email address"))
	}
	if userDTO.Password == "" {
		errMsg = append(errMsg, errors.NewBadRequest("invalid password"))
	}
	return errMsg
}

func (userDTO UserDTO) CopyToUser(user *model.User) *errors.RestErr {
	user.Id = userDTO.Id
	user.FirstName = userDTO.FirstName
	user.LastName = userDTO.LastName
	if errMsg := userDTO.SetBirthToUser(user); errMsg != nil {
		return errMsg
	}
	user.Gender = userDTO.Gender
	user.Phone = userDTO.Phone
	user.Email = userDTO.Email
	if password := crypto.GetMd5(userDTO.Password); user.Password != password {
		user.Password = password
	}
	user.DateUpdated = time.Now().UTC()
	user.Status = "active"
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
	userDTO.Status = user.Status
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
		if password := crypto.GetMd5(userDTO.Password); user.Password != password {
			user.Password = password
		}
	}
	user.DateUpdated = time.Now()
	return nil
}
