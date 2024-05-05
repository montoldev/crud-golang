package database

import (
	"context"
	"database/sql"
)

type IDatabase interface {
	GetMany(ctx context.Context) error
}

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("sqlite3", "./crud.db")
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

func (db *Database) Close() error {
	return db.db.Close()
}
