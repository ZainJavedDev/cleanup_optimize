package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	mysqlDB := os.Getenv("MYSQL_DB")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", mysqlUser, mysqlPassword, mysqlDB))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
