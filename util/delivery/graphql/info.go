package graphql

import (
	"context"

	"gitlab.com/lexicforlxd/backend-reloaded/host"
	"gitlab.com/lexicforlxd/backend-reloaded/models"
)

type InfoResolver struct{}

func NewInfoResolver(h host.Usecase) InfoResolver {
	resolver := InfoResolver{
		// hostUsecase: h,
	}
	return resolver
}

func (r *InfoResolver) Info(ctx context.Context) (*models.Info, error) {
	panic("not implemented")
}
func (r *InfoResolver) Infos(ctx context.Context) (*models.Info, error) {
	panic("not implemented")
}
