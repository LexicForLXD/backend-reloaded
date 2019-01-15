package resolvers

import (
	"context"

	"github.com/jinzhu/gorm"
	"gitlab.com/lexicforlxd/backend-reloaded/graphql"
	"gitlab.com/lexicforlxd/backend-reloaded/models"
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

type queryResolver struct{ *Resolver }

func (r *queryResolver) Info(ctx context.Context) (*models.Info, error) {
	panic("not implemented")
}
