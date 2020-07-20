package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)


func NewUserController(db *sql.DB) *UserController{
	return &UserController{
		conn: db,
	}
}

type UserController struct {
	conn *sql.DB
}

func (u *UserController) CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func (u *UserController) GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func (u *UserController) SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}