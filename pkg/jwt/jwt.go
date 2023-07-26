package jwt

import (
	"base-gin-golang/config"
	errorConstants "base-gin-golang/errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/wire"
)

type CustomJwtClaims struct {
	GenerateTokenInput
	jwt.StandardClaims
}

type GenerateTokenInput struct {
	UserID int
	Email  string
}

type Service interface {
	GenerateAccessToken(input *GenerateTokenInput) (string, error)
	GenerateRefreshToken(input *GenerateTokenInput) (string, error)
	ValidateAccessToken(tokenString string) (*CustomJwtClaims, error)
	ValidateRefreshToken(tokenString string) (*CustomJwtClaims, error)
}

type jwtService struct {
	cfg *config.Environment
}

var ProviderSet = wire.NewSet(NewJwtService)

func NewJwtService(cfg *config.Environment) Service {
	return &jwtService{
		cfg: cfg,
	}
}

func (s *jwtService) GenerateAccessToken(input *GenerateTokenInput) (string, error) {
	claims := &CustomJwtClaims{
		*input,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(s.cfg.AccessTokenExpireMinute) * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.cfg.AccessTokenSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *jwtService) GenerateRefreshToken(input *GenerateTokenInput) (string, error) {
	claims := &CustomJwtClaims{
		*input,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(s.cfg.RefreshTokenExpireHour) * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.cfg.RefreshTokenSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *jwtService) ValidateAccessToken(tokenString string) (*CustomJwtClaims, error) {
	claims := &CustomJwtClaims{}
	jwtKey := []byte(s.cfg.AccessTokenSecretKey)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		v, ok := err.(*jwt.ValidationError)
		if !ok {
			return nil, errorConstants.ErrTokenInvalid
		}
		if v.Errors == jwt.ValidationErrorExpired {
			return nil, errorConstants.ErrTokenExpired
		}
		return nil, errorConstants.ErrTokenInvalid
	}
	if !token.Valid {
		return nil, errorConstants.ErrTokenInvalid
	}
	return claims, nil
}

func (s *jwtService) ValidateRefreshToken(tokenString string) (*CustomJwtClaims, error) {
	claims := &CustomJwtClaims{}
	jwtKey := []byte(s.cfg.RefreshTokenSecretKey)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		v, ok := err.(*jwt.ValidationError)
		if !ok {
			return nil, errorConstants.ErrTokenInvalid
		}
		if v.Errors == jwt.ValidationErrorExpired {
			return nil, errorConstants.ErrTokenExpired
		}
		return nil, errorConstants.ErrTokenInvalid
	}
	if !token.Valid {
		return nil, errorConstants.ErrTokenInvalid
	}
	return claims, nil
}
