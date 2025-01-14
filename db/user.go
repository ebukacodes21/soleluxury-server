package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/utils"
	"github.com/ebukacodes21/soleluxury-server/validate"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// create user
func (r *Repository) CreateUser(ctx context.Context, args *pb.CreateUserRequest) (*User, error) {
	email := args.GetEmail()
	var existingUser User

	err := r.userColl.FindOne(ctx, bson.M{"email": email}).Decode(&existingUser)
	if err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("email %s is already registered", email)
	}

	code := validate.RandomString(32)
	hash, err := utils.HashPassword(args.GetPassword())
	if err != nil {
		return nil, fmt.Errorf("unable to hash password: %v", err)
	}

	user := &User{
		Username:         args.GetUsername(),
		Password:         hash,
		Email:            email,
		Role:             "admin",
		VerificationCode: code,
		IsVerified:       false,
		CreatedAt:        time.Now(),
	}

	result, err := r.userColl.InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("unable to create user: %v", err)
	}

	user.ID = result.InsertedID.(bson.ObjectID)
	return user, nil
}

// find user by email
func (r *Repository) FindUser(ctx context.Context, args *pb.LoginUserRequest) (*User, error) {
	email := args.GetEmail()
	var user User

	err := r.userColl.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("user with email %s not found", email)
	} else if err != nil {
		return nil, fmt.Errorf("an error occurred %s ", err)
	}

	err = utils.ComparePassword(args.GetPassword(), user.Password)
	if err != nil {
		return nil, fmt.Errorf("incorrect password %s ", err)
	}

	return &user, nil
}
