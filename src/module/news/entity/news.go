package entity

import "database/sql"

type News struct {
	ID            int            `db:"id"`
	Title         string         `db:"title"`
	URL           string         `db:"url"`
	Content       string         `db:"content"`
	Summary       sql.NullString `db:"summary"`
	ArticleTS     int64          `db:"article_ts"`
	PublishedDate sql.NullTime   `db:"published_date"`
	Inserted      sql.NullTime   `db:"inserted"`
	Updated       sql.NullTime   `db:"updated"`
}
