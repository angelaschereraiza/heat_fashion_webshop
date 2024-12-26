package handlers

import (
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func ProductHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
			if len(pathParts) == 2 && pathParts[0] == "products" {
				handleProductDetail(w, pathParts[1])
				return
			}

			products := []models.Product{}
			err := db.DB.Select(&products, "SELECT * FROM products")
			if err != nil {
				log.Println("Error fetching products:", err)
				http.Error(w, "Error fetching products", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products)
			return
		}

		// for _, product := range products {
		// 	aliexpressapi.GetAliExpressDetails(cfg, &product)
		// }

		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func handleProductDetail(w http.ResponseWriter, id string) {
	product := models.Product{}
	err := db.DB.Get(&product, "SELECT * FROM products WHERE id = ?", id)
	if err != nil {
		log.Println("Product not found:", err)
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
