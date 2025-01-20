package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) CreateOrder(ctx context.Context, args *pb.CreateOrderRequest) ([]Product, error) {
	var productIDs []bson.ObjectID
	for _, id := range args.GetItems() {
		objectID, err := bson.ObjectIDFromHex(id)
		if err != nil {
			return nil, fmt.Errorf("invalid product id: %s", id)
		}
		productIDs = append(productIDs, objectID)
	}

	cursor, err := r.productColl.Find(ctx, bson.M{"_id": bson.M{"$in": productIDs}})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}
	defer cursor.Close(ctx)

	var products []Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, fmt.Errorf("failed to decode products: %w", err)
	}

	if len(products) == 0 {
		return nil, fmt.Errorf("no products found with the given ids")
	}

	storeID := products[0].StoreID

	var items []bson.M
	for _, product := range products {
		items = append(items, bson.M{
			"product_id": product.ID,
			"name":       product.Name,
			"price":      product.Price,
		})
	}

	order := &Order{
		StoreID:   storeID,
		Items:     items,
		IsPaid:    false,
		CreatedAt: time.Now(),
	}

	result, err := r.orderColl.InsertOne(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	order.ID = result.InsertedID.(bson.ObjectID)
	return products, nil
}
