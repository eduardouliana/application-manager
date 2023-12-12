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

	pendingCommands := router.Group("/pending-commands")
	{
		pendingCommands.GET("/", controllers.GetAllPendingCommands)
		pendingCommands.GET("/:id", controllers.GetPendingCommand)
		pendingCommands.POST("/", controllers.AddPendingCommand)
		pendingCommands.PUT("/:id", controllers.UpdatePendingCommand)
		pendingCommands.DELETE("/:id", controllers.DeletePendingCommand)
	}

	return router
}
