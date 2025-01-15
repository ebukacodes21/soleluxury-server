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

	err = r.populateCategoryDetails(ctx, &category)
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

		err = r.populateCategoryDetails(ctx, &category)
		if err != nil {
			return nil, fmt.Errorf("unable to fetch associated details: %v", err)
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

func (r *Repository) populateCategoryDetails(ctx context.Context, category *Category) error {
	if err := r.populateCategoryBillboard(ctx, category); err != nil {
		return fmt.Errorf("unable to populate category billboard: %v", err)
	}

	if err := r.populateCategoryProducts(ctx, category); err != nil {
		return fmt.Errorf("unable to populate category products: %v", err)
	}

	return nil
}

func (r *Repository) populateCategoryBillboard(ctx context.Context, category *Category) error {
	var billboard Billboard
	var store Store
	err := r.billboardColl.FindOne(ctx, bson.M{"store_id": category.StoreID}).Decode(&billboard)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("billboard for store %s category not found: %v", category.StoreID.Hex(), err)
		}
		return fmt.Errorf("unable to fetch billboard for store category %s: %v", category.StoreID.Hex(), err)
	}

	err = r.storeColl.FindOne(ctx, bson.M{"_id": billboard.StoreID}).Decode(&store)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("store with id %s not found: %v", billboard.StoreID.Hex(), err)
		}
		return fmt.Errorf("unable to fetch store %s: %v", billboard.StoreID.Hex(), err)
	}

	billboard.Store = store
	category.Billboard = billboard
	return nil
}

func (r *Repository) populateCategoryProducts(ctx context.Context, category *Category) error {
	var products []Product

	cursor, err := r.productColl.Find(ctx, bson.M{"category_id": category.ID})
	if err != nil {
		return fmt.Errorf("unable to fetch products for category %s: %v", category.ID.Hex(), err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product Product
		var store Store
		var color Color
		var size Size

		if err := cursor.Decode(&product); err != nil {
			return fmt.Errorf("unable to decode product: %v", err)
		}

		//find product store, size and color
		err = r.storeColl.FindOne(ctx, bson.M{"_id": product.StoreID}).Decode(&store)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return fmt.Errorf("store with id %s not found: %v", product.StoreID.Hex(), err)
			}
			return fmt.Errorf("unable to fetch store %s: %v", product.StoreID.Hex(), err)
		}

		err = r.colorColl.FindOne(ctx, bson.M{"_id": product.ColorID}).Decode(&color)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return fmt.Errorf("color with id %s not found: %v", product.ColorID.Hex(), err)
			}
			return fmt.Errorf("unable to fetch color %s: %v", product.ColorID.Hex(), err)
		}

		err = r.sizeColl.FindOne(ctx, bson.M{"_id": product.SizeID}).Decode(&size)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return fmt.Errorf("size with id %s not found: %v", product.SizeID.Hex(), err)
			}
			return fmt.Errorf("unable to fetch size %s: %v", product.SizeID.Hex(), err)
		}

		// append to product
		size.Store = store
		color.Store = store
		product.Store = store
		product.Color = color
		product.Size = size
		product.Category = *category
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return fmt.Errorf("cursor error: %v", err)
	}

	category.Products = products
	return nil
}
