package resolvers

import (
	"gitlab.com/lexicforlxd/backend-reloaded/graphql"
	"gitlab.com/lexicforlxd/backend-reloaded/host"
)

type Resolver struct {
	HostUsecase host.Usecase
}

func NewResolver(h host.Usecase) graphql.ResolverRoot {
	resolver := &Resolver{
		HostUsecase: h,
	}
	return resolver
}

func (r *Resolver) Container() graphql.ContainerResolver {
	return &container{}
}
func (r *Resolver) Mutation() graphql.MutationResolver {
	return newMutation(r.HostUsecase)
}
func (r *Resolver) Query() graphql.QueryResolver {
	return newQuery(r.HostUsecase)
}
