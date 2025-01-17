package routes

import (
	"log"

	"github.com/themrgeek/CarCleaningApp/backend/handlers"

	"github.com/gorilla/mux"
)

// SetupRoutes sets up all routes for the application
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Route for login
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Admin-specific route
	router.HandleFunc("/admin/dashboard", handlers.AdminDashboard).Methods("GET")

	// User-specific route
	router.HandleFunc("/user/dashboard", handlers.UserDashboard).Methods("GET")

	// Cleaner-specific route
	router.HandleFunc("/cleaner/dashboard", handlers.CleanerDashboard).Methods("GET")

	log.Println("Routes initialized successfully")
	return router
}
