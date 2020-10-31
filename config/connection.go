package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "your_username:your_password@tcp(your_IP:3306)/your_database")
	if err != nil {
		log.Fatal(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}