package repository

import (
	"context"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/database"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/user/dto"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/user/entity"

)

type UserRepository interface {
	CreateUser(ctx context.Context, user dto.UserReqBody) (entity.User, error)
}

type UserRepositoryImpl struct {
	db0 *database.MysqlImpl
}

func NewUserRepository(
	db0 *database.MysqlImpl,
) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db0: db0,
	}
}
