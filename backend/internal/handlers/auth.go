package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	// "github.com/CarCleaningApp/backend/internal/models"
	// "github.com/CarCleaningApp/backend/pkg/logger"
	"github.com/CarCleaningApp/backend/internal/models"
	"github.com/CarCleaningApp/backend/pkg/logger"
)

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			logger.Logger.Println("Invalid request payload")
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		role, err := models.Authenticate(db, user.Email, user.Password)
		if err != nil {
			logger.Logger.Printf("Authentication failed for email: %s", user.Email)
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		logger.Logger.Printf("User authenticated: %s, Role: %s", user.Email, role)
		response := map[string]string{
			"message": "Login successful",
			"role":    role,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
