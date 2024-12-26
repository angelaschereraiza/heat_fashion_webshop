package handlers

import (
	"backend/internal/db"
	"backend/internal/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
)

func RegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handleRegister(w, r)
			return
		}

		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handleLogin(w, r)
			return
		}

		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handleLogout(w, r)
			return
		}

		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Invalid request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if user.Mail == "" || user.Password == "" {
		http.Error(w, "Mail and password are required", http.StatusBadRequest)
		return
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user.ID = uuid.New()
	user.Password = hashedPassword

	_, err = db.DB.Exec("INSERT INTO users (id, mail, password, first_name, last_name, phone_number, address, city, postal_code, country) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		user.ID, user.Mail, user.Password, user.FirstName, user.LastName, user.PhoneNumber, user.Address, user.City, user.PostalCode, user.Country)
	if err != nil {
		log.Println("Error creating user:", err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Mail     string `json:"mail"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		log.Println("Invalid request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var user models.User
	err := db.DB.Get(&user, "SELECT id, mail, password FROM users WHERE mail = ?", credentials.Mail)
	if err != nil {
		log.Println("Invalid mail or password:", err)
		http.Error(w, "Invalid mail or password", http.StatusUnauthorized)
		return
	}

	if !verifyPassword(user.Password, credentials.Password) {
		http.Error(w, "Invalid mail or password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
}

func hashPassword(password string) (string, error) {
	salt := []byte("somesecurestaticrandomsalt")
	hashed := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	return string(hashed), nil
}

func verifyPassword(hashedPassword, password string) bool {
	salt := []byte("somesecurestaticrandomsalt")
	expectedHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	return string(expectedHash) == hashedPassword
}
