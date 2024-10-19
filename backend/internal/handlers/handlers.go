package handlers

import (
	"backend/internal/db"
	"backend/internal/models"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {

	// Handle GET requests to return all products
	if r.Method == http.MethodGet {
		products := []models.Product{}
		err := db.DB.Select(&products, "SELECT * FROM products")
		if err != nil {
			http.Error(w, "Error fetching products", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
		return
	}

	// Handle POST requests to create a new product
	if r.Method == http.MethodPost {
		var product models.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		product.ID = uuid.New()

		_, err = db.DB.Exec("INSERT INTO products (id, name, price) VALUES (?, ?, ?)", product.ID, product.Name, product.Price)
		if err != nil {
			http.Error(w, "Error inserting product", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(product)
		return
	}

	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}
