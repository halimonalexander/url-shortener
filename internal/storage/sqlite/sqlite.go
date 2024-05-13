package sqlite

import (
	"database/sql"
	"github.com/mattn/go-sqlite3"
	"link_shortener/internal/storage"
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
	CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	`)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Storage) SaveUrl(incomingUrl string, alias string) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO url (url, alias) VALUES (?, ?)")
	if err != nil {
		return 0, e.Wrap("storage.sqlite.SaveUrl", err)
	}

	res, err := stmt.Exec(incomingUrl, alias)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, e.Wrap("storage.sqlite.SaveUrl", storage.ErrURLExists)
		}

		return 0, e.Wrap("storage.sqlite.SaveUrl", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, e.Wrap("unable to fetch id", storage.ErrUnableToFetchRecordId)
	}

	return id, nil
}

func (s Storage) GetUrl(alias string) (string, error) {
	var requestedUrl string

	if err := s.db.QueryRow("SELECT url FROM url WHERE alias = ? LIMIT 1", alias).Scan(&requestedUrl); err != nil {
		return "", e.Wrap("storage.sqlite.GetUrl", e.IfIsChangeTo(err, sql.ErrNoRows, storage.ErrURLNotFound))
	}

	return requestedUrl, nil
}

func (s Storage) RemoveUrl(alias string) error {
	stmt, err := s.db.Prepare("DELETE FROM url WHERE alias = ?")
	if err != nil {
		return e.Wrap("storage.sqlite.RemoveUrl", err)
	}

	if _, err := stmt.Exec(alias); err != nil {
		return e.Wrap("storage.sqlite.RemoveUrl", err)
	}

	return nil
}
