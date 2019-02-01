package graphql

import (
	"context"
	"fmt"

	"github.com/lexicforlxd/backend-reloaded/host"
	"github.com/lexicforlxd/backend-reloaded/models"
)

type HostResolver struct {
	hostUsecase host.Usecase
}

func NewHostResolver(h host.Usecase) HostResolver {
	resolver := HostResolver{
		hostUsecase: h,
	}
	return resolver
}

// mutations
func (r *HostResolver) CreateHost(ctx context.Context, hostReq models.HostReq) (*models.Host, error) {
	host := &models.Host{
		Name:    *hostReq.Name,
		Address: *hostReq.Address,
	}

	if hostReq.Desc != nil {
		host.Desc = *hostReq.Desc
	}
	if hostReq.Password != nil {
		fmt.Println("auth with lxd")
	}

	host, err := r.hostUsecase.Store(ctx, host)
	if err != nil {
		return nil, err
	}

	return host, nil
}

func (r *HostResolver) UpdateHost(ctx context.Context, id string, hostReq models.HostReq) (*models.Host, error) {
	host, err := r.hostUsecase.Show(ctx, id)

	if err != nil {
		return nil, err
	}

	if hostReq.Desc != nil {
		host.Desc = *hostReq.Desc
	}

	if hostReq.Name != nil {
		host.Name = *hostReq.Name
	}

	if hostReq.Address != nil {
		host.Address = *hostReq.Address
	}

	if host, err = r.hostUsecase.Update(ctx, host); err != nil {
		return nil, err
	}

	return host, nil

}
func (r *HostResolver) DeleteHost(ctx context.Context, id string) (*models.DeleteRes, error) {
	if err := r.hostUsecase.Delete(ctx, id); err != nil {
		return nil, err
	}

	res := models.DeleteRes{
		Message: "delete successful",
		Entity:  "Host",
	}
	return &res, nil

}
func (r *HostResolver) AuthHost(ctx context.Context, id string, authReq models.AuthHostReq) (*models.Host, error) {
	panic("not implemented")
}

// Queries
func (r *HostResolver) Hosts(ctx context.Context, limit *int, offset *int) ([]*models.Host, error) {
	limitInt := -1
	offsetInt := -1
	if limit != nil {
		limitInt = *limit
	}
	if offset != nil {
		offsetInt = *offset
	}

	hosts, err := r.hostUsecase.Fetch(ctx, limitInt, offsetInt)
	if err != nil {
		return nil, err
	}

	return hosts, nil
}
func (r *HostResolver) Host(ctx context.Context, id string) (*models.Host, error) {
	host, err := r.hostUsecase.Show(ctx, id)
	if err != nil {
		return nil, err
	}
	return host, nil
}
