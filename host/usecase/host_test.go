package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/lexicforlxd/backend-reloaded/host/mocks"
	"github.com/lexicforlxd/backend-reloaded/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShow(t *testing.T) {
	mockHostRepo := new(mocks.Repository)
	mockHost := models.Host{
		Name:    "Hello",
		Address: "Content",
	}

	t.Run("success", func(t *testing.T) {
		mockHostRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(&mockHost, nil).Once()
		u := NewHostUsecase(mockHostRepo, time.Second*2)

		a, err := u.Show(context.TODO(), mockHost.ID)

		assert.NoError(t, err)
		assert.NotNil(t, a)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockHostRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("Unexpected")).Once()

		u := NewHostUsecase(mockHostRepo, time.Second*2)

		a, err := u.Show(context.TODO(), mockHost.ID)

		assert.Error(t, err)
		assert.Nil(t, a)
	})

	t.Run("assert mock", func(t *testing.T) {
		mockHostRepo.AssertExpectations(t)
	})

}

func TestStore(t *testing.T) {
	mockHostRepo := new(mocks.Repository)
	mockHost := models.Host{
		Name:    "Hello",
		Address: "Content",
	}

	t.Run("success without password", func(t *testing.T) {
		tempMockHost := mockHost
		mockHostRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(&tempMockHost, nil).Once()
		mockHostRepo.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("error")).Once()
		mockHostRepo.On("GetByAddress", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("error")).Once()
		mockHostRepo.On("Store", mock.Anything, mock.AnythingOfType("*models.Host")).Return(nil).Once()

		u := NewHostUsecase(mockHostRepo, time.Second*2)

		h, err := u.Store(context.TODO(), &tempMockHost)

		assert.NoError(t, err)
		assert.NotNil(t, h)
		assert.Equal(t, mockHost.Name, h.Name)
	})

	t.Run("success with password", func(t *testing.T) {
		tempMockHost := mockHost
		tempMockHost.Password = "password"
		mockHostRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(&tempMockHost, nil).Once()
		mockHostRepo.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("error")).Once()
		mockHostRepo.On("GetByAddress", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("error")).Once()
		mockHostRepo.On("Store", mock.Anything, mock.AnythingOfType("*models.Host")).Return(nil).Once()

		u := NewHostUsecase(mockHostRepo, time.Second*2)

		h, err := u.Store(context.TODO(), &tempMockHost)

		assert.NoError(t, err)
		assert.NotNil(t, h)
		assert.Equal(t, mockHost.Name, h.Name)
	})

	t.Run("failing existing address", func(t *testing.T) {
		mockHostRepo.On("GetByAddress", mock.Anything, mock.AnythingOfType("string")).Return(&mockHost, nil).Once()
		u := NewHostUsecase(mockHostRepo, time.Second*2)

		h, err := u.Store(context.TODO(), &mockHost)

		assert.Nil(t, h)
		assert.Error(t, err)
	})

	t.Run("failing existing name", func(t *testing.T) {
		mockHostRepo.On("GetByAddress", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("error")).Once()
		mockHostRepo.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(&mockHost, nil).Once()
		u := NewHostUsecase(mockHostRepo, time.Second*2)

		h, err := u.Store(context.TODO(), &mockHost)

		assert.Nil(t, h)
		assert.Error(t, err)
	})

	t.Run("assert mock", func(t *testing.T) {
		mockHostRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockHostRepo := new(mocks.Repository)
	mockHost := models.Host{
		Name:    "Hello",
		Address: "Content",
	}

	t.Run("success", func(t *testing.T) {
		tempMockHost := mockHost
		tempMockHost.ID = "adhfgasjklsjb"
		tempMockHost.Name = "Help"

		mockHostRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(&mockHost, nil).Once()
		mockHostRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(&tempMockHost, nil).Once()
		mockHostRepo.On("Update", mock.Anything, mock.AnythingOfType("*models.Host")).Return(nil).Once()

		u := NewHostUsecase(mockHostRepo, time.Second*2)

		h, err := u.Update(context.TODO(), &tempMockHost)

		assert.NoError(t, err)
		assert.NotNil(t, h)
		assert.NotEqual(t, mockHost.Name, h.Name)
		assert.Equal(t, tempMockHost.Name, h.Name)
	})

	t.Run("not-found", func(t *testing.T) {
		mockHostRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("not found")).Once()
		u := NewHostUsecase(mockHostRepo, time.Second*2)
		h, err := u.Update(context.TODO(), &mockHost)

		assert.Error(t, err)
		assert.Nil(t, h)
	})

	t.Run("assert mock", func(t *testing.T) {
		mockHostRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockHostRepo := new(mocks.Repository)
	mockHost := models.Host{
		Name:    "Hello",
		Address: "Content",
	}

	t.Run("success", func(t *testing.T) {
		mockHostRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(&mockHost, nil).Once()
		mockHostRepo.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		u := NewHostUsecase(mockHostRepo, time.Second*2)
		err := u.Delete(context.TODO(), "asuidgz")

		assert.NoError(t, err)
	})

	t.Run("not-found", func(t *testing.T) {
		mockHostRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("not found")).Once()

		u := NewHostUsecase(mockHostRepo, time.Second*2)
		err := u.Delete(context.TODO(), "asuidgz")

		assert.Error(t, err)
	})

	t.Run("assert mock", func(t *testing.T) {
		mockHostRepo.AssertExpectations(t)
	})

}

func TestFetch(t *testing.T) {
	mockHostRepo := new(mocks.Repository)
	mockHost := models.Host{
		Name:    "Hello",
		Address: "Content",
	}

	mockHosts := []*models.Host{
		&mockHost,
	}

	t.Run("success", func(t *testing.T) {
		mockHostRepo.On("Fetch", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(mockHosts, nil).Once()

		u := NewHostUsecase(mockHostRepo, time.Second*2)
		h, err := u.Fetch(context.TODO(), -1, -1)

		assert.NotNil(t, h)
		assert.NoError(t, err)
	})

	t.Run("empty-db", func(t *testing.T) {
		mockHostRepo.On("Fetch", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, errors.New("no hosts")).Once()

		u := NewHostUsecase(mockHostRepo, time.Second*2)
		h, err := u.Fetch(context.TODO(), -1, -1)

		assert.Nil(t, h)
		assert.Error(t, err)
	})

	t.Run("assert mock", func(t *testing.T) {
		mockHostRepo.AssertExpectations(t)
	})
}
