package database

import (
	"database/sql"
	"fmt"
	_ "gorm.io/driver/mysql"
	"os"
)

func Open() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_DSN"))
	if err != nil {
		fmt.Printf("Error opening database connection %v\n", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error Pinging database: %v\n", err)
	}
	fmt.Printf("Connected to database\n")
	return db
}
