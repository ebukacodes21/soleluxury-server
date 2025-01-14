package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// create billboard
func (r *Repository) CreateBillboard(ctx context.Context, args *pb.CreateBillboardRequest) (*Billboard, error) {
	var exlabel Billboard
	var store Store
	label := args.GetLabel()

	storeId, err := bson.ObjectIDFromHex(args.GetStoreId())
	if err != nil {
		return nil, fmt.Errorf("invalid store id")
	}

	err = r.storeColl.FindOne(ctx, bson.M{"_id": storeId}).Decode(&store)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("store with id %s not found", storeId)
	}

	err = r.billboardColl.FindOne(ctx, bson.M{"label": label}).Decode(&exlabel)
	if err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("billboard %s already existing", label)
	}

	billboard := &Billboard{
		StoreID:   store.ID,
		Label:     label,
		ImageURL:  args.GetImageUrl(),
		CreatedAt: time.Now(),
	}

	result, err := r.billboardColl.InsertOne(ctx, billboard)
	if err != nil {
		return nil, fmt.Errorf("unable to create billboard: %v", err)
	}

	billboard.ID = result.InsertedID.(bson.ObjectID)
	return billboard, nil
}

// get billboard
func (r *Repository) GetBillboard(ctx context.Context, args *pb.GetBillboardRequest) (*Billboard, error) {
	var billboard Billboard
	id, err := bson.ObjectIDFromHex(args.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid billboard id")
	}

	err = r.billboardColl.FindOne(ctx, bson.M{"_id": id}).Decode(&billboard)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("billboard with ID %s not found", id)
		}
		return nil, fmt.Errorf("unable to fetch billboard by ID: %v", err)
	}

	if err := r.populateBillboardStore(ctx, &billboard); err != nil {
		return nil, fmt.Errorf("unable to populate billboard store: %v", err)
	}

	return &billboard, nil
}

// get billboards
func (r *Repository) GetAllBillboards(ctx context.Context, args *pb.GetBillboardsRequest) ([]Billboard, error) {
	var billboards []Billboard
	sId, err := bson.ObjectIDFromHex(args.GetStoreId())
	if err != nil {
		return nil, fmt.Errorf("invalid store id for billboard")
	}

	cursor, err := r.billboardColl.Find(ctx, bson.M{"store_id": sId})
	if err != nil {
		return nil, fmt.Errorf("unable to fetch billboards for store: %v", err)
	}
	defer cursor.Close(ctx)

	// decode all the billboards
	for cursor.Next(ctx) {
		var billboard Billboard
		if err := cursor.Decode(&billboard); err != nil {
			return nil, fmt.Errorf("unable to decode billboard: %v", err)
		}

		if err := r.populateBillboardCategories(ctx, &billboard); err != nil {
			return nil, fmt.Errorf("unable to populate billboard categories: %v", err)
		}

		if err := r.populateBillboardStore(ctx, &billboard); err != nil {
			return nil, fmt.Errorf("unable to populate billboard store: %v", err)
		}

		billboards = append(billboards, billboard)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return billboards, nil
}

// update billboard with id
func (r *Repository) UpdateBillboard(ctx context.Context, args *pb.UpdateBillboardRequest) (string, error) {
	var billboard Billboard
	id, err := bson.ObjectIDFromHex(args.GetId())
	if err != nil {
		return "", fmt.Errorf("invalid billboard id: %v", err)
	}

	err = r.billboardColl.FindOne(ctx, bson.M{"_id": id}).Decode(&billboard)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("billboard with ID %s not found", id)
		}
		return "", fmt.Errorf("unable to fetch billboard: %v", err)
	}

	update := bson.M{
		"$set": bson.M{
			"label":      args.GetLabel(),
			"image_url":  args.GetImageUrl(),
			"updated_at": time.Now(),
		},
	}

	_, err = r.billboardColl.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return "", fmt.Errorf("unable to update billboard: %v", err)
	}

	return "billboard update successful", nil
}

// delete billboard with id
func (r *Repository) DeleteBillboard(ctx context.Context, arg *pb.DeleteBillboardRequest) (string, error) {
	bId, err := bson.ObjectIDFromHex(arg.GetId())
	if err != nil {
		return "", fmt.Errorf("invalid billboard id: %v", err)
	}

	// start a session
	session, err := r.client.StartSession()
	if err != nil {
		return "", fmt.Errorf("unable to start session: %v", err)
	}
	defer session.EndSession(ctx)

	// start a tx
	err = session.StartTransaction()
	if err != nil {
		return "", fmt.Errorf("unable to start transaction: %v", err)
	}

	// clean up
	defer func() {
		if err != nil {
			session.AbortTransaction(ctx)
		}
	}()

	_, err = r.categoryColl.DeleteMany(ctx, bson.M{"billboard_id": bId})
	if err != nil {
		return "", fmt.Errorf("unable to delete categories for billboard %s: %v", bId.Hex(), err)
	}

	_, err = r.billboardColl.DeleteOne(ctx, bson.M{"_id": bId})
	if err != nil {
		return "", fmt.Errorf("unable to delete billboard: %v", err)
	}

	err = session.CommitTransaction(ctx)
	if err != nil {
		return "", fmt.Errorf("unable to commit transaction: %v", err)
	}

	return "billboard delete successful", nil
}

func (r *Repository) populateBillboardCategories(ctx context.Context, billboard *Billboard) error {
	var categories []Category
	cursor, err := r.categoryColl.Find(ctx, bson.M{"billboard_id": billboard.ID})
	if err != nil {
		return fmt.Errorf("unable to fetch categories for billboard %s: %v", billboard.ID.Hex(), err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var category Category
		if err := cursor.Decode(&category); err != nil {
			return fmt.Errorf("unable to decode category: %v", err)
		}
		categories = append(categories, category)
	}

	if err := cursor.Err(); err != nil {
		return fmt.Errorf("cursor error: %v", err)
	}

	billboard.Categories = categories
	return nil
}

func (r *Repository) populateBillboardStore(ctx context.Context, billboard *Billboard) error {
	var store Store
	err := r.storeColl.FindOne(ctx, bson.M{"_id": billboard.StoreID}).Decode(&store)
	if err != nil {
		return fmt.Errorf("unable to fetch store for billboard %s: %v", billboard.StoreID.Hex(), err)
	}

	billboard.Store = store
	return nil
}
