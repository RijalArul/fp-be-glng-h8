package exceptions

import (
	"github.com/gin-gonic/gin"
)

func Errors(ctx *gin.Context, code int, message string, err interface{}) {
	switch code != 0 {
	case code == 400:
		ctx.AbortWithStatusJSON(code, gin.H{
			"message": message,
			"err":     err,
		})

	case code == 401:
		ctx.AbortWithStatusJSON(code, gin.H{
			"message": message,
			"err":     err,
		})
	case code == 404:
		ctx.AbortWithStatusJSON(code, gin.H{
			"message": message,
			"err":     err,
		})
	default:
		ctx.AbortWithStatusJSON(code, gin.H{
			"message": "Internal Server Error",
		})
	}
}
