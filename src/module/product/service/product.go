package service

import (
	"context"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/auth"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/product/dto"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/product/repository"
)

type ProductService interface {
	FindProductByID(ctx context.Context, productID int64) (dto.ProductRespBody, error)
}

type ProductServiceImpl struct {
	jwtAuth           auth.JwtToken
	productRepository repository.ProductRepository
}

func NewProductService(
	jwtAuth auth.JwtToken,
	productRepository repository.ProductRepository,
) *ProductServiceImpl {
	return &ProductServiceImpl{
		jwtAuth:           jwtAuth,
		productRepository: productRepository,
	}
}
