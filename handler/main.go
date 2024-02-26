package handler

import (
	"base-gin-golang/cmd/wire"
	"base-gin-golang/config"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	Config    *config.Environment
	App       *wire.App
	Validator *validator.Validate
}

func NewHandler(cfg *config.Environment, app *wire.App, validator *validator.Validate) *Handler {
	return &Handler{
		Config:    cfg,
		App:       app,
		Validator: validator,
	}
}
