package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/lexicforlxd/backend-reloaded/container"
	"github.com/lexicforlxd/backend-reloaded/models"
)

type containerRepository struct {
	db *gorm.DB
}

func NewContainerRepository(db *gorm.DB) container.Repository {
	return &containerRepository{db}
}

func (cr *containerRepository) Fetch(ctx context.Context, limit int, offset int) ([]*models.Container, error) {
	var containers []*models.Container

	if err := cr.db.Limit(limit).Offset(offset).Order("name").Find(&containers).Error; err != nil {
		return nil, err
	}

	return containers, nil
}

// Set all the fields in the Container struct you want to search for
func (cr *containerRepository) Get(ctx context.Context, searchedContainer models.Container) (*models.Container, error) {
	container := models.Container{}
	if err := cr.db.Or(searchedContainer).First(&container).Error; err != nil {
		return nil, err
	}
	return &container, nil
}

func (cr *containerRepository) GetByID(ctx context.Context, ID string) (*models.Container, error) {
	container := models.Container{}
	if err := cr.db.Where("id = ?", ID).First(&container).Error; err != nil {
		return nil, err
	}
	return &container, nil
}
