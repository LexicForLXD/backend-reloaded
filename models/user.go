package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Birthday  time.Time
	Email     string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type UserRepositoryInt interface {
	GetAllUsers() ([]*User, error)
	GetUser(ID string) (*User, error)
	CreateUser(user *User) error
	UpdateUser(ID string, user *User) error
	DeleteUser(ID string) error
	GetUserByToken(token string) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

type UserRepository struct {
	Db *gorm.DB
}

func (repo *UserRepository) GetAllUsers() ([]*User, error) {
	var users []*User

	if err := repo.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) GetUser(id string) (*User, error) {
	var user *User
	user.ID = id
	if err := repo.Db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) CreateUser(user *User) error {
	if err := repo.Db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) UpdateUser(id string, user *User) error {
	if err := repo.Db.Update(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) DeleteUser(user *User) error {
	if err := repo.Db.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetUserByToken(token string) (*User, error) {
	var user *User
	user.Token = token
	if err := repo.Db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetUserByEmail(email string) (*User, error) {
	var user *User
	user.Email = email
	if err := repo.Db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
