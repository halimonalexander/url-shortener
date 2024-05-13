package storage

import "errors"

var (
	ErrURLNotFound           = errors.New("url not found")
	ErrURLExists             = errors.New("url exists")
	ErrUnableToFetchRecordId = errors.New("unable to fetch id for a new record")
)
