package database

import (
	"database/sql"

	"github.com/lhdv/treasure_hunt/util"

	_ "github.com/mattn/go-sqlite3" // SQLite3 Driver
)

// Open connects to database server
func Open() *sql.DB {

	db, err := sql.Open("sqlite3", DBFile)
	if err != nil {
		util.LogError("Open Connection", err)
	}

	return db

}

// Close finish connection to database server
func Close(dbCon *sql.DB) {
	dbCon.Close()
}
