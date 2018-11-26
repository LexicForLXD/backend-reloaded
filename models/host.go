package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Host struct {
	ID         string
	Name       string
	Address    string
	Containers []Container
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

type HostRepositoryInt interface {
	GetAllHosts() ([]*Host, error)
	GetHost(ID string) (*Host, error)
	CreateHost(Host *Host) error
	UpdateHost(ID string, Host *Host) error
	DeleteHost(ID string) error
}

type HostRepository struct {
	Db *gorm.DB
}
