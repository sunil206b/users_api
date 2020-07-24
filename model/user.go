package model

import "time"

type User struct {
	Id          int64       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Birth       time.Time `json:"birth"`
	Gender      string    `json:"gender"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}
