package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/config"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/database"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/encryption"
)

type Token struct {
	Type         string `json:"type"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
	Expired      int64  `json:"expired"`
}

type JwtToken interface {
	Sign(claims jwt.MapClaims) Token
	SignRSA(claims jwt.MapClaims) Token
}

type JwtTokenImpl struct {
	cached                 *database.ScribleImpl
	jwtTokenTimeExp        time.Duration
	jwtRefreshTokenTimeExp time.Duration
}

func NewJwt(cached *database.ScribleImpl) *JwtTokenImpl {
	jwtTokenDuration, err := time.ParseDuration(config.Get().Auth.JWTToken.Expired)
	if err != nil {
		log.Err(err).Msg(config.Get().Auth.JWTToken.Expired)
	}

	jwtRefreshDuration, err := time.ParseDuration(config.Get().Auth.JWTToken.RefreshExpired)
	if err != nil {
		log.Err(err).Msg(config.Get().Auth.JWTToken.RefreshExpired)
	}

	return &JwtTokenImpl{
		cached:                 cached,
		jwtTokenTimeExp:        jwtTokenDuration,
		jwtRefreshTokenTimeExp: jwtRefreshDuration,
	}
}

func (o JwtTokenImpl) Sign(claims jwt.MapClaims) Token {
	timeNow := time.Now()
	tokenExpired := timeNow.Add(o.jwtTokenTimeExp).Unix()

	if claims["id"] == nil {
		return Token{}
	}

	token := jwt.New(jwt.SigningMethodHS256)
	// setup userdata
	var _, checkExp = claims["exp"]
	var _, checkIat = claims["exp"]

	// if user didn't define claims expired
	if !checkExp {
		claims["exp"] = tokenExpired
	}

	// if user didn't define claims iat
	if !checkIat {
		claims["iat"] = timeNow.Unix()
	}

	claims["token_type"] = "access_token"
	token.Claims = claims
	authToken := new(Token)

	tokenString, err := token.SignedString([]byte(config.Get().Application.Key.Default))
	if err != nil {
		log.Err(err)
		return Token{}
	}

	authToken.Token = tokenString
	authToken.Type = "Bearer"

	//create refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenExpired := timeNow.Add(o.jwtRefreshTokenTimeExp).Unix()

	claims["exp"] = refreshTokenExpired
	claims["token_type"] = "refresh_token"
	refreshToken.Claims = claims

	refreshTokenString, err := refreshToken.SignedString([]byte(config.Get().Application.Key.Default))
	if err != nil {
		return Token{}
	}

	authToken.RefreshToken = refreshTokenString

	//save token to redis db
	go func() {
		encryptedRefreshToken, err := encryption.AesCFBEncryption(refreshTokenString, config.Get().Application.Key.Default)
		if err != nil {
			log.Err(err)
		}

		// check data type of the claims
		switch claims["id"].(type) {
		case int:
			claims["id"] = fmt.Sprintf("%d", claims["id"].(int))
		case int32:
			claims["id"] = fmt.Sprintf("%d", claims["id"].(int32))
		case float64:
			claims["id"] = fmt.Sprintf("%d", int(claims["id"].(float64)))
		default:
		}

		o.cached.DB().Write("refresh_token", claims["id"].(string), RefreshToken{RefreshToken: encryptedRefreshToken, Expired: refreshTokenExpired})
		if err != nil {
			log.Err(err).Msgf("Failed to save refresh token to scrible")
		} else {
			log.Info().Msg("Successfully to save refresh token to scrible")
		}
	}()

	return Token{
		Type:         config.Get().Auth.JWTToken.Type,
		Token:        authToken.Token,
		RefreshToken: authToken.RefreshToken,
	}
}

func (o JwtTokenImpl) SignRSA(claims jwt.MapClaims) Token {
	timeNow := time.Now()
	tokenExpired := timeNow.Add(o.jwtTokenTimeExp).Unix()

	if claims["id"] == nil {
		return Token{}
	}

	token := jwt.New(jwt.SigningMethodRS256)
	// setup userdata
	var _, checkExp = claims["exp"]
	var _, checkIat = claims["exp"]

	// if user didn't define claims expired
	if !checkExp {
		claims["exp"] = tokenExpired
	}
	// if user didn't define claims iat
	if !checkIat {
		claims["iat"] = timeNow.Unix()
	}
	claims["token_type"] = "access_token"

	token.Claims = claims
	authToken := new(Token)
	privateRsa, err := encryption.ReadPrivateKeyFromEnv(config.Get().Application.Key.RSA.Private)
	if err != nil {
		log.Err(err).Msg("err read private key rsa from env")
		return Token{}
	}
	tokenString, err := token.SignedString(privateRsa)
	if err != nil {
		log.Err(err).Msg("err read private rsa")
		return Token{}
	}

	authToken.Token = tokenString
	authToken.Type = "Bearer"

	//create refresh token
	refreshToken := jwt.New(jwt.SigningMethodRS256)
	refreshTokenExpired := timeNow.Add(o.jwtRefreshTokenTimeExp).Unix()

	claims["exp"] = refreshTokenExpired
	claims["token_type"] = "refresh_token"
	refreshToken.Claims = claims
	refreshTokenString, err := refreshToken.SignedString(privateRsa)
	if err != nil {
		log.Err(err).Msg("")
		return Token{}
	}
	authToken.RefreshToken = refreshTokenString

	//save token to redis db
	go func() {
		encryptedRefreshToken, err := encryption.AesCFBEncryption(refreshTokenString, config.Get().Application.Key.Default)
		if err != nil {
			log.Err(err)
		}
		// check data type of the claims
		switch claims["id"].(type) {
		case int:
			claims["id"] = fmt.Sprintf("%d", claims["id"].(int))
		case int32:
			claims["id"] = fmt.Sprintf("%d", claims["id"].(int32))
		case float64:
			claims["id"] = fmt.Sprintf("%d", int(claims["id"].(float64)))
		default:
		}
		o.cached.DB().Write("refresh_token", claims["id"].(string), RefreshToken{RefreshToken: encryptedRefreshToken, Expired: refreshTokenExpired})
		if err != nil {
			log.Err(err).Msg("Failed to save refresh token to redis")
		} else {
			log.Info().Msg("Successfully to save refresh token to redis")
		}
	}()

	return Token{
		Type:         "Bearer",
		Token:        authToken.Token,
		RefreshToken: authToken.RefreshToken,
	}
}
