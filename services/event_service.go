package services

import (
	"context"
	"database/sql"
	"myStandardModel/models"
	//"myStandardModel/repository"
)

type EventService struct {
	db *sql.DB
}

func NewEventService(db *sql.DB) *EventService {
	return &EventService{db: db}
}

func (s *EventService) CreateEvent(ctx context.Context, event models.Event) (models.Event, error) {
	query := `INSERT INTO events (name, description, location, "userId")
			  VALUES ($1, $2, $3, $4) RETURNING id, name, description, location, datatimecreating, "userId", datetimeupdating`

	row := s.db.QueryRowContext(ctx, query, event.Name, event.Description, event.Location, event.UserId)
	var newEvent models.Event
	err := row.Scan(&newEvent.ID, &newEvent.Name, &newEvent.Description, &newEvent.Location, &newEvent.Datatimecreating, &newEvent.UserId, &newEvent.Datetimeupdating)
	return newEvent, err
}

func (s *EventService) ListEvents(ctx context.Context) ([]models.Event, error) {
	query := `SELECT id, name, description, location, datatimecreating, "userId", datetimeupdating FROM events ORDER BY datatimecreating DESC`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datatimecreating, &event.UserId, &event.Datetimeupdating); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
