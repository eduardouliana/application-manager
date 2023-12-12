package controllers

import (
	"appmanager/database"
	"appmanager/models"

	"github.com/gin-gonic/gin"
)

func GetAllApplications(c *gin.Context) {
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

func GetApplication(c *gin.Context) {
	var application models.Applications

	db := database.GetDB()
	err := db.Where("id = ?", c.Param("id")).First(&application).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get application! Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, application)
}

func AddApplication(c *gin.Context) {
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

func UpdateApplication(c *gin.Context) {
	var application models.Applications

	err := c.ShouldBindJSON(&application)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to bind JSON! Error: " + err.Error(),
		})
		return
	}

	db := database.GetDB()
	err = db.Save(&application).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to update application! Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully updated application!",
	})
}

func DeleteApplication(c *gin.Context) {
	var application models.Applications

	db := database.GetDB()
	err := db.Where("id = ?", c.Param("id")).Delete(&application).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to delete application! Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully deleted application!",
	})
}
