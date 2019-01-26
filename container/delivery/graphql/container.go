package graphql

import (
	"context"
	"time"

	"github.com/lxc/lxd/shared/api"
	"gitlab.com/lexicforlxd/backend-reloaded/host"
	"gitlab.com/lexicforlxd/backend-reloaded/models"
)

type ContainerResolver struct {
}

type ContainerFieldResolver struct{}

func NewContainerResolver(h host.Usecase) ContainerResolver {
	resolver := ContainerResolver{
		// hostUsecase: h,
	}
	return resolver
}

// Mutations
func (r *ContainerResolver) CreateContainer(ctx context.Context, container models.ContainerReq) (*api.Container, error) {
	// c, err := lxd.ConnectLXDUnix("", nil)
	// if err != nil {
	// 	return err
	// }
	panic("not implemented")

}
func (r *ContainerResolver) UpdateContainer(ctx context.Context, id string, container models.ContainerReq) (*api.Container, error) {
	panic("not implemented")
}
func (r *ContainerResolver) DeleteContainer(ctx context.Context, id string) (*models.DeleteRes, error) {
	panic("not implemented")
}

// Queries
func (r *ContainerResolver) Containers(ctx context.Context, hostID *string, limit *int, offset *int) ([]*api.Container, error) {
	panic("not implemented")
}
func (r *ContainerResolver) Container(ctx context.Context, id string) (*api.Container, error) {
	panic("not implemented")
}

// Fields
func (r *ContainerFieldResolver) ID(ctx context.Context, obj *api.Container) (string, error) {
	panic("not implemented")
}
func (r *ContainerFieldResolver) Host(ctx context.Context, obj *api.Container) (*models.Host, error) {
	panic("not implemented")
}
func (r *ContainerFieldResolver) HostID(ctx context.Context, obj *api.Container) (string, error) {
	panic("not implemented")
}
func (r *ContainerFieldResolver) Source(ctx context.Context, obj *api.Container) (*models.ContianerSource, error) {
	panic("not implemented")
}
func (r *ContainerFieldResolver) UpdatedAt(ctx context.Context, obj *api.Container) (*time.Time, error) {
	panic("not implemented")
}
func (r *ContainerFieldResolver) DeletedAt(ctx context.Context, obj *api.Container) (*time.Time, error) {
	panic("not implemented")
}
