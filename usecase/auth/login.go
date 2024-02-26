package auth

import (
	"context"

	jwtPkg "base-gin-golang/pkg/jwt"
	"base-gin-golang/pkg/logger"
)

type LoginInputBody struct {
	Email    string `json:"email"    validate:"required,customEmail,email,max=255"`
	Password string `json:"password" validate:"required"`
}

type LoginInput struct {
	Body LoginInputBody
}

type LoginUser struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type LoginOutputBody struct {
	User         LoginUser `json:"user"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
}

type LoginOutput struct {
	Body LoginOutputBody
}

func (u *authUseCase) Login(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	user, err := u.userRepository.GetByEmail(ctx, input.Body.Email)
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	err = u.passwordService.CheckHashPassword(user.Password, input.Body.Password)
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	accessToken, err := u.jwtService.GenerateAccessToken(&jwtPkg.GenerateTokenInput{
		UserID: user.ID,
		Email:  user.Email,
	})
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	refreshToken, err := u.jwtService.GenerateRefreshToken(&jwtPkg.GenerateTokenInput{
		UserID: user.ID,
		Email:  user.Email,
	})
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	return &LoginOutput{
		Body: LoginOutputBody{
			LoginUser{
				ID:    user.ID,
				Email: user.Email,
				Name:  user.Name,
			},
			accessToken,
			refreshToken,
		},
	}, nil
}
