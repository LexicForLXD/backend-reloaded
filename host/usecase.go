package host

import (
	"context"

	"github.com/lexicforlxd/backend-reloaded/models"
)

type Usecase interface {
	Fetch(ctx context.Context, limit int, offset int) ([]*models.Host, error)
	Show(ctx context.Context, ID string) (*models.Host, error)
	Update(ctx context.Context, host *models.Host) (*models.Host, error)
	Store(ctx context.Context, host *models.Host) (*models.Host, error)
	Delete(ctx context.Context, id string) error
}
