package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_HOST          = "172.17.0.1"
	DB_PORT          = 5432
	DB_USER          = "postgres"
	DB_PASSWORD      = "postgres10"
	DB_DATABASE_NAME = "demo_connect_golang"
)

func GetPostgresDb() (*sql.DB, error) {
	desc := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_DATABASE_NAME)

	db, err := ConnectToDb(desc)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectToDb(desc string) (*sql.DB, error) {
	db, err := sql.Open("postgres", desc)
	if err != nil {
		return nil, err
	}

	return db, nil
}
