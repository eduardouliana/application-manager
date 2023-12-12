package routes

import (
	"appmanager/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	applications := router.Group("/applications")
	{
		applications.GET("/", controllers.GetApplications)
		applications.POST("/", controllers.RegisterApplication)
	}

	return router
}
