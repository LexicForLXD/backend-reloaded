package container

import (
	"context"

	"github.com/lexicforlxd/backend-reloaded/models"
)

type Repository interface {
	Fetch(ctx context.Context, limit int, offset int) ([]*models.Container, error)
	FetchByHost(ctx context.Context, limit int, offset int, hostID string) ([]*models.Container, error)
	GetByID(ctx context.Context, ID string) (*models.Container, error)
	GetByName(ctx context.Context, name string) (*models.Container, error)
	Get(ctx context.Context, field string, value string) (*models.Container, error)
	Update(ctx context.Context, container *models.Container) error
	Store(ctx context.Context, container *models.Container) error
	Delete(ctx context.Context, id string) error
}
