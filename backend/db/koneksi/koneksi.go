package koneksi

import (
	"database/sql"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func GetConnection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dbname := os.Getenv("DB_NAME")
	db, err := sql.Open("sqlite3", `db/migration/`+dbname+`?_loc=Local`)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db
}
func GetConnection1() *sql.DB {
	db, err := sql.Open("sqlite3", "../db/migration/app1.db?_loc=Local")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db
}
