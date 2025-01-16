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

		// Fetch products for the category using an aggregation pipeline
		// This combines category, color, size, and store in one query
		pipeline := []bson.M{
			{
				"$match": bson.M{
					"category_id": category.ID,
				},
			},
			// Fetch category information (though we already have it in `category`)
			{
				"$lookup": bson.M{
					"from":         "categories", // Collection name for categories
					"localField":   "category_id",
					"foreignField": "_id",
					"as":           "category",
				},
			},
			// Fetch the color information for the product
			{
				"$lookup": bson.M{
					"from":         "colors", // Collection name for colors
					"localField":   "color_id",
					"foreignField": "_id",
					"as":           "color",
				},
			},
			// Fetch the size information for the product
			{
				"$lookup": bson.M{
					"from":         "sizes", // Collection name for sizes
					"localField":   "size_id",
					"foreignField": "_id",
					"as":           "size",
				},
			},
			// Add the store to each product (we can use `$addFields` to add it directly)
			{
				"$addFields": bson.M{
					"store": store, // Add store to each product
				},
			},
			{
				"$unwind": "$category", // Unwind to flatten the category array
			},
			{
				"$unwind": "$color", // Unwind to flatten the color array
			},
			{
				"$unwind": "$size", // Unwind to flatten the size array
			},
		}

		// Execute the aggregation query for products
		productCursor, err := r.productColl.Aggregate(ctx, pipeline)
		if err != nil {
			return fmt.Errorf("unable to aggregate products: %v", err)
		}
		defer productCursor.Close(ctx)

		var products []Product
		// Iterate over the cursor and decode each product
		for productCursor.Next(ctx) {
			var product Product
			if err := productCursor.Decode(&product); err != nil {
				return fmt.Errorf("unable to decode product: %v", err)
			}

			// Set store, category, size, and color to the product
			product.Store = *store
			product.Category = category
			product.Size.Store = *store
			product.Color.Store = *store

			// Append to the products list
			products = append(products, product)
		}

		// Check for any errors during cursor iteration
		if err := productCursor.Err(); err != nil {
			return fmt.Errorf("product cursor error: %v", err)
		}

		// Assign products to the category
		category.Products = products

		// Populate the Billboard in the category and set its store
		if err := r.billboardColl.FindOne(ctx, bson.M{"_id": category.BillboardID}).Decode(&category.Billboard); err != nil {
			return fmt.Errorf("unable to fetch billboard for category %s: %v", category.ID.Hex(), err)
		}

		// Set the store on the billboard
		category.Billboard.Store = *store

		// Append the populated category to the categories list
		categories = append(categories, category)
	}

	// Check if there was any error with the category cursor
	if err := cursor.Err(); err != nil {
		return fmt.Errorf("cursor error: %v", err)
	}

	// Finally, assign the populated categories to the store
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

	// Aggregation pipeline to fetch products with category, size, color, and store
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"store_id": store.ID, // Match products for the given store
			},
		},
		{
			"$lookup": bson.M{
				"from":         "categories", // Join with categories collection
				"localField":   "category_id",
				"foreignField": "_id",
				"as":           "category",
			},
		},
		{
			"$unwind": "$category", // Unwind to flatten the category array
		},
		{
			"$lookup": bson.M{
				"from":         "billboards", // Join with billboards collection
				"localField":   "category.billboard_id",
				"foreignField": "_id",
				"as":           "billboard",
			},
		},
		{
			"$unwind": "$billboard", // Unwind to flatten the billboard array
		},
		{
			"$lookup": bson.M{
				"from":         "colors", // Join with colors collection
				"localField":   "color_id",
				"foreignField": "_id",
				"as":           "color",
			},
		},
		{
			"$unwind": "$color", // Unwind to flatten the color array
		},
		{
			"$lookup": bson.M{
				"from":         "sizes", // Join with sizes collection
				"localField":   "size_id",
				"foreignField": "_id",
				"as":           "size",
			},
		},
		{
			"$unwind": "$size", // Unwind to flatten the size array
		},
		// Add the store to each product
		{
			"$addFields": bson.M{
				"store": store, // Add store information to the product
			},
		},
	}

	// Execute the aggregation query to get products with related data
	cursor, err := r.productColl.Aggregate(ctx, pipeline)
	if err != nil {
		return fmt.Errorf("unable to aggregate products for store %s: %v", store.ID.Hex(), err)
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode each product
	for cursor.Next(ctx) {
		var product Product
		if err := cursor.Decode(&product); err != nil {
			return fmt.Errorf("unable to decode product: %v", err)
		}

		// Populate the product's category, size, color, and store
		product.Category.Store = *store // Set the store on category
		product.Size.Store = *store     // Set the store on size
		product.Color.Store = *store    // Set the store on color

		// Populate the billboard for the category and set its store
		product.Category.Billboard.Store = *store

		// Append the product to the products list
		products = append(products, product)
	}

	// Check for any errors during cursor iteration
	if err := cursor.Err(); err != nil {
		return fmt.Errorf("cursor error: %v", err)
	}

	// Finally, assign the populated products to the store
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
