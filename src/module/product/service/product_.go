package service

import (
	"context"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/errors"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/product/dto"
	"github.com/rs/zerolog/log"
)

func (n *ProductServiceImpl) FindProductByID(ctx context.Context, productID int64) (dto.ProductRespBody, error) {
	productResp := dto.ProductRespBody{}

	product, err := n.productRepository.FindProductByID(ctx, productID)
	if err != nil {
		log.Err(err).Msg("Error fetch product data")

		return productResp, errors.FindErrorType(err)
	}

	productResp = dto.CreateProductResp(product)

	return productResp, nil
}
