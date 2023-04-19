package middleware

import (
	"Project-Akhir/helpers"
	"Project-Akhir/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	token := strings.Split(auth, " ")[1]

	jwtToken, err := helpers.VerifyToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	ctx.Set("user_id", claims["user_id"])

	ctx.Next()
}
