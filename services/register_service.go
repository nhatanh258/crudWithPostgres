package services

import (
	"context"
	"database/sql"
	"myStandardModel/models"
)

type RegisterService struct {
	db *sql.DB
}

func NewRegisterService(db *sql.DB) *RegisterService {
	return &RegisterService{db: db}
}

func (s *RegisterService) RegisterUserForEvent(ctx context.Context, register models.Register) (models.Register, error) {
	query := `INSERT INTO register (userId, eventid) VALUES ($1, $2) RETURNING id, userid, eventid`
	row := s.db.QueryRowContext(ctx, query, register.Userid, register.Eventid)

	var newRegister models.Register
	err := row.Scan(&newRegister.ID, &newRegister.Userid, &newRegister.Eventid)
	return newRegister, err
}
