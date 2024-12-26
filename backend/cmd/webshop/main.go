package main

import (
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/handlers"
	"backend/internal/middleware"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DbName)

	db.ConnectDB(connectionString)

	http.HandleFunc("/products/", middleware.WithCORS(handlers.ProductHandler(cfg)))
	http.HandleFunc("/cart/", middleware.WithCORS(handlers.CartHandler()))
	http.HandleFunc("/register/", middleware.WithCORS(handlers.RegisterHandler()))
	http.HandleFunc("/login/", middleware.WithCORS(handlers.LoginHandler()))
	http.HandleFunc("/logout/", middleware.WithCORS(handlers.LogoutHandler()))

	log.Println("Server is running on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
