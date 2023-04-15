package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	//
}

type DBStore struct {
	db *sql.DB
}

func InitDB(dbFileName string) (*DBStore, error) {
	db, err := sql.Open("sqlite3", dbFileName)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return &DBStore{
		db: db,
	}, nil
}
