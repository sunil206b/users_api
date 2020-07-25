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

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil || userId <= 0 {
		return 0, errors.NewBadRequest("invalid user id")
	}
	return userId, nil
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
	userId, errMsg := getUserId(c.Param("user_id"))
	if errMsg != nil {
		c.JSON(errMsg.StatusCode, errMsg)
		return
	}
	userDTO, errMsg := u.us.GetUser(userId)
	if errMsg != nil {
		c.JSON(errMsg.StatusCode, errMsg)
		return
	}
	c.JSON(http.StatusOK, userDTO)
}

func (u *UserController) SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func (u *UserController) UpdateUser(c *gin.Context) {
	userId, errMsg := getUserId(c.Param("user_id"))
	if errMsg != nil {
		c.JSON(errMsg.StatusCode, errMsg)
		return
	}

	var userDto dto.UserDTO
	if err := c.ShouldBindJSON(&userDto); err != nil {
		errMsg := errors.NewBadRequest("Invalid json body")
		c.JSON(errMsg.StatusCode, errMsg)
		return
	}

	userDto.Id = userId
	isPartial := c.Request.Method == http.MethodPatch
	result, errMSg := u.us.UpdateUser(isPartial, userDto)
	if errMSg != nil {
		c.JSON(errMSg.StatusCode, errMSg)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (u *UserController) DeleteUser(c *gin.Context) {
	userId, errMsg := getUserId(c.Param("user_id"))
	if errMsg != nil {
		c.JSON(errMsg.StatusCode, errMsg)
		return
	}

	if errMsg = u.us.DeleteUser(userId); errMsg != nil {
		c.JSON(errMsg.StatusCode, errMsg)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}