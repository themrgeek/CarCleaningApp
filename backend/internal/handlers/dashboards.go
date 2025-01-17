package handlers

import (
	"net/http"
)

// UserDashboard handles the user dashboard requests

func UserDashboard(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("User Dashboard"))

}

// CleanerDashboard handles the user dashboard requests

func CleanerDashboard(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Cleaner Dashboard"))

}
func AdminDashboard(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Admin Dashboard"))

}
