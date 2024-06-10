package repository

import (
	"context"
	"database/sql"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/errors"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/product/entity"
)

func (n *ProductRepositoryImpl) FindProductByID(ctx context.Context, productID int64) (entity.Product, error) {
	entityProduct := entity.Product{}

	tx, err := n.db1.SQL.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	if err != nil {
		return entityProduct, errors.FindErrorType(err)
	}

	tx, entityProduct, err = n.getSQLProductByID(ctx, tx, productID)
	if err != nil {
		tx.Rollback()
		return entityProduct, err
	}

	if err := tx.Commit(); err != nil {
		return entityProduct, errors.FindErrorType(err)
	}

	return entityProduct, nil
}
