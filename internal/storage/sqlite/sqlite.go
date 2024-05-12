package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"link_shortener/lib/e"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, e.Wrap("storage.sqlite.New", err)
	}

	if err = initDb(db); err != nil {
		return nil, e.Wrap("init db error", err)
	}

	return &Storage{db: db}, nil
}

func initDb(db *sql.DB) error {
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS url (
		id INTEGER PRIMARY KEY,
		alias TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL	    
	);
	CREATE INDEX IS NOT EXISTS idx_alias ON url(alias);
	`)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(); err != nil {
		return err
	}

	return nil
}
