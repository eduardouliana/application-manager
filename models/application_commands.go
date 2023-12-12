package models

type ApplicationCommands struct {
	ID            string `json:"id" gorm:"primaryKey"`
	ApplicationID string `json:"application_id"`
	Data          string `json:"data"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Cron          string `json:"cron"`
	CreatedAt     string `json:"created_at"`
	DeletedAt     string `json:"deleted_at"`
}
