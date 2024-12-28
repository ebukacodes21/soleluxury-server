package db

import "context"

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user User) error
}

type CreateUserTxResponse struct {
	User User
}

func (sr *SoleluxuryRepository) CreateUserTx(ctx context.Context, args CreateUserTxParams) (CreateUserTxResponse, error) {
	var result CreateUserTxResponse

	err := sr.execTx(ctx, func(queries *Queries) error {
		var err error
		result.User, err = sr.CreateUser(ctx, args.CreateUserParams)
		if err != nil {
			return err
		}

		return args.AfterCreate(result.User)
	})

	return result, err
}
