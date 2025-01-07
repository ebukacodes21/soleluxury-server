package db

import "context"

type CreateProductTxParams struct {
	CreateProductColorParams
	*CreateProductSizeParams
	*CreateProductCategoryParams
}

type CreateProductTxResponse struct {
	ProductColor    ProductColor
	ProductSize     ProductSize
	ProductCategory ProductCategory
}

func (sr *SoleluxuryRepository) CreateProductTx(ctx context.Context, args CreateProductTxParams) (CreateProductTxResponse, error) {
	var result CreateProductTxResponse

	err := sr.execTx(ctx, func(queries *Queries) error {
		var err error
		result.ProductColor, err = sr.CreateProductColor(ctx, args.CreateProductColorParams)
		if err != nil {
			return err
		}

		result.ProductSize, err = sr.CreateProductSize(ctx, *args.CreateProductSizeParams)
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
