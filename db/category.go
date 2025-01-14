package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// create category
func (r *Repository) CreateCategory(ctx context.Context, args *pb.CreateCategoryRequest) (*Category, error) {
	var exCategory Category
	var store Store
	var billboard Billboard
	name := args.GetName()

	err := r.categoryColl.FindOne(ctx, bson.M{"name": name}).Decode(&exCategory)
	if err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("category %s already existing", name)
	}

	storeId, err := bson.ObjectIDFromHex(args.GetStoreId())
	if err != nil {
		return nil, fmt.Errorf("invalid store id")
	}

	billId, err := bson.ObjectIDFromHex(args.GetBillboardId())
	if err != nil {
		return nil, fmt.Errorf("invalid billboard id")
	}

	err = r.storeColl.FindOne(ctx, bson.M{"_id": storeId}).Decode(&store)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("store with id %s not found", storeId)
	}

	err = r.billboardColl.FindOne(ctx, bson.M{"_id": billId}).Decode(&billboard)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("billboard with id %s not found", billId)
	}

	category := &Category{
		StoreID:     store.ID,
		BillboardID: billboard.ID,
		Name:        name,
		CreatedAt:   time.Now(),
	}

	result, err := r.categoryColl.InsertOne(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("unable to create category: %v", err)
	}

	category.ID = result.InsertedID.(bson.ObjectID)
	return category, nil
}

// get category by id
func (r *Repository) GetCatgeoryByID(ctx context.Context, req *pb.GetCategoryRequest) (*Category, error) {
	var category Category
	sID, err := bson.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid category id: %v", err)
	}

	err = r.categoryColl.FindOne(ctx, bson.M{"_id": sID}).Decode(&category)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("category with ID %s not found", sID)
		}
		return nil, fmt.Errorf("unable to fetch category by ID: %v", err)
	}

	err = r.populateCategoryBillboard(ctx, &category)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch associated details: %v", err)
	}

	return &category, nil
}

// get categories
func (r *Repository) GetAllCategories(ctx context.Context, args *pb.GetCategoriesRequest) ([]Category, error) {
	var categories []Category
	sId, err := bson.ObjectIDFromHex(args.GetStoreId())
	if err != nil {
		return nil, fmt.Errorf("invalid store id for category")
	}

	cursor, err := r.categoryColl.Find(ctx, bson.M{"store_id": sId})
	if err != nil {
		return nil, fmt.Errorf("unable to fetch categories for store: %v", err)
	}
	defer cursor.Close(ctx)

	// decode all the categories
	for cursor.Next(ctx) {
		var category Category
		if err := cursor.Decode(&category); err != nil {
			return nil, fmt.Errorf("unable to decode category: %v", err)
		}

		if err := r.populateCategoryBillboard(ctx, &category); err != nil {
			return nil, fmt.Errorf("unable to populate category billboard: %v", err)
		}

		categories = append(categories, category)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return categories, nil
}

// update category with id
func (r *Repository) UpdateCategory(ctx context.Context, args *pb.UpdateCategoryRequest) (string, error) {
	var category Category
	id, err := bson.ObjectIDFromHex(args.GetId())
	if err != nil {
		return "", fmt.Errorf("invalid category id: %v", err)
	}

	err = r.categoryColl.FindOne(ctx, bson.M{"_id": id}).Decode(&category)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("category with ID %s not found", id)
		}
		return "", fmt.Errorf("unable to fetch category: %v", err)
	}

	update := bson.M{
		"$set": bson.M{
			"name":         args.GetName(),
			"billboard_id": args.GetBillboardId(),
			"updated_at":   time.Now(),
		},
	}

	_, err = r.categoryColl.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return "", fmt.Errorf("unable to update category: %v", err)
	}

	return "category update successful", nil
}

// delete category with id
func (r *Repository) DeleteCategory(ctx context.Context, arg *pb.DeleteCategoryRequest) (string, error) {
	cId, err := bson.ObjectIDFromHex(arg.GetId())
	if err != nil {
		return "", fmt.Errorf("invalid category id: %v", err)
	}

	_, err = r.categoryColl.DeleteOne(ctx, bson.M{"_id": cId})
	if err != nil {
		return "", fmt.Errorf("unable to delete category: %v", err)
	}

	return "category delete successful", nil
}

func (r *Repository) populateCategoryBillboard(ctx context.Context, category *Category) error {
	var billboard Billboard
	err := r.billboardColl.FindOne(ctx, bson.M{"store_id": category.StoreID}).Decode(&billboard)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("billboard for store %s category not found: %v", category.StoreID.Hex(), err)
		}
		return fmt.Errorf("unable to fetch billboard for store category %s: %v", category.StoreID.Hex(), err)
	}

	category.Billboard = billboard
	return nil
}
