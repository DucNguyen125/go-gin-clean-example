package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var ExcludingApisForAuth = map[string]bool{
	"/api/ping":          true,
	"/api/v1/auth/login": true,
	"/docs":              true,
	"/openapi.json":      true,
	"/openapi.yaml":      true,
}

func (m *middleware) Authentication(ctx *gin.Context) {
	if ExcludingApisForAuth[ctx.Request.URL.Path] {
		return
	}
	headerAuthorization := ctx.GetHeader("Authorization")
	var token string
	_, err := m.stringService.Sscanf(headerAuthorization, "Bearer %s", &token)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if token == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	verifiedToken, err := m.jwtService.ValidateAccessToken(token)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.Set("userID", verifiedToken.UserID)
}
