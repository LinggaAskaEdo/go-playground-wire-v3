package repository

import (
	"context"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/database"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/product/entity"
)

type ProductRepository interface {
	FindProductByID(ctx context.Context, productID int64) (entity.Product, error)
}

type ProductRepositoryImpl struct {
	db0 *database.MysqlImpl
	db1 *database.PostgresImpl
}

func NewProductRepository(
	db0 *database.MysqlImpl,
	db1 *database.PostgresImpl,
) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		db0: db0,
		db1: db1,
	}
}
