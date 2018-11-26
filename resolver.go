//go:generate gorunpkg github.com/99designs/gqlgen

package backend_reloaded

import (
	context "context"

	graphql "gitlab.com/lexicforlxd/backend-reloaded/graphql"
	models "gitlab.com/lexicforlxd/backend-reloaded/models"
)

type Resolver struct{}

func (r *Resolver) Mutation() graphql.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() graphql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateHost(ctx context.Context, host models.NewHost) (models.Host, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateUser(ctx context.Context, user models.NewUser) (models.User, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Info(ctx context.Context) (*models.Info, error) {
	panic("not implemented")
}
func (r *queryResolver) Hosts(ctx context.Context) ([]*models.Host, error) {
	panic("not implemented")
}
func (r *queryResolver) Host(ctx context.Context, id string) (models.Host, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context, id string) (models.User, error) {
	panic("not implemented")
}
