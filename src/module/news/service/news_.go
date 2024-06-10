package service

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xmlquery"
	"github.com/rs/zerolog/log"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/config"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/errors"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/dto"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/entity"
)

func (nsi *NewsServiceImpl) GetLatestNewsRSS(ctx context.Context) error {
	log.Info().Msg("GetLatestNewsRSS...")

	start := time.Now()
	counter := 0

	doc, err := xmlquery.LoadURL(config.Get().Module.News.Service.UrlNewsRSS)
	if err != nil {
		log.Err(err).Msg("Error GetLatestNews")
	}

	v := entity.News{}

	channel := xmlquery.Find(doc, "//item")

	for _, n := range channel {
		if n := n.SelectElement("title"); n != nil {
			title := n.InnerText()
			v.Title = title
		}

		if n := n.SelectElement("link"); n != nil {
			link := n.InnerText()
			v.URL = link

			docDetail, err := htmlquery.LoadURL(link)
			if err != nil {
				log.Err(err).Msg("Error GetLatestNews")
			}

			docDataDetail := htmlquery.FindOne(docDetail, "//div[@class = 'post-content clearfix']")
			strDocDataDetail := htmlquery.InnerText(docDataDetail)
			strDocDataDetail = strings.TrimSpace(strDocDataDetail)
			strDocDataDetail = strings.ReplaceAll(strDocDataDetail, "\t", "")
			strDocDataDetail = strings.ReplaceAll(strDocDataDetail, "\n", "")

			v.Content = strDocDataDetail
		}

		if n := n.SelectElement("pubDate"); n != nil {
			pubDate := n.InnerText()

			timePub, err := time.Parse(time.RFC1123Z, pubDate)
			if err != nil {
				log.Err(err).Msg("Error GetLatestNews")
			}

			v.PublishedDate = sql.NullTime{Time: timePub, Valid: true}

			timestamp := timePub.Unix()
			v.ArticleTS = int64(timestamp)
		}

		v.Inserted = sql.NullTime{Time: time.Now(), Valid: true}

		status, err := nsi.newsRepository.GetNewsByUrl(ctx, v.URL)
		if err != nil {
			log.Err(err).Msg("Error fetch news data")

			return errors.FindErrorType(err)
		}

		log.Info().Msgf("URL: %s, Status:  %t", v.URL, status)

		if !status {
			_, err := nsi.newsRepository.CreateNews(ctx, v)
			if err != nil {
				log.Err(err).Msg("Error GetLatestNews")
			}

			counter++
		}
	}

	duration := time.Since(start)

	log.Info().Msgf("%d data added successfully in %f seconds", counter, duration.Seconds())

	return nil
}

func (nsi *NewsServiceImpl) GetLatestNewsIndex(ctx context.Context) error {
	log.Info().Msg("GetLatestNewsIndex...")

	start := time.Now()
	counter := 0

	doc, err := htmlquery.LoadURL(config.Get().Module.News.Service.UrlNewsIndex)
	if err != nil {
		return errors.FindErrorType(err)
	}

	log.Info().Msg("AAA")
	list := htmlquery.FindOne(doc, "//ul[@class = 'list-news indeks-new']")
	data := htmlquery.Find(list, "//div[@class = 'col-sm-4']//a")
	for _, n := range data {
		span := htmlquery.FindOne(n, "//span") // premium content always contains span tag
		if span == nil {
			log.Info().Msgf("BBB --> href: %s, title: %s", htmlquery.SelectAttr(n, "href"), htmlquery.SelectAttr(n, "title"))
		}
	}

	// for _, n := range list {
	// 	// link := htmlquery.FindOne(n, "//href")
	// 	log.Info().Msg("BBB")
	// 	log.Info().Msgf("HREF: %s", htmlquery.SelectAttr(n, "href"))
	// 	log.Info().Msgf("TITLE: %s", htmlquery.SelectAttr(n, "title"))
	// }
	log.Info().Msg("CCC")
	duration := time.Since(start)

	log.Info().Msgf("%d data added successfully in %f seconds", counter, duration.Seconds())

	return nil
}

func (nsi *NewsServiceImpl) FindNewsByID(ctx context.Context, newsID int64) (dto.NewsRespBody, error) {
	newsResp := dto.NewsRespBody{}

	news, err := nsi.newsRepository.FindNewsByID(ctx, newsID)
	if err != nil {
		log.Err(err).Msg("Error fetch news data")

		return newsResp, errors.FindErrorType(err)
	}

	newsResp = dto.CreateNewsResp(news)

	return newsResp, nil
}
