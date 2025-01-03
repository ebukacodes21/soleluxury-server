// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: billboard.sql

package db

import (
	"context"
	"database/sql"
)

const createBillboard = `-- name: CreateBillboard :one
INSERT INTO billboards (
  store_id, label, image_url
) VALUES (
  $1, $2, $3
)
RETURNING id, store_id, label, image_url, created_at
`

type CreateBillboardParams struct {
	StoreID  int64  `db:"store_id" json:"store_id"`
	Label    string `db:"label" json:"label"`
	ImageUrl string `db:"image_url" json:"image_url"`
}

func (q *Queries) CreateBillboard(ctx context.Context, arg CreateBillboardParams) (Billboard, error) {
	row := q.db.QueryRowContext(ctx, createBillboard, arg.StoreID, arg.Label, arg.ImageUrl)
	var i Billboard
	err := row.Scan(
		&i.ID,
		&i.StoreID,
		&i.Label,
		&i.ImageUrl,
		&i.CreatedAt,
	)
	return i, err
}

const deleteBillboard = `-- name: DeleteBillboard :exec
DELETE FROM billboards
WHERE id = $1
`

func (q *Queries) DeleteBillboard(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBillboard, id)
	return err
}

const getBillboard = `-- name: GetBillboard :one
SELECT id, store_id, label, image_url, created_at FROM billboards
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetBillboard(ctx context.Context, id int64) (Billboard, error) {
	row := q.db.QueryRowContext(ctx, getBillboard, id)
	var i Billboard
	err := row.Scan(
		&i.ID,
		&i.StoreID,
		&i.Label,
		&i.ImageUrl,
		&i.CreatedAt,
	)
	return i, err
}

const getBillboards = `-- name: GetBillboards :many
SELECT id, store_id, label, image_url, created_at FROM billboards
WHERE store_id = $1
ORDER BY id
`

func (q *Queries) GetBillboards(ctx context.Context, storeID int64) ([]Billboard, error) {
	rows, err := q.db.QueryContext(ctx, getBillboards, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Billboard{}
	for rows.Next() {
		var i Billboard
		if err := rows.Scan(
			&i.ID,
			&i.StoreID,
			&i.Label,
			&i.ImageUrl,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBillboard = `-- name: UpdateBillboard :exec
UPDATE billboards
SET
  label = COALESCE($1, label),
  image_url = COALESCE($2, image_url)
WHERE 
  id = $3
  AND store_id = $4
`

type UpdateBillboardParams struct {
	Label    sql.NullString `db:"label" json:"label"`
	ImageUrl sql.NullString `db:"image_url" json:"image_url"`
	ID       int64          `db:"id" json:"id"`
	StoreID  int64          `db:"store_id" json:"store_id"`
}

func (q *Queries) UpdateBillboard(ctx context.Context, arg UpdateBillboardParams) error {
	_, err := q.db.ExecContext(ctx, updateBillboard,
		arg.Label,
		arg.ImageUrl,
		arg.ID,
		arg.StoreID,
	)
	return err
}
