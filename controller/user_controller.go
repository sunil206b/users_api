package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/sunil206b/users_api/dto"
	"github.com/sunil206b/users_api/service"
	"github.com/sunil206b/users_api/utils/errors"
	"net/http"
	"strconv"
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
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (u *UserController) GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil || userId <= 0{
		errMsg := errors.NewBadRequest("invalid user id")
		c.JSON(errMsg.StatusCode, errMsg)
		return
	}
	userDto, errMsg := u.us.GetUser(userId)
	if errMsg != nil {
		c.JSON(errMsg.StatusCode, errMsg)
		return
	}
	c.JSON(http.StatusOK, userDto)
}

func (u *UserController) SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}