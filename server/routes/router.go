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
		applications.GET("/", controllers.GetAllApplications)
		applications.GET("/:id", controllers.GetApplication)
		applications.POST("/", controllers.AddApplication)
		applications.PUT("/:id", controllers.UpdateApplication)
		applications.DELETE("/:id", controllers.DeleteApplication)
	}

	}

	return router
}
