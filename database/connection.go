package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite3 Driver
)

// Open connects to database server
func Open() *sql.DB {

	db, err := sql.Open("sqlite3", DBFile)
	if err != nil {
		log.Fatal(err)
	}

	return db

}

// Close finish connection to database server
func Close(dbCon *sql.DB) {
	dbCon.Close()
}
