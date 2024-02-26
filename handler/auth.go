package handler

import (
	"context"

	errors "base-gin-golang/pkg/errors/custom"
	"base-gin-golang/usecase/auth"
)

func (h *Handler) Login(
	ctx context.Context,
	input *auth.LoginInput,
) (*auth.LoginOutput, error) {
	if err := h.Validator.Struct(&input.Body); err != nil {
		errValidate := errors.NewValidateError(ctx, input, err)
		return nil, errValidate
	}
	output, err := h.App.AuthUseCase.Login(ctx, input)
	if err != nil {
		errConverted := h.App.ErrorService.ParseInternalServer(err)
		return nil, errConverted
	}
	return output, nil
}
