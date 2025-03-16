package services

import (
	"context"
	"database/sql"
	"myStandardModel/models"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	query := `INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id, username, email, created_at`
	row := s.db.QueryRowContext(ctx, query, user.Username, user.Email)

	var newUser models.User
	err := row.Scan(&newUser.ID, &newUser.Username, &newUser.Email, &newUser.CreatedAt)
	return newUser, err
}

func (s *UserService) GetUserByID(ctx context.Context, id int64) (models.User, error) {
	query := `SELECT id, username, email, created_at FROM users WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	return user, err
}
