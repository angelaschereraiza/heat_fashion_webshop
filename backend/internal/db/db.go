package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDB(dataSourceName string) {
	var err error
	DB, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}

	runMigrations()

	if err = DB.Ping(); err != nil {
		log.Fatalln(err)
	}
}

func runMigrations() {
	sqlDB := DB.DB

	driver, err := mysql.WithInstance(sqlDB, &mysql.Config{})
	if err != nil {
		log.Fatalf("Failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"mysql", driver)
	if err != nil {
		log.Fatalf("Failed to start migrations: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Database migration completed")
}
