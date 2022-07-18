package database

import (
	"database/sql"
	"fmt"
	_ "gorm.io/driver/mysql"
	"os"
)

func Open() *sql.DB {
	databaseDsn := os.Getenv("DATABASE_DSN")
	if databaseDsn != "" {

	} else {
		databaseName := os.Getenv("DATABASE_NAME")
		databasePassword := os.Getenv("DATABASE_PASSWORD")
		databaseHost := os.Getenv("DATABASE_HOST")
		databaseUserName := os.Getenv("DATABASE_USER_NAME")
		databasePort := os.Getenv("DATABASE_PORT")
		databaseDsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?multiStatements=true&parseTime=true",
			databaseUserName,
			databasePassword,
			databaseHost,
			databasePort,
			databaseName)

	}
	db, err := sql.Open("mysql", databaseDsn)
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
