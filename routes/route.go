package routes

import (
	"fp-be-glng-h8/configs"
	"fp-be-glng-h8/handlers"
	"fp-be-glng-h8/middlewares"
	"fp-be-glng-h8/repositories"
	"fp-be-glng-h8/services"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()
	db := configs.GetDB()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	userRouter := r.Group("/users")

	{
		userRouter.POST("/register", userHandler.Register)
		userRouter.POST("/login", userHandler.Login)
		userRouter.Use(middlewares.Authenthication())
		userRouter.GET("/profile", userHandler.Profile)
		userRouter.PUT("/", userHandler.UpdateUser)
		userRouter.DELETE("/", userHandler.DeleteUser)

	}

	r.Run()
}
