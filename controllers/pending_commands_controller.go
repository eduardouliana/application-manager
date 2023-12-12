package controllers

import (
	"appmanager/database"
	"appmanager/models"

	"github.com/gin-gonic/gin"
)

func GetAllPendingCommands(c *gin.Context) {
	var pendingCommands []models.PendingCommands

	db := database.GetDB()
	err := db.Find(&pendingCommands).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get pending commands! Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, pendingCommands)
}

func GetPendingCommand(c *gin.Context) {
	var pendingCommand models.PendingCommands

	db := database.GetDB()
	err := db.Where("id = ?", c.Param("id")).First(&pendingCommand).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get pending command! Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, pendingCommand)
}

func AddPendingCommand(c *gin.Context) {
	var pendingCommand models.PendingCommands

	err := c.ShouldBindJSON(&pendingCommand)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to bind JSON! Error: " + err.Error(),
		})
		return
	}

	db := database.GetDB()
	err = db.Create(&pendingCommand).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to create pending command! Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully added pending command!",
	})
}

func UpdatePendingCommand(c *gin.Context) {
	var pendingCommand models.PendingCommands

	db := database.GetDB()
	err := db.Where("id = ?", c.Param("id")).First(&pendingCommand).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get pending command! Error: " + err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&pendingCommand)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to bind JSON! Error: " + err.Error(),
		})
		return
	}

	err = db.Save(&pendingCommand).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to update pending command! Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully updated pending command!",
	})
}

func DeletePendingCommand(c *gin.Context) {
	var pendingCommand models.PendingCommands

	db := database.GetDB()
	err := db.Where("id = ?", c.Param("id")).First(&pendingCommand).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get pending command! Error: " + err.Error(),
		})
		return
	}

	err = db.Delete(&pendingCommand).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to delete pending command! Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully deleted pending command!",
	})
}
