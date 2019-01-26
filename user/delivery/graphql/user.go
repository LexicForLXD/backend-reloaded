package graphql

import (
	"context"

	"github.com/google/uuid"
	"github.com/vektah/gqlparser/gqlerror"
	"gitlab.com/lexicforlxd/backend-reloaded/host"
	"gitlab.com/lexicforlxd/backend-reloaded/models"
	"golang.org/x/crypto/bcrypt"
)

type UserResolver struct {
}

func NewUserResolver(h host.Usecase) UserResolver {
	resolver := UserResolver{
		// hostUsecase: h,
	}
	return resolver
}

// Mutations
func (r *UserResolver) CreateUser(ctx context.Context, userReq models.UserReq) (*models.User, error) {
	user := models.User{
		ID:        uuid.New().String(),
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		Email:     userReq.Email,
		Token:     uuid.New().String(),
	}

	if userReq.Password != nil {
		// user.Password = *userReq.Desc
	}

	if userReq.Birthday != nil {
		user.Birthday = *userReq.Birthday
	}

	if userReq.Password != nil {
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(*userReq.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPass)
	}

	// if err := r.db.Create(&user).Error; err != nil {
	// 	return nil, err
	// }

	return &user, nil
}

func (r *UserResolver) UpdateUser(ctx context.Context, id string, userUpdate models.UserUpdate) (*models.User, error) {
	user := models.User{}
	// if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
	// 	return nil, err
	// }

	if userUpdate.FirstName != nil {
		user.FirstName = *userUpdate.FirstName
	}

	if userUpdate.LastName != nil {
		user.LastName = *userUpdate.LastName
	}

	if userUpdate.Email != nil {
		user.Email = *userUpdate.Email
	}

	if userUpdate.Birthday != nil {
		user.Birthday = *userUpdate.Birthday
	}

	if userUpdate.Password != nil {
		if userUpdate.Token != nil && *userUpdate.Token != user.Token {
			return nil, gqlerror.Errorf("Invalid token")
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*userUpdate.OldPassword)); err != nil {
			return nil, gqlerror.Errorf("Password missmatch")
		}

		if hashedPass, err := bcrypt.GenerateFromPassword([]byte(*userUpdate.OldPassword), bcrypt.DefaultCost); err != nil {
			return nil, err
		} else {
			user.Password = string(hashedPass)
		}

	}

	// if err := r.db.Save(&user).Error; err != nil {
	// 	return nil, err
	// }

	return &user, nil
}
func (r *UserResolver) DeleteUser(ctx context.Context, id string) (*models.DeleteRes, error) {
	panic("not implemented")
}

// Queries

func (r *UserResolver) Users(ctx context.Context, limit *int) ([]*models.User, error) {
	var users []*models.User

	if limit != nil {
		// if err := r.db.Limit(*limit).Find(&users).Error; err != nil {
		// 	return nil, err
		// }
	} else {
		// if err := r.db.Find(&users).Error; err != nil {
		// 	return nil, err
		// }
	}

	return users, nil
}
func (r *UserResolver) User(ctx context.Context, id string) (*models.User, error) {
	user := models.User{}
	// if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
	// 	return nil, err
	// }

	return &user, nil
}
