package models

type CommandExecutionLogs struct {
	ID               string `json:"id" gorm:"primaryKey"`
	PendingCommandID string `json:"pending_command_id" gorm:"not null"`
	Data             string `json:"data" gorm:"not null"`
}
