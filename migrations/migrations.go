package migrations

import (
	"appmanager/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Applications{})
	db.AutoMigrate(models.ApplicationCommands{})
	db.AutoMigrate(models.PendingCommands{})
	db.AutoMigrate(models.CommandExecutionLogs{})
}
