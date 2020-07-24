package repo

import (
	"database/sql"
	"fmt"
	"github.com/sunil206b/users_api/model"
	"github.com/sunil206b/users_api/utils/errors"
	"strings"
)

const (
	indexUnique = "unique"
	userInsert = "INSERT INTO users(first_name, last_name, birth, gender, phone, email, password, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)"
	queryGetUser = "SELECT id, first_name, last_name, birth, gender, phone, email, password,  created_at,  updated_at  FROM users WHERE id=?"
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

func (u *userRepo) CreateUser(user *model.User) *errors.RestErr {
	stmt, err := u.conn.Prepare(userInsert)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s", err.Error()))
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.FirstName, user.LastName, user.Birth, user.Gender, user.Phone, user.Email, user.Password, user.DateCreated, user.DateUpdated)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), indexUnique) {
			return errors.NewBadRequest(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s", err.Error()))
	}
	 userId, err := result.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s", err.Error()))
	}
	user.Id = userId
	return nil
}

func (u *userRepo)  GetUser(userId int64) (*model.User, *errors.RestErr) {
	stmt, err := u.conn.Prepare(queryGetUser)
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s", err.Error()))
	}
	defer stmt.Close()

	result := stmt.QueryRow(userId)
	var user model.User
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Birth, &user.Gender, &user.Phone, &user.Email, &user.Password, &user.DateCreated, &user.DateUpdated); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError(fmt.Sprintf("user %d not found", userId))
		}
		return nil, errors.NewInternalServerError(fmt.Sprintf("error when trying to get user with %d: %s", userId, err.Error()))
	}
	return &user, nil
}

func (u *userRepo) UpdateUser(user *model.User) *errors.RestErr {
	return nil
}

func (u *userRepo) DeleteUser(userId int64) (bool, *errors.RestErr) {
	return false, nil
}