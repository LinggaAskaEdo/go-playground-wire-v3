package service

import (
	"context"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/auth"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/user/dto"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/user/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user dto.UserReqBody) (dto.UserRespBody, error)
}

type UserServiceImpl struct {
	jwtAuth        auth.JwtToken
	userRepository repository.UserRepository
}

func NewUserService(
	jwtAuth auth.JwtToken,
	userRepository repository.UserRepository,
) *UserServiceImpl {
	return &UserServiceImpl{
		jwtAuth:        jwtAuth,
		userRepository: userRepository,
	}
}
