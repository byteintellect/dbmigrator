package database

import (
	"database/sql"
	"fmt"
	_ "gorm.io/driver/mysql"
	"log"
	"os"
	"strconv"
)

func Open() *sql.DB {
	databaseDsn := os.Getenv("DATABASE_DSN")
	if databaseDsn == "" {
		var databasePort int64 = 3306
		var err error
		port, ok := os.LookupEnv("DATABASE_PORT")
		if ok {
			databasePort, err = strconv.ParseInt(port, 10, 32)
			if err != nil {
				log.Fatalf("Invalid Database Port")
			}
		}
		databaseName := os.Getenv("DATABASE_NAME")
		databasePassword := os.Getenv("DATABASE_PASSWORD")
		databaseHost := os.Getenv("DATABASE_HOST")
		databaseUserName := os.Getenv("DATABASE_USER_NAME")
		databaseDsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%v?multiStatements=true&parseTime=true",
			databaseUserName,
			databasePassword,
			databaseHost,
			databasePort,
			databaseName)

	}
	db, err := sql.Open("mysql", databaseDsn)
	if err != nil {
		log.Fatalf("Error opening database connection %v\n", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error Pinging database: %v\n", err)
	}
	fmt.Printf("Connected to database\n")
	return db
}
