package main

import (
	"database/sql"
	"github.com/sunil206b/users_api/app"
)

func main() {
	app.StartApp(&sql.DB{})
}
