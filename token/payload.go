package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	Id        uuid.UUID     `json:"id"`
	Username  string        `json:"username"`
	UserId    bson.ObjectID `json:"user_id"`
	Role      string        `json:"role"`
	IssuedAt  time.Time     `json:"issued_at"`
	ExpiredAt time.Time     `json:"expired_at"`
}

func NewPayload(username string, userId bson.ObjectID, role string, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		Id:        id,
		Username:  username,
		UserId:    userId,
		Role:      role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
