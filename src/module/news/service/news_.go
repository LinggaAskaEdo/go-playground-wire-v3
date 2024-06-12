package service

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xmlquery"
	"github.com/rs/zerolog/log"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/config"
	errs "github.com/linggaaskaedo/go-playground-wire-v3/lib/errors"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/dto"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/entity"
)

const (
	GOOGLE_TAG = "googletag.cmd.push(function() { googletag.display(\"div-gpt-ad-parallax\"); });"
	CEK_BERITA = "Cek Berita dan Artikel yang lain di Google News dan WA Channel"
)

func (nsi *NewsServiceImpl) GetLatestNewsRSS(ctx context.Context) error {
	log.Info().Msg("GetLatestNewsRSS...")

	start := time.Now()
	counter := 0

	doc, err := xmlquery.LoadURL(config.Get().Module.News.Service.UrlNewsRSS)
	if err != nil {
		log.Err(err).Msg("Error GetLatestNewsRSS")
		return err
	}

	channel := xmlquery.Find(doc, "//item")

	for _, n := range channel {
		v := entity.News{}

		if n := n.SelectElement("title"); n != nil {
			title := n.InnerText()
			v.Title = title
		}

		if n := n.SelectElement("link"); n != nil {
			link := n.InnerText()
			v.URL = link

			docDetail, err := htmlquery.LoadURL(link)
			if err != nil {
				log.Err(err).Msg("Error GetLatestNewsRSS")
				return err
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
				log.Err(err).Msg("Error GetLatestNewsRSS")
				return err
			}

			v.PublishedDate = sql.NullTime{Time: timePub, Valid: true}

			timestamp := timePub.Unix()
			v.ArticleTS = int64(timestamp)
		}

		v.Inserted = sql.NullTime{Time: time.Now(), Valid: true}

		status, err := nsi.newsRepository.GetNewsByUrl(ctx, v.URL)
		if err != nil {
			log.Err(err).Msg("Error GetLatestNewsRSS")
			return err
		}

		log.Info().Msgf("URL RSS: %s, Status:  %t", v.URL, status)

		if !status {
			_, err := nsi.newsRepository.CreateNews(ctx, v)
			if err != nil {
				log.Err(err).Msg("Error GetLatestNewsRSS")
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
		log.Err(err).Msg("Error GetLatestNewsIndex")
		return err
	}

	list := htmlquery.FindOne(doc, "//ul[@class = 'list-news indeks-new']")
	data := htmlquery.Find(list, "//div[@class = 'col-sm-4']//a")
	for _, n := range data {
		v := entity.News{}

		span := htmlquery.FindOne(n, "//span") // premium content always contains span tag
		if span == nil {
			link := htmlquery.SelectAttr(n, "href")
			v.URL = link

			title := htmlquery.SelectAttr(n, "title")
			v.Title = title

			status, err := nsi.newsRepository.GetNewsByUrl(ctx, v.URL)
			if err != nil {
				log.Err(err).Msg("Error GetLatestNewsIndex")
				return err
			}

			log.Info().Msgf("URL IDX: %s, Status: %t", link, status)

			if !status {
				docContent, err := htmlquery.LoadURL(link)
				if err != nil {
					log.Err(err).Msg("Error GetLatestNewsIndex")
					return err
				}

				docDataDetail := htmlquery.FindOne(docContent, "//article")
				if docDataDetail == nil {
					log.Err(errors.New("Tag article missing")).Msg("Error GetLatestNewsIndex")
				} else {
					strDocDataDetail := htmlquery.InnerText(docDataDetail)
					if strDocDataDetail == "" {
						log.Err(errors.New("Doc data detail missing")).Msg("Error GetLatestNewsIndex")
					} else {
						strDocDataDetail = regexp.MustCompile(`\s+`).ReplaceAllString(strDocDataDetail, " ")
						strDocDataDetail = strings.ReplaceAll(strDocDataDetail, GOOGLE_TAG, "")
						strDocDataDetail = strings.ReplaceAll(strDocDataDetail, CEK_BERITA, "")
						v.Content = strDocDataDetail

						docDataDate := htmlquery.Find(docContent, "//meta[@name]")
						for _, n := range docDataDate {
							metaName := htmlquery.SelectAttr(n, "name")
							if metaName == "publishdate" {
								strPubDate := htmlquery.SelectAttr(n, "content")
								pubDate, err := time.Parse("2006/01/02 15:04:05", strPubDate)
								if err != nil {
									log.Err(err).Msg("Error GetLatestNewsIndex")
									return err
								}

								v.PublishedDate = sql.NullTime{Time: pubDate, Valid: true}

								timestamp := pubDate.Unix()
								v.ArticleTS = int64(timestamp)

								v.Inserted = sql.NullTime{Time: time.Now(), Valid: true}
							}
						}

						_, err = nsi.newsRepository.CreateNews(ctx, v)
						if err != nil {
							log.Err(err).Msg("Error GetLatestNewsIndex")
							return err
						}

						counter++
					}
				}

			}
		}
	}

	duration := time.Since(start)

	log.Info().Msgf("%d data added successfully in %f seconds", counter, duration.Seconds())

	return nil
}

func (nsi *NewsServiceImpl) FindNewsByID(ctx context.Context, newsID int64) (dto.NewsRespBody, error) {
	newsResp := dto.NewsRespBody{}

	news, err := nsi.newsRepository.FindNewsByID(ctx, newsID)
	if err != nil {
		log.Err(err).Msg("Error FindNewsByID")

		return newsResp, errs.FindErrorType(err)
	}

	newsResp = dto.CreateNewsResp(news)

	return newsResp, nil
}
