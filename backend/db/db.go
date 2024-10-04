package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDB(dataSourceName string) {
	var err error
	DB, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalln(err)
	}
}
