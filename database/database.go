package database

import (
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/simple_restful_api_golang")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)

	return db
}
