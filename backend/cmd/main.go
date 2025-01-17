package main

import (
	"log"
	"net/http"

	"github.com/themrgeek/CarCleaningApp/backend/config"
	"github.com/themrgeek/CarCleaningApp/backend/pkg/db"
	"github.com/themrgeek/CarCleaningApp/backend/pkg/logger"
	"github.com/themrgeek/CarCleaningApp/backend/routes"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize logger
	logger.InitLogger(cfg.LogFile)

	// Connect to the database
	database := db.Connect(cfg)
	defer database.Close()

	// Register routes
	router := routes.RegisterRoutes(database)

	log.Printf("Server running on port %s...", cfg.AppPort)
	logger.Logger.Printf("Server started on port %s", cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, router))
}
