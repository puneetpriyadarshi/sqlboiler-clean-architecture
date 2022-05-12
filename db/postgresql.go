package db

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 8080
	user     = "postgres"
	password = "Singhasan26!"
	dbname   = "postgres"
)

func Connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Successfully connected")

	return db, nil
}
