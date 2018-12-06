//go:generate gorunpkg github.com/99designs/gqlgen

package resolvers

import (
	context "context"

	"github.com/jinzhu/gorm"
	graphql "gitlab.com/lexicforlxd/backend-reloaded/graphql"
	models "gitlab.com/lexicforlxd/backend-reloaded/models"
)

type Resolver struct {
	Db *gorm.DB
}

func (r *Resolver) Container() graphql.ContainerResolver {
	return &containerResolver{r}
}
func (r *Resolver) Mutation() graphql.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() graphql.QueryResolver {
	return &queryResolver{r}
}

type containerResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, user models.UserReq) (models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, user models.UserReq) (models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (models.DeleteRes, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Info(ctx context.Context) (*models.Info, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context, id string) (models.User, error) {
	panic("not implemented")
}
