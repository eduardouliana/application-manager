package models

import (
	"time"
)

type Applications struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	LastUpdate  time.Time `json:"last_update" gorm:"autoUpdateTime"`
}
