package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/sunil206b/users_api/model"
	"github.com/sunil206b/users_api/service"
	"net/http"
)


func NewUserController(db *sql.DB) *UserController{
	return &UserController{
		us: service.NewUserService(db),
	}
}

type UserController struct {
	us *service.UserService
}

func (u *UserController) CreateUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		// TODO : Handle Error
		return
	}

	result, err := u.us.CreateUser(user)
	if err != nil {
		// TODO: handle user creation error
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (u *UserController) GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func (u *UserController) SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}