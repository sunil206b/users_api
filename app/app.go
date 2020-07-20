package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)
// StartApp function is used to configure all the url's
func StartApp(db *sql.DB) {
	mapUrls(db)
	router.Run("8080")
}
