// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: session.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (
  id, username, user_id,refresh_token, user_agent, client_ip, is_blocked, expired_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING id, user_id, username, refresh_token, user_agent, client_ip, is_blocked, expired_at, created_at
`

type CreateSessionParams struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Username     string    `db:"username" json:"username"`
	UserID       int64     `db:"user_id" json:"user_id"`
	RefreshToken string    `db:"refresh_token" json:"refresh_token"`
	UserAgent    string    `db:"user_agent" json:"user_agent"`
	ClientIp     string    `db:"client_ip" json:"client_ip"`
	IsBlocked    bool      `db:"is_blocked" json:"is_blocked"`
	ExpiredAt    time.Time `db:"expired_at" json:"expired_at"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession,
		arg.ID,
		arg.Username,
		arg.UserID,
		arg.RefreshToken,
		arg.UserAgent,
		arg.ClientIp,
		arg.IsBlocked,
		arg.ExpiredAt,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Username,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiredAt,
		&i.CreatedAt,
	)
	return i, err
}

const logout = `-- name: Logout :exec
DELETE FROM sessions
WHERE id = $1
`

func (q *Queries) Logout(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, logout, id)
	return err
}