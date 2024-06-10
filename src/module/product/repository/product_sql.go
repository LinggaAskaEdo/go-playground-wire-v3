package repository

import (
	"context"
	"database/sql"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/errors"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/product/entity"
)

func (n *ProductRepositoryImpl) getSQLProductByID(ctx context.Context, tx *sql.Tx, newsID int64) (*sql.Tx, entity.Product, error) {
	result := entity.Product{}

	rows, err := tx.QueryContext(ctx, GetProductByID, newsID)
	if err != nil {
		return tx, result, errors.FindErrorType(err)
	}

	if rows.Next() {
		err := rows.Scan(
			&result.ID,
			&result.Name,
			&result.Description,
			&result.Cost,
			&result.Price,
			&result.Category.ID,
			&result.Category.Name,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return tx, result, errors.NotFound("Data not found")
			}

			return tx, result, errors.FindErrorType(err)
		}

		return tx, result, nil
	} else {
		return tx, result, errors.NotFound("Data not found")
	}
}
