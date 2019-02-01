package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lexicforlxd/backend-reloaded/host"
	"github.com/lexicforlxd/backend-reloaded/models"
)

type hostUsecase struct {
	hostRepo       host.Repository
	contextTimeout time.Duration
}

// NewHostUsecase will retrun a new businesslogic
func NewHostUsecase(h host.Repository, timeout time.Duration) host.Usecase {
	return &hostUsecase{
		hostRepo:       h,
		contextTimeout: timeout,
	}
}

func (h *hostUsecase) Store(ctx context.Context, host *models.Host) (*models.Host, error) {

	if host, _ := h.hostRepo.GetByAddress(ctx, host.Address); host != nil {
		return nil, errors.New("Host with same address already in database")
	}
	if host, _ := h.hostRepo.GetByName(ctx, host.Name); host != nil {
		return nil, errors.New("Host with same name already in database")
	}

	if host.Password != "" {
		fmt.Println("auth with lxd")
	}

	host.ID = uuid.New().String()
	if err := h.hostRepo.Store(ctx, host); err != nil {
		return nil, err
	}
	host, err := h.hostRepo.GetByID(ctx, host.ID)
	if err != nil {
		return nil, err
	}

	return host, nil
}

func (h *hostUsecase) Delete(ctx context.Context, ID string) error {
	if _, err := h.hostRepo.GetByID(ctx, ID); err != nil {
		return err
	}
	if err := h.hostRepo.Delete(ctx, ID); err != nil {
		return err
	}

	return nil
}

func (h *hostUsecase) Fetch(ctx context.Context, limit int, offset int) ([]*models.Host, error) {
	if limit == 0 {
		limit = -1
	}
	if offset == 0 {
		offset = -1
	}

	hosts, err := h.hostRepo.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return hosts, nil
}

func (h *hostUsecase) Update(ctx context.Context, host *models.Host) (*models.Host, error) {
	if _, err := h.hostRepo.GetByID(ctx, host.ID); err != nil {
		return nil, err
	}
	if err := h.hostRepo.Update(ctx, host); err != nil {
		return nil, err
	}
	host, err := h.hostRepo.GetByID(ctx, host.ID)
	if err != nil {
		return nil, err
	}

	return host, nil
}

func (h *hostUsecase) Show(ctx context.Context, ID string) (*models.Host, error) {
	host, err := h.hostRepo.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return host, nil
}
