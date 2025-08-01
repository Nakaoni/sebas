package util

import (
	"context"
	"database/sql"
	"log"
)

type databaseConfig struct {
	driverName     string
	dataSourceName string
}

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := db.PingContext(context.Background()); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
