package repository

import (
	"context"
	"myStandardModel/models"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, event models.Event) (models.Event, error)
	GetEventByID(ctx context.Context, id int64) (models.Event, error)
	ListEvents(ctx context.Context) ([]models.Event, error)
	UpdateEvent(ctx context.Context, event models.Event) (models.Event, error)
	DeleteEvent(ctx context.Context, id int64) error
}
