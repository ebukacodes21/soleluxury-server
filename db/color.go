package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// create color
func (r *Repository) CreateColor(ctx context.Context, args *pb.CreateColorRequest) (*Color, error) {
	var store Store
	name := args.GetName()

	storeId, err := bson.ObjectIDFromHex(args.GetStoreId())
	if err != nil {
		return nil, fmt.Errorf("invalid store id")
	}

	err = r.storeColl.FindOne(ctx, bson.M{"_id": storeId}).Decode(&store)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("store with id %s not found", storeId)
	}

	color := &Color{
		StoreID:   store.ID,
		Name:      name,
		Value:     args.GetValue(),
		CreatedAt: time.Now(),
	}

	result, err := r.colorColl.InsertOne(ctx, color)
	if err != nil {
		return nil, fmt.Errorf("unable to create color: %v", err)
	}

	color.ID = result.InsertedID.(bson.ObjectID)
	return color, nil
}

// get color by id
func (r *Repository) GetColor(ctx context.Context, args *pb.GetColorRequest) (*Color, error) {
	var color Color
	cId, err := bson.ObjectIDFromHex(args.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid color id: %v", err)
	}

	err = r.colorColl.FindOne(ctx, bson.M{"_id": cId}).Decode(&color)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("color with ID %s not found", cId)
		}
		return nil, fmt.Errorf("unable to fetch color by ID: %v", err)
	}

	err = r.populateColorStore(ctx, &color)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch associated details: %v", err)
	}

	return &color, nil
}

// get colors
func (r *Repository) GetAllColors(ctx context.Context, args *pb.GetColorsRequest) ([]Color, error) {
	var colors []Color
	sId, err := bson.ObjectIDFromHex(args.GetStoreId())
	if err != nil {
		return nil, fmt.Errorf("invalid store id for color")
	}

	cursor, err := r.colorColl.Find(ctx, bson.M{"store_id": sId})
	if err != nil {
		return nil, fmt.Errorf("unable to fetch colors for store: %v", err)
	}
	defer cursor.Close(ctx)

	// decode all the colors
	for cursor.Next(ctx) {
		var color Color
		if err := cursor.Decode(&color); err != nil {
			return nil, fmt.Errorf("unable to decode color: %v", err)
		}

		if err := r.populateColorStore(ctx, &color); err != nil {
			return nil, fmt.Errorf("unable to populate color store: %v", err)
		}

		colors = append(colors, color)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return colors, nil
}

// update color with id
func (r *Repository) UpdateColor(ctx context.Context, args *pb.UpdateColorRequest) (string, error) {
	var color Color
	id, err := bson.ObjectIDFromHex(args.GetId())
	if err != nil {
		return "", fmt.Errorf("invalid color id: %v", err)
	}

	err = r.colorColl.FindOne(ctx, bson.M{"_id": id}).Decode(&color)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("color with ID %s not found", id)
		}
		return "", fmt.Errorf("unable to fetch color: %v", err)
	}

	update := bson.M{
		"$set": bson.M{
			"name":       args.GetName(),
			"value":      args.GetValue(),
			"updated_at": time.Now(),
		},
	}

	_, err = r.colorColl.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return "", fmt.Errorf("unable to update color: %v", err)
	}

	return "color update successful", nil
}

// delete color with id
func (r *Repository) DeleteColor(ctx context.Context, arg *pb.DeleteColorRequest) (string, error) {
	id, err := bson.ObjectIDFromHex(arg.GetId())
	if err != nil {
		return "", fmt.Errorf("invalid color id: %v", err)
	}

	_, err = r.colorColl.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return "", fmt.Errorf("unable to delete color: %v", err)
	}

	return "color delete successful", nil
}

func (r *Repository) populateColorStore(ctx context.Context, color *Color) error {
	var store Store
	err := r.storeColl.FindOne(ctx, bson.M{"_id": color.StoreID}).Decode(&store)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("color for store %s not found: %v", color.StoreID.Hex(), err)
		}
		return fmt.Errorf("unable to fetch color for store %s: %v", color.StoreID.Hex(), err)
	}

	color.Store = store
	return nil
}
