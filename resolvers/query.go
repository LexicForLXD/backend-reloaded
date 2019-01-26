package resolvers

import (
	_container "gitlab.com/lexicforlxd/backend-reloaded/container/delivery/graphql"
	"gitlab.com/lexicforlxd/backend-reloaded/host"
	_host "gitlab.com/lexicforlxd/backend-reloaded/host/delivery/graphql"
	_user "gitlab.com/lexicforlxd/backend-reloaded/user/delivery/graphql"
	_info "gitlab.com/lexicforlxd/backend-reloaded/util/delivery/graphql"
)

func newQuery(h host.Usecase) *query {
	query := &query{
		_host.NewHostResolver(h),
		_container.NewContainerResolver(h),
		_user.NewUserResolver(h),
		_info.NewInfoResolver(h),
	}

	return query
}

type query struct {
	_host.HostResolver
	_container.ContainerResolver
	_user.UserResolver
	_info.InfoResolver
	// queryResolver
}
