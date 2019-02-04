package container

import (
	"context"

	"github.com/lexicforlxd/backend-reloaded/models"
)

type Usecase interface {
	Fetch(ctx context.Context, limit int, offset int) ([]*models.Container, error)
	FetchByHost(ctx context.Context, limit int, offset int, hostID string) ([]*models.Container, error)
	Show(ctx context.Context, ID string) (*models.Container, error)
	Update(ctx context.Context, container *models.Container) (*models.Container, error)
	Store(ctx context.Context, container *models.Container) (*models.Container, error)
	Delete(ctx context.Context, id string) error
}
