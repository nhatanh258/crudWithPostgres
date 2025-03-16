package repository

import (
	"context"
	"myStandardModel/models"
)

type RegisterRepository interface {
	GetRegistrationsByEventID(ctx context.Context, eventID int32) ([]models.Register, error)
	GetRegistrationsByUserID(ctx context.Context, userID int32) ([]models.Register, error)
	RegisterUserForEvent(ctx context.Context, register models.Register) (models.Register, error)
}
