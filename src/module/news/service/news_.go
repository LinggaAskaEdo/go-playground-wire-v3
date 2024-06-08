package service

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/errors"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/dto"
)

func (n *NewsServiceImpl) FindNewsByID(ctx context.Context, newsID int64) (dto.NewsRespBody, error) {
	newsResp := dto.NewsRespBody{}

	news, err := n.newsRepository.FindNewsByID(ctx, newsID)
	if err != nil {
		log.Err(err).Msg("Error fetch news data")

		return newsResp, errors.FindErrorType(err)
	}

	newsResp = dto.CreateNewsResp(news)

	return newsResp, nil
}
