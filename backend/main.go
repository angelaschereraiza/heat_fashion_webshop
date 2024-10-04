package main

import (
	"fmt"
	"log"
	"net/http"

	"backend/config"
	"backend/db"
	"backend/handlers"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DbName)

	db.ConnectDB(connectionString)

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/products/", handlers.ProductHandler)

	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
