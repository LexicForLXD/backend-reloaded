package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/lexicforlxd/backend-reloaded/host/mocks"
	ucase "github.com/lexicforlxd/backend-reloaded/host/usecase"
	"github.com/lexicforlxd/backend-reloaded/models"
)

func TestGetByID(t *testing.T) {
	mockHostRepo := new(mocks.Repository)
	mockHost := models.Host{
		Name:    "Hello",
		Address: "Content",
	}

	t.Run("success", func(t *testing.T) {
		mockHostRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(&mockHost, nil).Once()
		u := ucase.NewHostUsecase(mockHostRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockHost.ID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockHostRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockHostRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("Unexpected")).Once()

		u := ucase.NewHostUsecase(mockHostRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockHost.ID)

		assert.Error(t, err)
		assert.Nil(t, a)

		mockHostRepo.AssertExpectations(t)
	})

}

func TestStore(t *testing.T) {
	mockHostRepo := new(mocks.Repository)
	mockHost := models.Host{
		Name:    "Hello",
		Address: "Content",
	}

	t.Run("success", func(t *testing.T) {
		tempMockHost := mockHost
		tempMockHost.ID = "adhfgasjklsjb"
		mockHostRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(nil, nil).Once()
		mockHostRepo.On("Store", mock.Anything, mock.AnythingOfType("*models.Host")).Return(nil).Once()

		u := ucase.NewHostUsecase(mockHostRepo, time.Second*2)

		h, err := u.Store(context.TODO(), &tempMockHost)

		assert.NoError(t, err)
		assert.NotNil(t, h)
		assert.Equal(t, mockHost.Name, tempMockHost.Name)
		mockHostRepo.AssertExpectations(t)
	})
	t.Run("existing-title", func(t *testing.T) {
		existingHost := mockHost
		mockHostRepo.On("GetByAddress", mock.Anything, mock.AnythingOfType("string")).Return(&existingHost, nil).Once()
		mockHostRepo.On("Store", mock.Anything, mock.AnythingOfType("*models.Host")).Return(errors.New("error")).Once()
		u := ucase.NewHostUsecase(mockHostRepo, time.Second*2)

		h, err := u.Store(context.TODO(), &mockHost)

		assert.Nil(t, h)
		assert.Error(t, err)
		mockHostRepo.AssertExpectations(t)
	})

}
