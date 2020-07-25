package app

import (
	"database/sql"
	"github.com/sunil206b/users_api/controller"
)

func mapUrls(db *sql.DB) {
	userH :=controller.NewUserController(db)
	router.POST("/users", userH.CreateUser)
	router.GET("/users/:user_id", userH.GetUser)
	router.PUT("/users/:user_id", userH.UpdateUser)
	router.PATCH("/users/:user_id", userH.UpdateUser)
	router.DELETE("/users/:user_id", userH.DeleteUser)
	//router.GET("/users/search", userH.SearchUser)
}
