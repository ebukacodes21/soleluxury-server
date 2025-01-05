// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Billboard struct {
	ID        int64     `db:"id" json:"id"`
	StoreID   int64     `db:"store_id" json:"store_id"`
	Label     string    `db:"label" json:"label"`
	ImageUrl  string    `db:"image_url" json:"image_url"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Category struct {
	ID             int64     `db:"id" json:"id"`
	StoreID        int64     `db:"store_id" json:"store_id"`
	BillboardID    int64     `db:"billboard_id" json:"billboard_id"`
	StoreName      string    `db:"store_name" json:"store_name"`
	BillboardLabel string    `db:"billboard_label" json:"billboard_label"`
	Name           string    `db:"name" json:"name"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

type Color struct {
	ID        int64     `db:"id" json:"id"`
	StoreID   int64     `db:"store_id" json:"store_id"`
	StoreName string    `db:"store_name" json:"store_name"`
	Name      string    `db:"name" json:"name"`
	Value     string    `db:"value" json:"value"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Order struct {
	ID        int64           `db:"id" json:"id"`
	StoreID   int64           `db:"store_id" json:"store_id"`
	Items     json.RawMessage `db:"items" json:"items"`
	IsPaid    bool            `db:"is_paid" json:"is_paid"`
	Phone     string          `db:"phone" json:"phone"`
	Address   string          `db:"address" json:"address"`
	CreatedAt time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt time.Time       `db:"updated_at" json:"updated_at"`
}

type Product struct {
	ID          int64           `db:"id" json:"id"`
	Name        string          `db:"name" json:"name"`
	Price       float64         `db:"price" json:"price"`
	IsFeatured  bool            `db:"is_featured" json:"is_featured"`
	IsArchived  bool            `db:"is_archived" json:"is_archived"`
	Description string          `db:"description" json:"description"`
	Images      json.RawMessage `db:"images" json:"images"`
	CreatedAt   time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time       `db:"updated_at" json:"updated_at"`
}

type ProductCategory struct {
	ProductID  int64 `db:"product_id" json:"product_id"`
	CategoryID int64 `db:"category_id" json:"category_id"`
}

type ProductColor struct {
	ProductID int64 `db:"product_id" json:"product_id"`
	ColorID   int64 `db:"color_id" json:"color_id"`
}

type ProductSize struct {
	ProductID int64 `db:"product_id" json:"product_id"`
	SizeID    int64 `db:"size_id" json:"size_id"`
}

type ProductStore struct {
	ProductID int64 `db:"product_id" json:"product_id"`
	StoreID   int64 `db:"store_id" json:"store_id"`
}

type Session struct {
	ID           uuid.UUID `db:"id" json:"id"`
	UserID       int64     `db:"user_id" json:"user_id"`
	Username     string    `db:"username" json:"username"`
	RefreshToken string    `db:"refresh_token" json:"refresh_token"`
	UserAgent    string    `db:"user_agent" json:"user_agent"`
	ClientIp     string    `db:"client_ip" json:"client_ip"`
	IsBlocked    bool      `db:"is_blocked" json:"is_blocked"`
	ExpiredAt    time.Time `db:"expired_at" json:"expired_at"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

type Size struct {
	ID        int64     `db:"id" json:"id"`
	StoreID   int64     `db:"store_id" json:"store_id"`
	StoreName string    `db:"store_name" json:"store_name"`
	Name      string    `db:"name" json:"name"`
	Value     string    `db:"value" json:"value"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Store struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type User struct {
	ID               int64     `db:"id" json:"id"`
	Username         string    `db:"username" json:"username"`
	Password         string    `db:"password" json:"password"`
	Email            string    `db:"email" json:"email"`
	IsVerified       bool      `db:"is_verified" json:"is_verified"`
	VerificationCode string    `db:"verification_code" json:"verification_code"`
	Role             string    `db:"role" json:"role"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
}
