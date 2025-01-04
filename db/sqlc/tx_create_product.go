package db

import "context"

type CreateProductTxParams struct {
	CreateProductParams
	CreateImageParams
	*CreateProductColorParams
	*CreateProductSizeParams
	*CreateProductStoreParams
	*CreateProductCategoryParams
}

type CreateProductTxResponse struct {
	Product         Product
	Image           Image
	ProductColor    ProductColor
	ProductSize     ProductSize
	ProductStore    ProductStore
	ProductCategory ProductCategory
}

func (sr *SoleluxuryRepository) CreateProductTx(ctx context.Context, args CreateProductTxParams) (CreateProductTxResponse, error) {
	var result CreateProductTxResponse

	err := sr.execTx(ctx, func(queries *Queries) error {
		var err error
		result.Product, err = sr.CreateProduct(ctx, args.CreateProductParams)
		if err != nil {
			return err
		}

		result.Image, err = sr.CreateImage(ctx, args.CreateImageParams)
		if err != nil {
			return err
		}

		result.ProductColor, err = sr.CreateProductColor(ctx, *args.CreateProductColorParams)
		if err != nil {
			return err
		}

		result.ProductSize, err = sr.CreateProductSize(ctx, *args.CreateProductSizeParams)
		if err != nil {
			return err
		}

		result.ProductStore, err = sr.CreateProductStore(ctx, *args.CreateProductStoreParams)
		if err != nil {
			return err
		}

		result.ProductCategory, err = sr.CreateProductCategory(ctx, *args.CreateProductCategoryParams)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
