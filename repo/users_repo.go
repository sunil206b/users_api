package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/sunil206b/users_api/model"
)

var (
	userDB = make(map[int]*model.User)
)
type userRepo struct {
	conn *sql.DB
}

func NewUserRepo(db *sql.DB) IUserRepo {
	return &userRepo{
		conn: db,
	}
}

func (u *userRepo) CreateUser(user *model.User) error {
	current := userDB[user.Id]
	if current != nil {
		return errors.New(fmt.Sprintf("user already exist with id %d", user.Id))
	}
	userDB[user.Id] = user
	return nil
}

func (u *userRepo)  GetUser(userId int) (*model.User, error) {
	result := userDB[userId]
	if result == nil {
		return nil, errors.New(fmt.Sprintf("user not found with id %d", userId))
	}
	var user model.User
	user.Id = result.Id
	user.Email = result.Email
	user.LastName = result.LastName
	user.FirstName = result.FirstName
	user.DateCreated = result.DateCreated
	user.DateUpdated = result.DateUpdated
	return &user, nil
}

func (u *userRepo) UpdateUser(user *model.User) error {
	return nil
}

func (u *userRepo) DeleteUser(userId int) (bool, error) {
	return false, nil
}