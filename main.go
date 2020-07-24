package main

import (
	"github.com/subosito/gotenv"
	"github.com/sunil206b/users_api/app"
	"github.com/sunil206b/users_api/driver"
)

func init() {
	gotenv.Load()
}

func main() {
	db := driver.MysqlConn()
	app.StartApp(db)
}
