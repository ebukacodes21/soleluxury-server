package mongo

import "context"

func (r *MongoRepository) CreateProduct(ctx context.Context, product Product) (*Product, error) {
	return nil, nil
}

func (r *MongoRepository) UpdateProduct(ctx context.Context, productId string, product Product) error {
	return nil
}

func (r *MongoRepository) DeleteProduct(ctx context.Context, productId string) error {
	return nil
}
