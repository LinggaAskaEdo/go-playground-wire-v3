package dto

import (
	"database/sql"

	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/entity"
)

const (
	timeFormat = "Monday, 02 January 2006 15:04 MST"
)

type NewsRespBody struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	URL           string `json:"url"`
	Content       string `json:"content"`
	Summary       string `json:"summary"`
	ArticleTS     int64  `json:"article_ts"`
	PublishedDate string `json:"published_date"`
	Inserted      string `json:"created_at"`
	Updated       string `json:"updated_at"`
}

func CreateNewsResp(data entity.News) NewsRespBody {
	return NewsRespBody{
		ID:            data.ID,
		Title:         data.Title,
		URL:           data.URL,
		Content:       data.Content,
		Summary:       isNullString(data.Summary),
		ArticleTS:     data.ArticleTS,
		PublishedDate: isNullTime(data.PublishedDate),
		Inserted:      isNullTime(data.Inserted),
		Updated:       isNullTime(data.Updated),
	}
}

func isNullString(s sql.NullString) string {
	if !s.Valid {
		return ""
	}

	return s.String
}

func isNullTime(s sql.NullTime) string {
	if !s.Valid {
		return ""
	}

	return s.Time.Format(timeFormat)
}
