package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"go-http-backend-server/pkg/db"
)

type UserHandler struct {
	DB *db.Database
}

func NewUserHandler(database *db.Database) *UserHandler {
	return &UserHandler{DB: database}
}

// Register handles user registration
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Insert into database
	query := `INSERT INTO users (username, password) VALUES (?, ?)`
	_, err = h.DB.Conn.Exec(query, data.Username, data.Password)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to register user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully."))

}

// Handles user login
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Query database
	var storedPassword string
	query := `SELECT password FROM users WHERE username = ?`
	err = h.DB.Conn.QueryRow(query, data.Username).Scan(&storedPassword)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if data.Password != storedPassword {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))

}