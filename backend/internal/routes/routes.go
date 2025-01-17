package routes

import (
	"database/sql"
	"log"

	"github.com/gorilla/mux"
	"github.com/themrgeek/CarCleaningApp/backend/internal/handlers"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", handlers.LoginHandler(db)).Methods("POST")

	// Admin-specific route
	router.HandleFunc("/admin/dashboard", handlers.AdminDashboard).Methods("GET")

	// User-specific route
	router.HandleFunc("/user/dashboard", handlers.UserDashboard).Methods("GET")

	// Cleaner-specific route
	router.HandleFunc("/cleaner/dashboard", handlers.CleanerDashboard).Methods("GET")

	log.Println("Routes initialized successfully")
	return router
}
