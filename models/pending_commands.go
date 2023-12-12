package models

import (
	"database/sql"
	"time"
)

type PendingCommands struct {
	ID                   string       `json:"id" gorm:"primaryKey"`
	ApplicationCommandID string       `json:"application_command_id" gorm:"not null"`
	CreatedAt            time.Time    `json:"created_at" gorm:"not null"`
	ReadedAt             sql.NullTime `json:"readed_at"`
}
