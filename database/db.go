package database

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	var err error
	db, err = sql.Open("mysql", "root@tcp(localhost:3306)/web-service-gin")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func GetDB() (*sql.DB, error) {
	if db == nil {
		return nil, errors.New("database connection is not initialized")
	}
	return db, nil
}
