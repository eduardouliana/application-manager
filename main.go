package main

import (
	"appmanager/database"
	"appmanager/migrations"
	"appmanager/server"
)

func main() {
	database.StartDB()

	migrations.RunMigrations(database.GetDB())

	server := server.NewServer("8080")

	server.Run()
}
