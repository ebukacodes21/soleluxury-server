package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) CreateOrder(ctx context.Context, args *pb.CreateOrderRequest) (*Order, []Product, error) {
	var productIDs []bson.ObjectID
	for _, id := range args.GetItems() {
		objectID, err := bson.ObjectIDFromHex(id)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid product id: %s", id)
		}
		productIDs = append(productIDs, objectID)
	}

	cursor, err := r.productColl.Find(ctx, bson.M{"_id": bson.M{"$in": productIDs}})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch products: %w", err)
	}
	defer cursor.Close(ctx)

	var products []Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, nil, fmt.Errorf("failed to decode products: %w", err)
	}

	if len(products) == 0 {
		return nil, nil, fmt.Errorf("no products found with the given ids")
	}

	storeID := products[0].StoreID

	var items []bson.M
	for _, product := range products {
		items = append(items, bson.M{
			"product_id": product.ID,
			"store_id":   product.StoreID,
			"name":       product.Name,
			"price":      product.Price,
		})
	}

	order := &Order{
		StoreID:   storeID,
		Items:     items,
		IsPaid:    false,
		Phone:     args.GetPhone(),
		Address:   args.GetAddress(),
		CreatedAt: time.Now(),
	}

	result, err := r.orderColl.InsertOne(ctx, order)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create order: %w", err)
	}

	order.ID = result.InsertedID.(bson.ObjectID)
	return order, products, nil
}

func (r *Repository) GetOrders(ctx context.Context) ([]Order, error) {
	var orders []Order
	cursor, err := r.orderColl.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("unable to fetch orders: %v", err)
	}
	defer cursor.Close(ctx)

	// decode all the orders
	for cursor.Next(ctx) {
		var order Order
		if err := cursor.Decode(&order); err != nil {
			return nil, fmt.Errorf("unable to decode order: %v", err)
		}
		orders = append(orders, order)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return orders, nil
}

func (r *Repository) UpdateOrder(ctx context.Context, args *pb.UpdateOrderRequest) (string, error) {
	var order Order
	id, err := bson.ObjectIDFromHex(args.GetOrderId())
	if err != nil {
		return "", fmt.Errorf("invalid order id: %s", id)
	}

	err = r.orderColl.FindOne(ctx, bson.M{"_id": id}).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("order with ID %s not found", id)
		}
		return "", fmt.Errorf("unable to fetch order: %v", err)
	}

	update := bson.M{
		"$set": bson.M{
			"is_paid":    true,
			"updated_at": time.Now(),
		},
	}

	_, err = r.orderColl.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return "", fmt.Errorf("unable to update order: %v", err)
	}

	return "order update successful", nil
}
