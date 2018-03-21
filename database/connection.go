package database

import (
	"database/sql"
	"path"

	"github.com/lhdv/treasure_hunt/util"

	_ "github.com/mattn/go-sqlite3" // SQLite3 Driver
)

// DBFile define a default database name(or file)
const dbFile = "treasure_hunt.dat"

// Open connects to database server
func Open(dbConn string) (*sql.DB, error) {

	var db *sql.DB
	var err error

	if len(dbConn) <= 0 {
		dbConn = dbFile
	}

	if path.Ext(dbConn) == ".dat" {
		db, err = sql.Open("sqlite3", dbConn)
	} else {
		//db, err = sql.Open("sqlite3", dbConn) // other DB
	}

	if err != nil {
		util.LogError("Open Connection", err)
		return nil, err
	}

	return db, nil
}

// Close finish connection to database server
func Close(dbCon *sql.DB) {
	dbCon.Close()
}
