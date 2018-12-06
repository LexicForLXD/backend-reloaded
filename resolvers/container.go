package resolvers

import (
	context "context"
	time "time"

	api "github.com/lxc/lxd/shared/api"
	models "gitlab.com/lexicforlxd/backend-reloaded/models"
)

// Mutations
func (r *mutationResolver) CreateContainer(ctx context.Context, container models.ContainerReq) (api.Container, error) {
	// c, err := lxd.ConnectLXDUnix("", nil)
	// if err != nil {
	// 	return err
	// }
	panic("not implemented")

}
func (r *mutationResolver) UpdateContainer(ctx context.Context, id string, container models.ContainerReq) (api.Container, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteContainer(ctx context.Context, id string) (models.DeleteRes, error) {
	panic("not implemented")
}

// Queries
func (r *queryResolver) Containers(ctx context.Context, hostID *string) ([]*api.Container, error) {
	panic("not implemented")
}
func (r *queryResolver) Container(ctx context.Context, id string) (api.Container, error) {
	panic("not implemented")
}

// Fields
func (r *containerResolver) ID(ctx context.Context, obj *api.Container) (string, error) {
	panic("not implemented")
}
func (r *containerResolver) Host(ctx context.Context, obj *api.Container) (*models.Host, error) {
	panic("not implemented")
}
func (r *containerResolver) HostID(ctx context.Context, obj *api.Container) (string, error) {
	panic("not implemented")
}
func (r *containerResolver) Source(ctx context.Context, obj *api.Container) (*models.ContianerSource, error) {
	panic("not implemented")
}
func (r *containerResolver) UpdatedAt(ctx context.Context, obj *api.Container) (*time.Time, error) {
	panic("not implemented")
}
func (r *containerResolver) DeletedAt(ctx context.Context, obj *api.Container) (*time.Time, error) {
	panic("not implemented")
}
