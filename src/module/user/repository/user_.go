package repository

import (
	"context"

	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/user/dto"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/user/entity"
)

func (n *UserRepositoryImpl) CreateUser(ctx context.Context, user dto.UserReqBody) (entity.User, error) {
	entityUser := entity.User{}

	return entityUser, nil
}
