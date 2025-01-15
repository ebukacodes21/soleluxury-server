package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// create store
func (r *Repository) CreateStore(ctx context.Context, arg *pb.CreateStoreRequest) (*Store, error) {
	name := arg.GetName()
	var exStore Store

	err := r.storeColl.FindOne(ctx, bson.M{"name": name}).Decode(&exStore)
	if err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("store %s already exists", name)
	}

	store := &Store{
		Name:      name,
		CreatedAt: time.Now(),
	}

	result, err := r.storeColl.InsertOne(ctx, store)
	if err != nil {
		return nil, fmt.Errorf("unable to create store: %v", err)
	}

	store.ID = result.InsertedID.(bson.ObjectID)
	store.CreatedAt = time.Now()
	return store, nil
}

// get first created store
func (r *Repository) GetFirstStore(ctx context.Context) (*Store, error) {
	var store Store
	cursor, err := r.storeColl.Find(ctx, bson.M{}, options.Find().SetSort(bson.M{"created_at": 1}).SetLimit(1))
	if err != nil {
		return nil, fmt.Errorf("unable to fetch stores: %v", err)
	}
	defer cursor.Close(ctx)

	// any docs?
	if !cursor.Next(ctx) {
		if err := cursor.Err(); err != nil {
			return nil, fmt.Errorf("cursor error: %v", err)
		}
		return nil, fmt.Errorf("no stores found")
	}

	if err := cursor.Decode(&store); err != nil {
		return nil, fmt.Errorf("unable to decode the first store: %v", err)
	}

	err = r.populateStoreDetails(ctx, &store)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch associated details: %v", err)
	}

	return &store, nil
}

// get store by id
func (r *Repository) GetStoreByID(ctx context.Context, req *pb.GetStoreRequest) (*Store, error) {
	var store Store
	sID, err := bson.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid store id: %v", err)
	}

	err = r.storeColl.FindOne(ctx, bson.M{"_id": sID}).Decode(&store)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("store with ID %s not found", sID)
		}
		return nil, fmt.Errorf("unable to fetch store by ID: %v", err)
	}

	err = r.populateStoreDetails(ctx, &store)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch associated details: %v", err)
	}

	return &store, nil
}

// get all stores.
func (r *Repository) GetAllStores(ctx context.Context) ([]Store, error) {
	var stores []Store
	cursor, err := r.storeColl.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("unable to fetch stores: %v", err)
	}
	defer cursor.Close(ctx)

	// decode all the stores
	for cursor.Next(ctx) {
		var store Store
		if err := cursor.Decode(&store); err != nil {
			return nil, fmt.Errorf("unable to decode store: %v", err)
		}

		err = r.populateStoreDetails(ctx, &store)
		if err != nil {
			return nil, fmt.Errorf("unable to fetch associated details: %v", err)
		}

		stores = append(stores, store)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return stores, nil
}

// update store with id
func (r *Repository) UpdateStore(ctx context.Context, args *pb.UpdateStoreRequest) (string, error) {
	var store Store
	storeId, err := bson.ObjectIDFromHex(args.GetId())
	if err != nil {
		return "", fmt.Errorf("invalid store id: %v", err)
	}

	err = r.storeColl.FindOne(ctx, bson.M{"_id": storeId}).Decode(&store)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("store with ID %s not found", storeId)
		}
		return "", fmt.Errorf("unable to fetch store: %v", err)
	}

	update := bson.M{
		"$set": bson.M{
			"name":       args.GetName(),
			"updated_at": time.Now(),
		},
	}

	_, err = r.storeColl.UpdateOne(ctx, bson.M{"_id": storeId}, update)
	if err != nil {
		return "", fmt.Errorf("unable to update store: %v", err)
	}

	return "store update successful", nil
}

