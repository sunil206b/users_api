package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/sunil206b/users_api/logger"
)

var (
	router = gin.Default()
)

// StartApp function is used to configure all the url's
func StartApp(db *sql.DB) {
	mapUrls(db)
	logger.Info("about to start the application...")
	router.Run(":8080")
}
