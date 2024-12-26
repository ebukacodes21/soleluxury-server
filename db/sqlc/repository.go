package db

import "database/sql"

type DatabaseContract interface {
	Querier
}

type SoleluxuryRepository struct {
	*Queries
	db *sql.DB
}

func NewSoleluxuryRepository(db *sql.DB) DatabaseContract {
	return &SoleluxuryRepository{db: db, Queries: New(db)}
}
