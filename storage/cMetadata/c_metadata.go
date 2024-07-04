package cMetadata

import "database/sql"

type CMetadata struct {
	DB *sql.DB
}

func NewCMetadata(db *sql.DB) *CMetadata {
	return &CMetadata{DB: db}
}
