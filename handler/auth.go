package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"base-gin-golang/usecase/auth"
)

func Login(ctx *gin.Context, authUseCase auth.UseCase) {
	var input auth.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	output, err := authUseCase.Login(ctx, &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, output)
}
