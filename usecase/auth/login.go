package auth

import (
	jwtPkg "base-gin-golang/pkg/jwt"
	"base-gin-golang/pkg/logger"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required,email,customEmail,max=255"`
	Password string `json:"password" binding:"required"`
}

type LoginUser struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type LoginOutput struct {
	User         LoginUser `json:"user"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
}

func (au *authUseCase) Login(ctx *gin.Context, input *LoginInput) (*LoginOutput, error) {
	user, err := au.userRepository.GetByEmail(ctx, input.Email)
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	err = au.passwordService.CheckHashPassword(user.Password, input.Password)
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	accessToken, err := au.jwtService.GenerateAccessToken(&jwtPkg.GenerateTokenInput{
		UserID: user.ID,
		Email:  user.Email,
	})
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	refreshToken, err := au.jwtService.GenerateRefreshToken(&jwtPkg.GenerateTokenInput{
		UserID: user.ID,
		Email:  user.Email,
	})
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	return &LoginOutput{
		LoginUser{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
		accessToken,
		refreshToken,
	}, nil
}
