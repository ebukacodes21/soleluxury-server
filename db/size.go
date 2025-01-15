package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// create size
func (r *Repository) CreateSize(ctx context.Context, args *pb.CreateSizeRequest) (*Size, error) {
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

	size := &Size{
		StoreID:   store.ID,
		Name:      name,
		Value:     args.GetValue(),
		CreatedAt: time.Now(),
	}

	result, err := r.sizeColl.InsertOne(ctx, size)
	if err != nil {
		return nil, fmt.Errorf("unable to create size: %v", err)
	}

	size.ID = result.InsertedID.(bson.ObjectID)
	return size, nil
}

// get size by id
func (r *Repository) GetSize(ctx context.Context, args *pb.GetSizeRequest) (*Size, error) {
	var size Size
	sID, err := bson.ObjectIDFromHex(args.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid size id: %v", err)
	}

	err = r.sizeColl.FindOne(ctx, bson.M{"_id": sID}).Decode(&size)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("size with ID %s not found", sID)
		}
		return nil, fmt.Errorf("unable to fetch size by ID: %v", err)
	}

	err = r.populateSizeStore(ctx, &size)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch associated details: %v", err)
	}

	return &size, nil
}

// get sizes
func (r *Repository) GetAllSizes(ctx context.Context, args *pb.GetSizesRequest) ([]Size, error) {
	var sizes []Size
	sId, err := bson.ObjectIDFromHex(args.GetStoreId())
	if err != nil {
		return nil, fmt.Errorf("invalid store id for size")
	}

	cursor, err := r.sizeColl.Find(ctx, bson.M{"store_id": sId})
	if err != nil {
		return nil, fmt.Errorf("unable to fetch sizes for store: %v", err)
	}
	defer cursor.Close(ctx)

	// decode all the sizes
	for cursor.Next(ctx) {
		var size Size
		if err := cursor.Decode(&size); err != nil {
			return nil, fmt.Errorf("unable to decode size: %v", err)
		}

		if err := r.populateSizeStore(ctx, &size); err != nil {
			return nil, fmt.Errorf("unable to populate size store: %v", err)
		}

		sizes = append(sizes, size)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return sizes, nil
}

// update size with id
func (r *Repository) UpdateSize(ctx context.Context, args *pb.UpdateSizeRequest) (string, error) {
	var size Size
	id, err := bson.ObjectIDFromHex(args.GetId())
	if err != nil {
		return "", fmt.Errorf("invalid size id: %v", err)
	}

	err = r.sizeColl.FindOne(ctx, bson.M{"_id": id}).Decode(&size)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("size with ID %s not found", id)
		}
		return "", fmt.Errorf("unable to fetch size: %v", err)
	}

	update := bson.M{
		"$set": bson.M{
			"name":       args.GetName(),
			"value":      args.GetValue(),
			"updated_at": time.Now(),
		},
	}

	_, err = r.sizeColl.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return "", fmt.Errorf("unable to update size: %v", err)
	}

	return "size update successful", nil
}

// delete size with id
func (r *Repository) DeleteSize(ctx context.Context, arg *pb.DeleteSizeRequest) (string, error) {
	sId, err := bson.ObjectIDFromHex(arg.GetId())
	if err != nil {
		return "", fmt.Errorf("invalid size id: %v", err)
	}

	_, err = r.sizeColl.DeleteOne(ctx, bson.M{"_id": sId})
	if err != nil {
		return "", fmt.Errorf("unable to delete size: %v", err)
	}

	return "size delete successful", nil
}

func (r *Repository) populateSizeStore(ctx context.Context, size *Size) error {
	var store Store
	err := r.storeColl.FindOne(ctx, bson.M{"_id": size.StoreID}).Decode(&store)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("size for store %s not found: %v", size.StoreID.Hex(), err)
		}
		return fmt.Errorf("unable to fetch size for store %s: %v", size.StoreID.Hex(), err)
	}

	size.Store = store
	return nil
}
