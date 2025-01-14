package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type TokenContract interface {
	CreateToken(username string, userId bson.ObjectID, role string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}

type Token struct {
	engine *paseto.V2
	key    []byte
}

func NewToken(key string) (TokenContract, error) {
	if len(key) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size")
	}

	token := &Token{
		engine: paseto.NewV2(),
		key:    []byte(key),
	}

	return token, nil
}

func (t *Token) CreateToken(username string, userId bson.ObjectID, role string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, userId, role, duration)
	if err != nil {
		return "", nil, fmt.Errorf("unable to create payload")
	}

	token, err := t.engine.Encrypt(t.key, payload, nil)
	return token, payload, err
}

func (t *Token) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := t.engine.Decrypt(token, t.key, &payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, err
}
