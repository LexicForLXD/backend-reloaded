package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/lexicforlxd/backend-reloaded/host"
	"github.com/lexicforlxd/backend-reloaded/models"
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

func (hr *hostRepository) GetByAddress(ctx context.Context, address string) (*models.Host, error) {
	host := models.Host{}
	if err := hr.db.Where("address = ?", address).First(&host).Error; err != nil {
		return nil, err
	}
	return &host, nil
}

func (hr *hostRepository) GetByName(ctx context.Context, name string) (*models.Host, error) {
	host := models.Host{}
	if err := hr.db.Where("name = ?", name).First(&host).Error; err != nil {
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
