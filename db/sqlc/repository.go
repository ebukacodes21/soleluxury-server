package db

import (
	"context"
	"database/sql"
	"log"
)

type DatabaseContract interface {
	Querier
	CreateUserTx(ctx context.Context, args CreateUserTxParams) (CreateUserTxResponse, error)
}

type SoleluxuryRepository struct {
	*Queries
	db *sql.DB
}

func NewSoleluxuryRepository(db *sql.DB) DatabaseContract {
	return &SoleluxuryRepository{db: db, Queries: New(db)}
}

func (sr *SoleluxuryRepository) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := sr.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit()
}
