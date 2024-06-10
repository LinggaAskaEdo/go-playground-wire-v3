package repository

import (
	"context"
	"database/sql"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/errors"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/entity"
)

func (n *NewsRepositoryImpl) GetNewsByUrl(ctx context.Context, url string) (bool, error) {
	var result bool

	tx, err := n.db0.SQL.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	if err != nil {
		return false, errors.FindErrorType(err)
	}

	tx, result, err = n.getSQLNewsByURL(ctx, tx, url)
	if err != nil {
		return result, errors.FindErrorType(err)
	}

	if err := tx.Commit(); err != nil {
		return result, errors.FindErrorType(err)
	}

	return result, nil
}

func (n *NewsRepositoryImpl) CreateNews(ctx context.Context, v entity.News) (entity.News, error) {
	tx, err := n.db0.SQL.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	if err != nil {
		return v, errors.FindErrorType(err)
	}

	tx, v, err = n.createSQLNews(ctx, tx, v)
	if err != nil {
		tx.Rollback()

		return v, err
	}

	if err := tx.Commit(); err != nil {
		return v, errors.FindErrorType(err)
	}

	return v, nil
}

func (n *NewsRepositoryImpl) FindNewsByID(ctx context.Context, newsID int64) (entity.News, error) {
	entityNews := entity.News{}

	tx, err := n.db0.SQL.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	if err != nil {
		return entityNews, errors.FindErrorType(err)
	}

	tx, entityNews, err = n.getSQLNewsByID(ctx, tx, newsID)
	if err != nil {
		tx.Rollback()

		return entityNews, err
	}

	if err := tx.Commit(); err != nil {
		return entityNews, errors.FindErrorType(err)
	}

	return entityNews, nil
}
