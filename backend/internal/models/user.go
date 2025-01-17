package models

import (
	"database/sql"
	"errors"
	"log"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"`
}

func Authenticate(db *sql.DB, email, password string) (string, error) {
	var role string
	query := "SELECT role FROM users WHERE email = ? AND password = ?"
	err := db.QueryRow(query, email, password).Scan(&role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Failed login attempt for email: %s", email)
			return "", errors.New("invalid credentials")
		}
		log.Printf("Error querying database for email %s: %v", email, err)
		return "", err
	}

	log.Printf("Successful login for email: %s, role: %s", email, role)
	return role, nil
}
