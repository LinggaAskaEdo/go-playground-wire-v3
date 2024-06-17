package service

import (
	"context"

	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/user/dto"
)

func (n *UserServiceImpl) CreateUser(ctx context.Context, user dto.UserReqBody) (dto.UserRespBody, error) {
	userResp := dto.UserRespBody{}

	return userResp, nil
}
