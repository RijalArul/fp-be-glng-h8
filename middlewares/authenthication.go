package middlewares

import (
	"fp-be-glng-h8/exceptions"
	"fp-be-glng-h8/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenthication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)
		_ = verifyToken

		if err != nil {
			exceptions.Errors(ctx, http.StatusUnauthorized, "Unauthenthicated", "Unauthenthicated")
			return
		}
		ctx.Set("userData", verifyToken)
		ctx.Next()

	}
}
