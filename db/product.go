package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// create product
func (r *Repository) CreateProduct(ctx context.Context, args *pb.CreateProductRequest) (*Product, error) {
	name := args.GetName()
	var existingProd Product
	var store Store
	var category Category
	var size Size
	var color Color

	// any exisiting product
	err := r.productColl.FindOne(ctx, bson.M{"name": name}).Decode(&existingProd)
	if err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("product %s already exists", name)
	}

	// get store
	storeId, err := bson.ObjectIDFromHex(args.GetStoreId())
	if err != nil {
		return nil, fmt.Errorf("invalid store id %s ", storeId)
	}
	err = r.storeColl.FindOne(ctx, bson.M{"_id": storeId}).Decode(&store)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("store with id %s not found", storeId)
	}

	// get category
	cateId, err := bson.ObjectIDFromHex(args.GetCategoryId())
	if err != nil {
		return nil, fmt.Errorf("invalid category id %s ", cateId)
	}
	err = r.categoryColl.FindOne(ctx, bson.M{"_id": cateId}).Decode(&category)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("category with id %s not found", cateId)
	}

	// get size
	sizeId, err := bson.ObjectIDFromHex(args.GetSizeId())
	if err != nil {
		return nil, fmt.Errorf("invalid size id %s ", sizeId)
	}
	err = r.sizeColl.FindOne(ctx, bson.M{"_id": sizeId}).Decode(&size)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("size with id %s not found", sizeId)
	}

	// get color
	colorId, err := bson.ObjectIDFromHex(args.GetColorId())
	if err != nil {
		return nil, fmt.Errorf("invalid color id %s ", cateId)
	}
	err = r.colorColl.FindOne(ctx, bson.M{"_id": colorId}).Decode(&color)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("color with id %s not found", colorId)
	}

	product := &Product{
		StoreID:     store.ID,
		CategoryID:  category.ID,
		SizeID:      size.ID,
		ColorID:     color.ID,
		Name:        args.GetName(),
		Description: args.GetDescription(),
		Price:       float64(args.GetPrice()),
		Images:      convertImage(args.GetImages()),
		IsFeatured:  args.GetIsFeatured(),
		IsArchived:  args.GetIsArchived(),
		CreatedAt:   time.Now(),
	}

	result, err := r.productColl.InsertOne(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("unable to create product: %v", err)
	}

	product.ID = result.InsertedID.(bson.ObjectID)
	return product, nil
}

// get products
func (r *Repository) GetProducts(ctx context.Context, args *pb.GetProductsRequest) ([]Product, error) {
	var products []Product
	sId, err := bson.ObjectIDFromHex(args.GetStoreId())
	if err != nil {
		return nil, fmt.Errorf("invalid store id for color")
	}

	cursor, err := r.productColl.Find(ctx, bson.M{"store_id": sId})
	if err != nil {
		return nil, fmt.Errorf("unable to fetch products for store: %v", err)
	}
	defer cursor.Close(ctx)

	// decode all the products
	for cursor.Next(ctx) {
		var product Product
		if err := cursor.Decode(&product); err != nil {
			return nil, fmt.Errorf("unable to decode product: %v", err)
		}

		if err := r.populateProductDetails(ctx, &product); err != nil {
			return nil, fmt.Errorf("unable to populate product with details: %v", err)
		}

		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return products, nil
}

// get product by id
func (r *Repository) GetProductByID(ctx context.Context, req *pb.GetProductRequest) (*Product, error) {
	var product Product
	pID, err := bson.ObjectIDFromHex(req.GetProductId())
	if err != nil {
		return nil, fmt.Errorf("invalid product id: %v", err)
	}

	err = r.productColl.FindOne(ctx, bson.M{"_id": pID}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("product with ID %s not found", pID)
		}
		return nil, fmt.Errorf("unable to fetch product by ID: %v", err)
	}

	err = r.populateProductDetails(ctx, &product)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch associated details: %v", err)
	}

	return &product, nil
}

