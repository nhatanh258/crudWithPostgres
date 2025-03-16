package models

import "time"

type Event struct {
	ID               int64
	Name             string `binding:"required"`
	Description      string `binding:"required"`
	Location         string `binding:"required"`
	Datatimecreating time.Time
	UserId           int32 
	Datetimeupdating time.Time
}