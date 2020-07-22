package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/sunil206b/users_api/dto"
	"github.com/sunil206b/users_api/service"
	"github.com/sunil206b/users_api/utils/errors"
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
	var user dto.UserDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		errMsg := errors.NewBadRequest("Invalid json body")
		c.JSON(errMsg.StatusCode, errMsg)
		return
	}
	result, err := u.us.CreateUse(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
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