// update product with id
func (r *Repository) UpdateProduct(ctx context.Context, args *pb.UpdateProductRequest) (string, error) {
	var product Product
	id, err := bson.ObjectIDFromHex(args.GetProductId())
	if err != nil {
		return "", fmt.Errorf("invalid product id: %v", err)
	}

	cateId, err := bson.ObjectIDFromHex(args.GetCategoryId())
	if err != nil {
		return "", fmt.Errorf("invalid category id: %v", err)
	}

	colorId, err := bson.ObjectIDFromHex(args.GetColorId())
	if err != nil {
		return "", fmt.Errorf("invalid color id: %v", err)
	}

	sizeId, err := bson.ObjectIDFromHex(args.GetSizeId())
	if err != nil {
		return "", fmt.Errorf("invalid size id: %v", err)
	}

	err = r.productColl.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("product with ID %s not found", id)
		}
		return "", fmt.Errorf("unable to fetch product: %v", err)
	}

	update := bson.M{
		"$set": bson.M{
			"name":        args.GetName(),
			"description": args.GetDescription(),
			"price":       args.GetPrice(),
			"category_id": cateId,
			"color_id":    colorId,
			"size_id":     sizeId,
			"images":      convertImage(args.GetImages()),
			"updated_at":  time.Now(),
		},
	}

	_, err = r.productColl.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return "", fmt.Errorf("unable to update product: %v", err)
	}

	return "product update successful", nil
}

// delete product with id
func (r *Repository) DeleteProduct(ctx context.Context, arg *pb.DeleteProductRequest) (string, error) {
	pId, err := bson.ObjectIDFromHex(arg.GetProductId())
	if err != nil {
		return "", fmt.Errorf("invalid product id: %v", err)
	}

	_, err = r.productColl.DeleteOne(ctx, bson.M{"_id": pId})
	if err != nil {
		return "", fmt.Errorf("unable to delete product: %v", err)
	}

	return "product delete successful", nil
}

func convertImage(images []*pb.Item) []Image {
	var dbImages []Image
	for _, image := range images {
		dbImage := Image{
			URL: image.Url,
		}

		dbImages = append(dbImages, dbImage)
	}
	return dbImages
}

func (r *Repository) populateProductDetails(ctx context.Context, product *Product) error {
	var store Store
	var category Category
	var size Size
	var color Color
	var billboard Billboard

	// find product store
	err := r.storeColl.FindOne(ctx, bson.M{"_id": product.StoreID}).Decode(&store)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("store with id %s not found: %v", product.StoreID.Hex(), err)
		}
		return fmt.Errorf("unable to fetch product store %s: %v", product.StoreID.Hex(), err)
	}

	// find product category
	err = r.categoryColl.FindOne(ctx, bson.M{"_id": product.CategoryID}).Decode(&category)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("category with id %s not found: %v", product.CategoryID.Hex(), err)
		}
		return fmt.Errorf("unable to fetch product category %s: %v", product.CategoryID.Hex(), err)
	}

	// here find category billboard
	err = r.billboardColl.FindOne(ctx, bson.M{"_id": category.BillboardID}).Decode(&billboard)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("billboard with id %s not found: %v", category.BillboardID.Hex(), err)
		}
		return fmt.Errorf("unable to fetch category billboard %s: %v", category.BillboardID.Hex(), err)
	}

	// find product size
	err = r.sizeColl.FindOne(ctx, bson.M{"_id": product.SizeID}).Decode(&size)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("size with id %s not found: %v", product.SizeID.Hex(), err)
		}
		return fmt.Errorf("unable to fetch size product %s: %v", product.SizeID.Hex(), err)
	}

	// find product color
	err = r.colorColl.FindOne(ctx, bson.M{"_id": product.ColorID}).Decode(&color)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("color with id %s not found: %v", product.ColorID.Hex(), err)
		}
		return fmt.Errorf("unable to fetch color product %s: %v", product.ColorID.Hex(), err)
	}

	// populate fields
	size.Store = store
	color.Store = store
	billboard.Store = store
	category.Billboard = billboard

	product.Store = store
	product.Category = category
	product.Color = color
	product.Size = size
	return nil
}
