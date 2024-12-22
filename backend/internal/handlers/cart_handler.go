package handlers

import (
	"backend/internal/db"
	"backend/internal/models"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func CartHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handleGetCart(w, r)
			return
		}
		if r.Method == http.MethodPost {
			handleAddToCart(w, r)
			return
		}
		if r.Method == http.MethodDelete {
			handleRemoveFromCart(w, r)
			return
		}
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func handleGetCart(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	cart := models.Cart{}
	err := db.DB.Get(&cart, "SELECT id, user_id FROM carts WHERE user_id = ?", userID)
	if err != nil {
		http.Error(w, "Cart not found", http.StatusNotFound)
		return
	}
	cartItems := []models.CartItem{}
	err = db.DB.Select(&cartItems, "SELECT id, cart_id, product_id, quantity FROM cart_items WHERE cart_id = ?", cart.ID)
	if err != nil {
		http.Error(w, "Error fetching cart items", http.StatusInternalServerError)
		return
	}
	cart.CartItems = cartItems
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func handleAddToCart(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	var item models.CartItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	cart := models.Cart{}
	err := db.DB.Get(&cart, "SELECT id FROM carts WHERE user_id = ?", userID)
	if err != nil {
		cartID := uuid.New()
		_, err := db.DB.Exec("INSERT INTO carts (id, user_id) VALUES (?, ?)", cartID, userID)
		if err != nil {
			http.Error(w, "Error creating cart", http.StatusInternalServerError)
			return
		}
		cart.ID = cartID
	}
	existingItem := models.CartItem{}
	err = db.DB.Get(&existingItem, "SELECT id, quantity FROM cart_items WHERE cart_id = ? AND product_id = ?", cart.ID, item.ProductID)
	if err == nil {
		_, err := db.DB.Exec("UPDATE cart_items SET quantity = ? WHERE id = ?", existingItem.Quantity+item.Quantity, existingItem.ID)
		if err != nil {
			http.Error(w, "Error updating cart item", http.StatusInternalServerError)
			return
		}
	} else {
		item.ID = uuid.New()
		item.CartID = cart.ID
		_, err := db.DB.Exec("INSERT INTO cart_items (id, cart_id, product_id, quantity) VALUES (?, ?, ?, ?)", item.ID, item.CartID, item.ProductID, item.Quantity)
		if err != nil {
			http.Error(w, "Error adding product to cart", http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product added to cart"})
}

func handleRemoveFromCart(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) != 2 || pathParts[0] != "cart" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	productID := pathParts[1]
	cart := models.Cart{}
	err := db.DB.Get(&cart, "SELECT id FROM carts WHERE user_id = ?", userID)
	if err != nil {
		http.Error(w, "Cart not found", http.StatusNotFound)
		return
	}
	_, err = db.DB.Exec("DELETE FROM cart_items WHERE cart_id = ? AND product_id = ?", cart.ID, productID)
	if err != nil {
		http.Error(w, "Error removing product from cart", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product removed from cart"})
}
