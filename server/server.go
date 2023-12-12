package server

import (
	"appmanager/server/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Print("Server running on port: " + s.port)
	log.Fatal(router.Run(":" + s.port))
}
