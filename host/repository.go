package host

import (
	"context"

	"github.com/lexicforlxd/backend-reloaded/models"
)

type Repository interface {
	Fetch(ctx context.Context, limit int, offset int) ([]*models.Host, error)
	GetByID(ctx context.Context, ID string) (*models.Host, error)
	Update(ctx context.Context, host *models.Host) error
	Store(ctx context.Context, host *models.Host) error
	Delete(ctx context.Context, id string) error
}
