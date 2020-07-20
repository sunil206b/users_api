package app

import (
	"database/sql"
	"github.com/sunil206b/users_api/controller"
)

func mapUrls(db *sql.DB) {
	userH :=controller.NewUserController(db)
	router.GET("/users/:user_id", userH.GetUser)
	router.GET("/users/search", userH.SearchUser)
	router.POST("/users", userH.CreateUser)
}
