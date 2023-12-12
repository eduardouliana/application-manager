package controllers

import (
	"appmanager/database"
	"appmanager/models"

	"github.com/gin-gonic/gin"
)

func RegisterApplication(c *gin.Context) {
	var application models.Applications

	err := c.ShouldBindJSON(&application)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to bind JSON! Error: " + err.Error(),
		})
		return
	}

	db := database.GetDB()
	err = db.Create(&application).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to create application! Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully created application!",
	})
}

func GetApplications(c *gin.Context) {
	var applications []models.Applications

	db := database.GetDB()
	err := db.Find(&applications).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get applications! Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, applications)
}
