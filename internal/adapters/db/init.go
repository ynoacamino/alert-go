package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB(DATABASE_URL string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DATABASE_URL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Default().Println("[LOGER] [DB]: Database connected")

	return db, nil
}
