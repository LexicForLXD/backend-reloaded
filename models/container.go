package models

import (
	"time"
)

type Container struct {
	ID     string
	Name   string
	Host   Host
	HostID string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
