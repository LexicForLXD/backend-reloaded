package models

import (
	"time"
)

type Host struct {
	ID            string
	Name          string
	Desc          string
	Address       string
	Authenticated bool
	Containers    []Container
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}
