package repository

import (
	"context"
	"myStandardModel/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	GetUserByID(ctx context.Context, id int64) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	ListUsers(ctx context.Context, limit, offset int32) ([]models.User, error)
	UpdateUser(ctx context.Context, user models.User) (models.User, error)
	DeleteUser(ctx context.Context, id int64) error
}
