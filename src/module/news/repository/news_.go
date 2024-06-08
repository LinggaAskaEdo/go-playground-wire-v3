package repository

import (
	"context"
	"database/sql"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/errors"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/entity"
)

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
