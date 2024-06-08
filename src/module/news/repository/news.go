package repository

import (
	"context"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/database"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/entity"
)

type NewsRepository interface {
	FindNewsByID(ctx context.Context, newsID int64) (entity.News, error)
}

type NewsRepositoryImpl struct {
	db0 *database.MysqlImpl
	db1 *database.PostgresImpl
}

func NewNewsRepository(
	db0 *database.MysqlImpl,
	db1 *database.PostgresImpl,
) *NewsRepositoryImpl {
	return &NewsRepositoryImpl{
		db0: db0,
		db1: db1,
	}
}
