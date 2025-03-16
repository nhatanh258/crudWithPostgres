package models

import "time"

type User struct {
	ID        int64
	Username  string `binding:"required"`
	Email     string `binding:"required"`
	CreatedAt time.Time
}