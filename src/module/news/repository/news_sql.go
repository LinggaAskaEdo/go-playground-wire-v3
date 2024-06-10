package repository

import (
	"context"
	"database/sql"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/errors"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/entity"
)

func (n *NewsRepositoryImpl) getSQLNewsByURL(ctx context.Context, tx *sql.Tx, url string) (*sql.Tx, bool, error) {
	var isExist bool

	err := tx.QueryRowContext(ctx, GetNewsByUrl, url).Scan(&isExist)
	if err != nil {
		return tx, isExist, errors.FindErrorType(err)
	}

	return tx, isExist, nil
}

func (n *NewsRepositoryImpl) createSQLNews(ctx context.Context, tx *sql.Tx, v entity.News) (*sql.Tx, entity.News, error) {
	_, err := tx.ExecContext(ctx, CreateNews, v.Title, v.URL, v.Content, v.Summary, v.ArticleTS, v.PublishedDate, v.Inserted, v.Updated)
	if err != nil {
		return tx, v, errors.FindErrorType(err)
	}

	return tx, v, nil
}

func (n *NewsRepositoryImpl) getSQLNewsByID(ctx context.Context, tx *sql.Tx, newsID int64) (*sql.Tx, entity.News, error) {
	result := entity.News{}

	rows, err := tx.QueryContext(ctx, GetNewsByID, newsID)
	if err != nil {
		return tx, result, errors.FindErrorType(err)
	}

	if rows.Next() {
		err := rows.Scan(
			&result.ID,
			&result.Title,
			&result.URL,
			&result.Content,
			&result.Summary,
			&result.ArticleTS,
			&result.PublishedDate,
			&result.Inserted,
			&result.Updated,
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
