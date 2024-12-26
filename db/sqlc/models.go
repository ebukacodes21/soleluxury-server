// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"time"
)

type Store struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type User struct {
	ID               int64          `db:"id" json:"id"`
	FirstName        string         `db:"first_name" json:"first_name"`
	LastName         string         `db:"last_name" json:"last_name"`
	Username         string         `db:"username" json:"username"`
	Password         string         `db:"password" json:"password"`
	Email            string         `db:"email" json:"email"`
	IsVerified       bool           `db:"is_verified" json:"is_verified"`
	VerificationCode string         `db:"verification_code" json:"verification_code"`
	Country          string         `db:"country" json:"country"`
	Phone            string         `db:"phone" json:"phone"`
	Role             string         `db:"role" json:"role"`
	ProfilePic       sql.NullString `db:"profile_pic" json:"profile_pic"`
	CreatedAt        time.Time      `db:"created_at" json:"created_at"`
}
