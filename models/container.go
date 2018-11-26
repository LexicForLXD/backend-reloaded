package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Container struct {
	ID        string
	Name      string
	Host      Host
	HostID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type ContainerRepositoryInt interface {
	GetAllContainers() ([]*Container, error)
	GetContainer(ID string) (*Container, error)
	CreateContainer(Container *Container) error
	UpdateContainer(ID string, Container *Container) error
	DeleteContainer(ID string) error
}

type ContainerRepository struct {
	Db *gorm.DB
}
