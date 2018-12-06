package resolvers

import (
	context "context"
	"fmt"

	"github.com/google/uuid"
	"gitlab.com/lexicforlxd/backend-reloaded/models"
)

// Mutations
func (r *mutationResolver) CreateHost(ctx context.Context, hostReq models.HostReq) (*models.Host, error) {
	// panic("not implemented")
	host := models.Host{
		ID:      uuid.New().String(),
		Name:    hostReq.Name,
		Address: hostReq.Address,
	}

	if hostReq.Desc != nil {
		host.Desc = *hostReq.Desc
	}
	if hostReq.Password != nil {
		fmt.Println("auth with lxd")
	}

	if err := r.Db.Create(&host).Error; err != nil {
		return nil, err
	}

	return &host, nil
}
func (r *mutationResolver) UpdateHost(ctx context.Context, id string, host models.HostReq) (models.Host, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteHost(ctx context.Context, id string) (models.DeleteRes, error) {
	panic("not implemented")
}
func (r *mutationResolver) AuthHost(ctx context.Context, id string, authReq models.AuthHostReq) (models.Host, error) {
	panic("not implemented")
}

// Queries
func (r *queryResolver) Hosts(ctx context.Context, limit *int) ([]*models.Host, error) {
	var hosts []*models.Host

	if limit != nil {
		if err := r.Db.Limit(*limit).Find(&hosts).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.Db.Find(&hosts).Error; err != nil {
			return nil, err
		}
	}

	return hosts, nil
}
func (r *queryResolver) Host(ctx context.Context, id string) (*models.Host, error) {
	host := models.Host{}
	if err := r.Db.Where("id = ?", id).First(&host).Error; err != nil {
		return nil, err
	}

	return &host, nil
}
