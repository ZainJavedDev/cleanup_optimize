package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	pgsqlDB := os.Getenv("PGSQL_DB")
	pgsqlUser := os.Getenv("PGSQL_USER")
	pgsqlPassword := os.Getenv("PGSQL_PASSWORD")
	pgsqlHost := os.Getenv("PGSQL_HOST")
	pgsqlPort := os.Getenv("PGSQL_PORT")

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pgsqlHost, pgsqlPort, pgsqlUser, pgsqlPassword, pgsqlDB))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
