package repository

import (
	"context"
	"errors"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lexicforlxd/backend-reloaded/models"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func newDB() (sqlmock.Sqlmock, *gorm.DB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}

	gormDB, gerr := gorm.Open("postgres", db)
	if gerr != nil {
		log.Fatalf("can't open gorm connection: %s", err)
	}
	gormDB.LogMode(true)

	// defer db.Close()
	return mock, gormDB.Set("gorm:update_column", true)
}

func TestGetByID(t *testing.T) {
	mock, gorm := newDB()
	mockHost := models.Host{
		ID:      "ahdsgfjs",
		Name:    "name1",
		Address: "192.168.1.2",
	}
	query := "SELECT * FROM \"hosts\""
	rows := sqlmock.NewRows([]string{"id", "name", "desc", "address", "password", "authenticated", "containers", "created_at", "updated_at", "deleted_at"}).
		AddRow(mockHost.ID, mockHost.Name, nil, mockHost.Address, nil, false, nil, time.Now(), time.Now(), nil)

	hr := NewHostRepository(gorm)

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(mockHost.ID).WillReturnRows(rows)

		resultHost, err := hr.GetByID(context.TODO(), mockHost.ID)
		assert.NoError(t, err)
		assert.NotNil(t, resultHost)
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(mockHost.ID).WillReturnError(errors.New("not found"))

		resultHost, err := hr.GetByID(context.TODO(), mockHost.ID)
		assert.Error(t, err)
		assert.Nil(t, resultHost)
	})

}

func TestGetByName(t *testing.T) {
	mock, gorm := newDB()
	mockHost := models.Host{
		ID:      "ahdsgfjs",
		Name:    "name1",
		Address: "192.168.1.2",
	}
	query := "SELECT * FROM \"hosts\""
	rows := sqlmock.NewRows([]string{"id", "name", "desc", "address", "password", "authenticated", "containers", "created_at", "updated_at", "deleted_at"}).
		AddRow(mockHost.ID, mockHost.Name, nil, mockHost.Address, nil, false, nil, time.Now(), time.Now(), nil)

	hr := NewHostRepository(gorm)

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(mockHost.Name).WillReturnRows(rows)

		resultHost, err := hr.GetByName(context.TODO(), mockHost.Name)
		assert.NoError(t, err)
		assert.NotNil(t, resultHost)
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(mockHost.Name).WillReturnError(errors.New("not found"))

		resultHost, err := hr.GetByName(context.TODO(), mockHost.Name)
		assert.Error(t, err)
		assert.Nil(t, resultHost)
	})
}

func TestGetByAddress(t *testing.T) {
	mock, gorm := newDB()
	mockHost := models.Host{
		ID:      "ahdsgfjs",
		Name:    "name1",
		Address: "192.168.1.2",
	}
	query := "SELECT * FROM \"hosts\""
	rows := sqlmock.NewRows([]string{"id", "name", "desc", "address", "password", "authenticated", "containers", "created_at", "updated_at", "deleted_at"}).
		AddRow(mockHost.ID, mockHost.Name, nil, mockHost.Address, nil, false, nil, time.Now(), time.Now(), nil)

	hr := NewHostRepository(gorm)

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(mockHost.Name).WillReturnRows(rows)

		resultHost, err := hr.GetByAddress(context.TODO(), mockHost.Name)
		assert.NoError(t, err)
		assert.NotNil(t, resultHost)
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(mockHost.Name).WillReturnError(errors.New("not found"))

		resultHost, err := hr.GetByAddress(context.TODO(), mockHost.Name)
		assert.Error(t, err)
		assert.Nil(t, resultHost)
	})
}
