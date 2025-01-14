package db

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Store struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	Name       string        `bson:"name,omitempty"`
	CreatedAt  time.Time     `bson:"created_at,omitempty"`
	UpdatedAt  time.Time     `bson:"updated_at,omitempty"`
	Products   []Product     `bson:"products,omitempty"`
	Billboards []Billboard   `bson:"billboards,omitempty"`
	Categories []Category    `bson:"categories,omitempty"`
	Colors     []Color       `bson:"colors,omitempty"`
	Sizes      []Size        `bson:"sizes,omitempty"`
	Orders     []Order       `bson:"orders,omitempty"`
}

type User struct {
	ID               bson.ObjectID `bson:"_id,omitempty"`
	Username         string        `bson:"username,omitempty"`
	Email            string        `bson:"email,omitempty"`
	Password         string        `bson:"password,omitempty"`
	VerificationCode string        `bson:"verification_code,omitempty"`
	IsVerified       bool          `bson:"is_verified,omitempty"`
	Role             string        `bson:"role,omitempty"`
	CreatedAt        time.Time     `bson:"created_at,omitempty"`
}

type Session struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Username     string        `bson:"username,omitempty"`
	UserID       bson.ObjectID `bson:"user_id,omitempty"`
	RefreshToken string        `bson:"refresh_token,omitempty"`
	UserAgent    string        `bson:"user_agent,omitempty"`
	ClientIp     string        `bson:"client_ip,omitempty"`
	IsBlocked    bool          `bson:"is_blocked,omitempty"`
	ExpiredAt    *time.Time    `bson:"expired_at,omitempty"`
}

type Billboard struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	StoreID    bson.ObjectID `bson:"store_id,omitempty"`
	Store      Store         `bson:"store,omitempty"`
	Label      string        `bson:"label,omitempty"`
	ImageURL   string        `bson:"image_url,omitempty"`
	Categories []Category    `bson:"categories,omitempty"`
	CreatedAt  time.Time     `bson:"created_at,omitempty"`
	UpdatedAt  time.Time     `bson:"updated_at,omitempty"`
}

type Category struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	StoreID     bson.ObjectID `bson:"store_id"`
	BillboardID bson.ObjectID `bson:"billboard_id"`
	Billboard   Billboard     `bson:"billboard,omitempty"`
	Products    []Product     `bson:"products,omitempty"`
	Store       Store         `bson:"store,omitempty"`
	Name        string        `bson:"name,omitempty"`
	CreatedAt   time.Time     `bson:"created_at,omitempty"`
	UpdatedAt   time.Time     `bson:"updated_at,omitempty"`
}

type Size struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	StoreID   bson.ObjectID `bson:"store_id,omitempty"`
	Store     Store         `bson:"store,omitempty"`
	Products  []Product     `bson:"products,omitempty"`
	Name      string        `bson:"name,omitempty"`
	Value     string        `bson:"value,omitempty"`
	CreatedAt time.Time     `bson:"created_at,omitempty"`
	UpdatedAt time.Time     `bson:"updated_at,omitempty"`
}

type Color struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	StoreID   bson.ObjectID `bson:"store_id,omitempty"`
	Store     Store         `bson:"store,omitempty"`
	Products  []Product     `bson:"products,omitempty"`
	Name      string        `bson:"name,omitempty"`
	Value     string        `bson:"value,omitempty"`
	CreatedAt time.Time     `bson:"created_at,omitempty"`
	UpdatedAt time.Time     `bson:"updated_at,omitempty"`
}

type Image struct {
	URL string `json:"url"`
}

type Product struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	StoreID     bson.ObjectID `bson:"store_id,omitempty"`
	Store       Store         `bson:"store,omitempty"`
	CategoryID  bson.ObjectID `bson:"category_id,omitempty"`
	Category    Category      `bson:"category,omitempty"`
	Name        string        `bson:"name,omitempty"`
	Price       float64       `bson:"price,omitempty"`
	IsFeatured  bool          `bson:"is_featured,omitempty"`
	IsArchived  bool          `bson:"is_archived,omitempty"`
	SizeID      bson.ObjectID `bson:"size_id,omitempty"`
	Size        Size          `bson:"size,omitempty"`
	ColorID     bson.ObjectID `bson:"color_id,omitempty"`
	Color       Color         `bson:"color,omitempty"`
	Images      []Image       `bson:"images,omitempty"`
	Description string        `bson:"description,omitempty"`
	CreatedAt   time.Time     `bson:"created_at,omitempty"`
	UpdatedAt   time.Time     `bson:"updated_at,omitempty"`
}

type Order struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	StoreID   bson.ObjectID `bson:"store_id"`
	Items     []bson.M      `bson:"items,omitempty"`
	IsPaid    bool          `bson:"is_paid,omitempty"`
	Phone     string        `bson:"phone,omitempty"`
	Address   string        `bson:"address,omitempty"`
	CreatedAt time.Time     `bson:"created_at,omitempty"`
	UpdatedAt time.Time     `bson:"updated_at,omitempty"`
}
