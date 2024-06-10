package repository

import (
	"context"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/database"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/entity"
)

type NewsRepository interface {
	GetNewsByUrl(ctx context.Context, url string) (bool, error)
	CreateNews(ctx context.Context, v entity.News) (entity.News, error)
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
