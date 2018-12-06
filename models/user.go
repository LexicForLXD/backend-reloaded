package models

import (
	"time"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Birthday  time.Time
	Email     string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
