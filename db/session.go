package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SessionReq struct {
	Username     string        `bson:"username,omitempty"`
	UserID       bson.ObjectID `bson:"user_id,omitempty"`
	RefreshToken string        `bson:"refresh_token,omitempty"`
	UserAgent    string        `bson:"user_agent,omitempty"`
	ClientIp     string        `bson:"client_ip,omitempty"`
	IsBlocked    bool          `bson:"is_blocked,omitempty"`
	ExpiredAt    *time.Time    `bson:"expired_at,omitempty"`
}

// create session for user
func (r *Repository) CreateSession(ctx context.Context, args SessionReq) (*Session, error) {
	user_id := args.UserID
	var prevSession Session

	err := r.sessionColl.FindOne(ctx, bson.M{"user_id": user_id}).Decode(&prevSession)
	if err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("user %s is already logged in", user_id)
	}

	session := &Session{
		Username:     args.Username,
		UserID:       args.UserID,
		RefreshToken: args.RefreshToken,
		UserAgent:    args.UserAgent,
		ClientIp:     args.ClientIp,
		IsBlocked:    args.IsBlocked,
		ExpiredAt:    args.ExpiredAt,
	}

	result, err := r.sessionColl.InsertOne(ctx, session)
	if err != nil {
		return nil, fmt.Errorf("unable to create session: %v", err)
	}

	session.ID = result.InsertedID.(bson.ObjectID)
	return session, nil
}

// delete session with user_id
func (r *Repository) LogOut(ctx context.Context, id bson.ObjectID) error {
	_, err := r.sessionColl.DeleteOne(ctx, bson.M{"user_id": id})
	if err != nil {
		return fmt.Errorf("unable to delete session: %v", err)
	}

	return nil
}