// delete store with associated models
func (r *Repository) DeleteStore(ctx context.Context, arg *pb.DeleteStoreRequest) (string, error) {
	storeId, err := bson.ObjectIDFromHex(arg.GetId())
	if err != nil {
		return "", fmt.Errorf("invalid store id: %v", err)
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

	_, err = r.categoryColl.DeleteMany(ctx, bson.M{"store_id": storeId})
	if err != nil {
		return "", fmt.Errorf("unable to delete categories for store %s: %v", storeId.Hex(), err)
	}

	_, err = r.billboardColl.DeleteMany(ctx, bson.M{"store_id": storeId})
	if err != nil {
		return "", fmt.Errorf("unable to delete billboards for store %s: %v", storeId.Hex(), err)
	}

	_, err = r.sizeColl.DeleteMany(ctx, bson.M{"store_id": storeId})
	if err != nil {
		return "", fmt.Errorf("unable to delete sizes for store %s: %v", storeId.Hex(), err)
	}

	_, err = r.colorColl.DeleteMany(ctx, bson.M{"store_id": storeId})
	if err != nil {
		return "", fmt.Errorf("unable to delete colors for store %s: %v", storeId.Hex(), err)
	}

	_, err = r.productColl.DeleteMany(ctx, bson.M{"store_id": storeId})
	if err != nil {
		return "", fmt.Errorf("unable to delete products for store %s: %v", storeId.Hex(), err)
	}

	_, err = r.orderColl.DeleteMany(ctx, bson.M{"store_id": storeId})
	if err != nil {
		return "", fmt.Errorf("unable to delete orders for store %s: %v", storeId.Hex(), err)
	}

	_, err = r.storeColl.DeleteOne(ctx, bson.M{"_id": storeId})
	if err != nil {
		return "", fmt.Errorf("unable to delete store: %v", err)
	}

	err = session.CommitTransaction(ctx)
	if err != nil {
		return "", fmt.Errorf("unable to commit transaction: %v", err)
	}

	return "store delete successful", nil
}

func (r *Repository) populateStoreDetails(ctx context.Context, store *Store) error {
	if err := r.populateStoreCategories(ctx, store); err != nil {
		return fmt.Errorf("unable to fetch categories: %v", err)
	}

	if err := r.populateBillboards(ctx, store); err != nil {
		return fmt.Errorf("unable to fetch billboards: %v", err)
	}

	if err := r.populateSizes(ctx, store); err != nil {
		return fmt.Errorf("unable to fetch sizes: %v", err)
	}

	if err := r.populateColors(ctx, store); err != nil {
		return fmt.Errorf("unable to fetch colors: %v", err)
	}

	if err := r.populateProducts(ctx, store); err != nil {
		return fmt.Errorf("unable to fetch products: %v", err)
	}

	if err := r.populateOrders(ctx, store); err != nil {
		return fmt.Errorf("unable to fetch orders: %v", err)
	}

	return nil
}

func (r *Repository) populateStoreCategories(ctx context.Context, store *Store) error {
	var categories []Category
	var billboard Billboard
	cursor, err := r.categoryColl.Find(ctx, bson.M{"store_id": store.ID})
	if err != nil {
		return fmt.Errorf("unable to fetch categories for store %s: %v", store.ID.Hex(), err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var category Category
		if err := cursor.Decode(&category); err != nil {
			return fmt.Errorf("unable to decode category: %v", err)
		}

		if err := r.billboardColl.FindOne(ctx, bson.M{"_id": category.BillboardID, "store_id": store.ID}).Decode(&billboard); err != nil {
			return fmt.Errorf("unable to decode billboard: %v", err)
		}

		billboard.Store = *store
		category.Store = *store
		category.Billboard = billboard
		categories = append(categories, category)
	}

	if err := cursor.Err(); err != nil {
		return fmt.Errorf("cursor error: %v", err)
	}

	store.Categories = categories
	return nil
}

func (r *Repository) populateBillboards(ctx context.Context, store *Store) error {
	var billboards []Billboard
	cursor, err := r.billboardColl.Find(ctx, bson.M{"store_id": store.ID})
	if err != nil {
		return fmt.Errorf("unable to fetch billboards for store %s: %v", store.ID.Hex(), err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var billboard Billboard
		if err := cursor.Decode(&billboard); err != nil {
			return fmt.Errorf("unable to decode billboard: %v", err)
		}

		billboard.Store = *store
		billboards = append(billboards, billboard)
	}

	if err := cursor.Err(); err != nil {
		return fmt.Errorf("cursor error: %v", err)
	}

	store.Billboards = billboards
	return nil
}

func (r *Repository) populateSizes(ctx context.Context, store *Store) error {
	var sizes []Size
	cursor, err := r.sizeColl.Find(ctx, bson.M{"store_id": store.ID})
	if err != nil {
		return fmt.Errorf("unable to fetch sizes for store %s: %v", store.ID.Hex(), err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var size Size
		if err := cursor.Decode(&size); err != nil {
			return fmt.Errorf("unable to decode size: %v", err)
		}

		size.Store = *store
		sizes = append(sizes, size)
	}

	if err := cursor.Err(); err != nil {
		return fmt.Errorf("cursor error: %v", err)
	}

	store.Sizes = sizes
	return nil
}

func (r *Repository) populateColors(ctx context.Context, store *Store) error {
	var colors []Color
	cursor, err := r.colorColl.Find(ctx, bson.M{"store_id": store.ID})
	if err != nil {
		return fmt.Errorf("unable to fetch colors for store %s: %v", store.ID.Hex(), err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var color Color
		if err := cursor.Decode(&color); err != nil {
			return fmt.Errorf("unable to decode color: %v", err)
		}
		color.Store = *store
		colors = append(colors, color)
	}

	if err := cursor.Err(); err != nil {
		return fmt.Errorf("cursor error: %v", err)
	}

	store.Colors = colors
	return nil
}

func (r *Repository) populateProducts(ctx context.Context, store *Store) error {
	var products []Product
	cursor, err := r.productColl.Find(ctx, bson.M{"store_id": store.ID})
	if err != nil {
		return fmt.Errorf("unable to fetch products for store %s: %v", store.ID.Hex(), err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var color Color
		var size Size
		var product Product
		var category Category
		var billboard Billboard
		if err := cursor.Decode(&product); err != nil {
			return fmt.Errorf("unable to decode product: %v", err)
		}

		if err := r.categoryColl.FindOne(ctx, bson.M{"_id": product.CategoryID}).Decode(&category); err != nil {
			return fmt.Errorf("unable to fetch category for product %s: %v", product.CategoryID.Hex(), err)
		}

		if err := r.billboardColl.FindOne(ctx, bson.M{"_id": category.BillboardID}).Decode(&billboard); err != nil {
			return fmt.Errorf("unable to fetch billboard for category %s: %v", category.BillboardID.Hex(), err)
		}

		if err := r.colorColl.FindOne(ctx, bson.M{"_id": product.ColorID}).Decode(&color); err != nil {
			return fmt.Errorf("unable to fetch color for product %s: %v", product.ColorID.Hex(), err)
		}

		if err := r.sizeColl.FindOne(ctx, bson.M{"_id": product.SizeID}).Decode(&size); err != nil {
			return fmt.Errorf("unable to fetch size for product %s: %v", product.SizeID.Hex(), err)
		}

		size.Store = *store
		color.Store = *store
		billboard.Store = *store
		category.Billboard = billboard
		product.Category = category
		product.Store = *store
		product.Size = size
		product.Color = color
		products = append(products, product)
		category.Products = products
	}

	if err := cursor.Err(); err != nil {
		return fmt.Errorf("cursor error: %v", err)
	}

	store.Products = products
	return nil
}

func (r *Repository) populateOrders(ctx context.Context, store *Store) error {
	var orders []Order
	cursor, err := r.orderColl.Find(ctx, bson.M{"store_id": store.ID})
	if err != nil {
		return fmt.Errorf("unable to fetch orders for store %s: %v", store.ID.Hex(), err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var order Order
		if err := cursor.Decode(&order); err != nil {
			return fmt.Errorf("unable to decode order: %v", err)
		}
		orders = append(orders, order)
	}

	if err := cursor.Err(); err != nil {
		return fmt.Errorf("cursor error: %v", err)
	}

	store.Orders = orders
	return nil
}
