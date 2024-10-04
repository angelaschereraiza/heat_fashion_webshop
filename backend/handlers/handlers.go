package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"backend/db"
	"backend/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to HeatUp Fashion API!")
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Product ID is missing", http.StatusBadRequest)
		return
	}
	productID := parts[2]

	var product models.Product
	err := db.DB.Get(&product, "SELECT * FROM products WHERE id = ?", productID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	fmt.Fprintf(w, "Product: %s, Price: %.2f", product.Name, product.Price)
}
