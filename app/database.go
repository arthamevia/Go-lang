package app

import (
	"database/sql"
	"time"
	"auth/helper"
)

func Database() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/auth")
	helper.PanicError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}