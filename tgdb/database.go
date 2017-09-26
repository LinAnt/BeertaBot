package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // required for sqlite3 database
)

func openDB() (*sql.DB, error) {
	// TODO: If the file does not exist, create the original table structure here?
	return sql.Open("sqlite3", databaseFilePath)
}

// Database - main struct for all database actions
type Database struct {
	path string
	db   *sql.DB
}

func (db *Database) setPath(path string) error {
	db.path = path
}

func (db *Database) openOrCreate() (*sql.DB, error) {

	_, err := sql.Open("sqlite3", db.path)
	if err != nil {
		//TODO create a new database
		log.infof()
	}

	return sql.Open("sqlite3", db.path)
}

// NewDatabase - returns an operational database instance
func NewDatabase(path string) (Database, error) {
	var db Database

	err := db.setPath(path)
	if err != nil {
		return Database{}, nil
	}
	return Database{}, nil
}
