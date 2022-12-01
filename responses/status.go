package responses

import "github.com/gin-gonic/gin"

func ConvertUserStatusResponse(ctx *gin.Context, code int, message string, data interface{}) {
	switch code != 0 {
	case code == 201:
		ctx.JSON(code, gin.H{
			"message": message,
			"user":    data,
		})
	case code == 200 && message == "Login Success":
		ctx.JSON(code, gin.H{
			"message": message,
			"token":   data,
		})
	case code == 200:
		ctx.JSON(code, gin.H{
			"message": message,
			"user":    data,
		})
	default:
		ctx.JSON(code, gin.H{
			"message": message,
		})
	}
}
