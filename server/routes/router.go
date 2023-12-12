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

	applicationCommands := router.Group("/application-commands")
	{
		applicationCommands.GET("/", controllers.GetAllCommands)
		applicationCommands.GET("/:id", controllers.GetCommand)
		applicationCommands.POST("/", controllers.AddCommand)
		applicationCommands.PUT("/:id", controllers.UpdateCommand)
		applicationCommands.DELETE("/:id", controllers.DeleteCommand)
	}
	}

	return router
}
