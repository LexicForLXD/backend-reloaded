package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"gitlab.com/lexicforlxd/backend-reloaded/host"
	"gitlab.com/lexicforlxd/backend-reloaded/models"
)

type hostRepository struct {
	db *gorm.DB
}

func NewHostRepository(db *gorm.DB) host.Repository {
	return &hostRepository{db}
}

func (hr *hostRepository) Fetch(ctx context.Context, limit int, offset int) ([]*models.Host, error) {
	var hosts []*models.Host

	if err := hr.db.Limit(limit).Offset(offset).Order("name").Find(&hosts).Error; err != nil {
		return nil, err
	}

	return hosts, nil
}

func (hr *hostRepository) GetByID(ctx context.Context, ID string) (*models.Host, error) {
	host := models.Host{}
	if err := hr.db.Where("id = ?", ID).First(&host).Error; err != nil {
		return nil, err
	}

	return &host, nil
}

func (hr *hostRepository) Update(ctx context.Context, host *models.Host) error {
	if err := hr.db.Save(&host).Error; err != nil {
		return err
	}

	return nil
}

func (hr *hostRepository) Store(ctx context.Context, host *models.Host) error {
	if err := hr.db.Create(&host).Error; err != nil {
		return err
	}

	return nil
}

func (hr *hostRepository) Delete(ctx context.Context, ID string) error {
	host := models.Host{
		ID: ID,
	}

	if err := hr.db.Delete(&host).Error; err != nil {
		return err
	}

	return nil
}
