package repo

import (
	"database/sql"
	"fmt"
	"github.com/sunil206b/users_api/logger"
	"github.com/sunil206b/users_api/model"
	"github.com/sunil206b/users_api/utils/errors"
	"strings"
	"time"
)

const (
	indexUnique = "unique"
	userInsert = "INSERT INTO users(first_name, last_name, birth, gender, phone, email, password, created_at, updated_at, status) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	queryGetUser = "SELECT id, first_name, last_name, birth, gender, phone, email, password, status  FROM users WHERE id=?"
	queryUpdate = "UPDATE users set first_name=?, last_name=?, birth=?, gender=?, phone=?, email=?, password=?,  updated_at=?, status=? WHERE id=?"
	queryDelete = "UPDATE users set status=?, updated_at=? WHERE id=?"
	queryByStatus = "SELECT id, first_name, last_name, birth, gender, phone, email, password, status  FROM users WHERE status=?"
	queryFindByEmail = "SELECT id, email, password FROM users WHERE email=? AND status=?"
	status = "active"
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
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("data base error when saving user")
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.FirstName, user.LastName, user.Birth, user.Gender, user.Phone, user.Email, user.Password, user.DateCreated, user.DateUpdated, user.Status)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), indexUnique) {
			logger.Error("unique constraint error when trying to save user", err)
			return errors.NewBadRequest(fmt.Sprintf("email %s already exists", user.Email))
		}
		logger.Error("error when trying to save user", err)
		return errors.NewInternalServerError("error when trying to save user")
	}
	 userId, err := result.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get user id after save", err)
		return errors.NewInternalServerError("error when trying to get user id after save")
	}
	user.Id = userId
	return nil
}

func (u *userRepo)  GetUser(userId int64) (*model.User, *errors.RestErr) {
	stmt, err := u.conn.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return nil, errors.NewInternalServerError("data base error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(userId)
	var user model.User
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Birth, &user.Gender, &user.Phone, &user.Email, &user.Password, &user.Status); err != nil {
		if err == sql.ErrNoRows {
			logger.Error("error when trying to get user by id ", err)
			return nil, errors.NewNotFoundError(fmt.Sprintf("user %d not found", userId))
		}
		logger.Error("error when trying to get user by id ", err)
		return nil, errors.NewInternalServerError(fmt.Sprintf("error when trying to get user with %d", userId))
	}
	return &user, nil
}

func (u *userRepo) UpdateUser(user *model.User) *errors.RestErr {
	stmt, err := u.conn.Prepare(queryUpdate)
	if err != nil {
		logger.Error("error when trying prepare update user query ", err)
		return errors.NewInternalServerError("data base error when updating user")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Birth, user.Gender, user.Phone, user.Email, user.Password, user.DateUpdated, user.Status, user.Id)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), indexUnique) {
			logger.Error("error when trying to update user", err)
			return errors.NewBadRequest(fmt.Sprintf("email %s already exists", user.Email))
		}
		logger.Error("error when trying to update user", err)
		return errors.NewInternalServerError("error when trying to update user")
	}
	return nil
}

func (u *userRepo) DeleteUser(userId int64)  *errors.RestErr {
	stmt, err := u.conn.Prepare(queryDelete)
	if err != nil {
		logger.Error("error when preparing delete user query", err)
		return errors.NewInternalServerError("data base error when deleting user")
	}
	defer stmt.Close()

	if _, err = stmt.Exec("inactive", time.Now(), userId); err != nil {
		logger.Error("error when trying to delete user", err)
		return errors.NewInternalServerError("error when trying to delete user")
	}
	return nil
}

func (u *userRepo) Search(status string) ([]model.User, *errors.RestErr) {
	stmt, err := u.conn.Prepare(queryByStatus)
	if err != nil {
		logger.Error("error when preparing search users query", err)
		return nil, errors.NewInternalServerError("data base when searching for users")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying search users", err)
		return nil, errors.NewInternalServerError("error when searching users")
	}
	defer rows.Close()

	users := make([]model.User, 0)
	for rows.Next() {
		user := model.User{}
		if err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Birth, &user.Gender, &user.Phone, &user.Email, &user.Password, &user.Status); err != nil {
			logger.Error("error when scanning users from row", err)
			return nil, errors.NewNotFoundError("error when searching for users")
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		logger.Info(fmt.Sprintf("no users found in the data base with given status %s", status))
		return nil, errors.NewNotFoundError(fmt.Sprintf("no user found with status %s", status))
	}
	return users, nil
}

func (u *userRepo) FindByEmail(email string) (*model.LoginUser, *errors.RestErr) {
	stmt, err := u.conn.Prepare(queryFindByEmail)
	if err != nil {
		logger.Error("error when trying to prepare find user by email statement", err)
		return nil, errors.NewInternalServerError("data base error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(email, status)
	var user model.LoginUser
	if err := result.Scan(&user.Id, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			logger.Error("error when trying to find user by email ", err)
			return nil, errors.NewNotFoundError(fmt.Sprintf("no active user found with given email %s", email))
		}
		logger.Error("error when trying to find user by emil ", err)
		return nil, errors.NewInternalServerError(fmt.Sprintf("error when trying to find user with %s", email))
	}
	return &user, nil
}