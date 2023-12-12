package controllers

import (
	"appmanager/database"
	"appmanager/models"

	"github.com/gin-gonic/gin"
)

func GetAllCommands(c *gin.Context) {
	var commands []models.ApplicationCommands

	db := database.GetDB()
	err := db.Find(&commands).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get commands! Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, commands)
}

func GetCommand(c *gin.Context) {
	var command models.ApplicationCommands

	db := database.GetDB()
	err := db.Where("id = ?", c.Param("id")).First(&command).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get command! Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, command)
}

func AddCommand(c *gin.Context) {
	var command models.ApplicationCommands

	err := c.ShouldBindJSON(&command)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to bind JSON! Error: " + err.Error(),
		})
		return
	}

	db := database.GetDB()
	err = db.Create(&command).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to create command! Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully created command!",
	})
}

func UpdateCommand(c *gin.Context) {
	var command models.ApplicationCommands

	err := c.ShouldBindJSON(&command)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to bind JSON! Error: " + err.Error(),
		})
		return
	}

	db := database.GetDB()
	err = db.Model(&command).Where("id = ?", c.Param("id")).Updates(&command).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to update command! Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully updated command!",
	})
}

func DeleteCommand(c *gin.Context) {
	var command models.ApplicationCommands

	db := database.GetDB()
	err := db.Where("id = ?", c.Param("id")).Delete(&command).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to delete command! Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully deleted command!",
	})
}
