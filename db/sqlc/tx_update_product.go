package db

import "context"

type UpdateProductTxParams struct {
	UpdateProductParams
	UpdateProductColorParams
	*UpdateProductSizeParams
	*UpdateProductCategoryParams
}

func (sr *SoleluxuryRepository) UpdateProductTx(ctx context.Context, args UpdateProductTxParams) error {
	err := sr.execTx(ctx, func(queries *Queries) error {
		var err error
		err = sr.UpdateProduct(ctx, args.UpdateProductParams)
		if err != nil {
			return err
		}

		err = sr.UpdateProductColor(ctx, args.UpdateProductColorParams)
		if err != nil {
			return err
		}

		err = sr.UpdateProductSize(ctx, *args.UpdateProductSizeParams)
		if err != nil {
			return err
		}

		err = sr.UpdateProductCategory(ctx, *args.UpdateProductCategoryParams)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
