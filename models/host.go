package models

import (
	"time"
)

type Host struct {
	ID            string      `json:"ID"`
	Name          string      `json:"name"`
	Desc          string      `json:"desc,omitempty"`
	Address       string      `json:"address"`
	Password      string      `json:"password,omitempty"`
	Authenticated bool        `json:"authenticated"`
	Containers    []Container `json:"containers"`
	CreatedAt     time.Time   `json:"createdAt"`
	UpdatedAt     time.Time   `json:"updatedAt"`
	DeletedAt     *time.Time  `json:"deletedAt"`
}
