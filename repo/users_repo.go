package repo

import "database/sql"

type UserRepo struct {
	conn *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		conn: db,
	}
}