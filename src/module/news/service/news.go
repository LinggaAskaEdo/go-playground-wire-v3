package service

import (
	"context"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/auth"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/dto"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/repository"
)

type NewsService interface {
	GetLatestNewsRSS(ctx context.Context) error
	GetLatestNewsIndex(ctx context.Context) error
	FindNewsByID(ctx context.Context, newsID int64) (dto.NewsRespBody, error)
}

type NewsServiceImpl struct {
	jwtAuth        auth.JwtToken
	newsRepository repository.NewsRepository
}

func NewNewsService(
	jwtAuth auth.JwtToken,
	newsRepository repository.NewsRepository,
) *NewsServiceImpl {
	return &NewsServiceImpl{
		jwtAuth:        jwtAuth,
		newsRepository: newsRepository,
	}
}
