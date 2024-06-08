package repository

const (
	GetNewsByID = `
		SELECT 
			id, title, url, content, summary, article_ts, published_date, inserted, updated 
		FROM 
			news_article 
		WHERE 
			id = ?
	`
)
