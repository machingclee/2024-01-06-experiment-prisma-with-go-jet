package pgsql

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://pguser:pguser@localhost:5432/udemy?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	return DB
}
