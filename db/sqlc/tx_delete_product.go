package db

import "context"

func (sr *SoleluxuryRepository) DeleteProductTx(ctx context.Context, arg int64) error {
	err := sr.execTx(ctx, func(queries *Queries) error {
		var err error
		err = sr.DeleteProductCategory(ctx, arg)
		if err != nil {
			return err
		}

		err = sr.DeleteProductColor(ctx, arg)
		if err != nil {
			return err
		}

		err = sr.DeleteProductSize(ctx, arg)
		if err != nil {
			return err
		}

		err = sr.DeleteProductCategory(ctx, arg)
		if err != nil {
			return err
		}

		err = sr.DeleteProduct(ctx, arg)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
