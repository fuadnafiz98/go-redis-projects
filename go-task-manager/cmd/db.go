package main

import (
	"database/sql"
)

type DBStore struct {
	dbFileName string
	db         *sql.DB
}

func (d *DBStore) Init() error {
	db, err := sql.Open("sqlite3", d.dbFileName)
	d.db = db

	return err
}
