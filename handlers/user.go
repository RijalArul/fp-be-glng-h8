package handlers

import (
	"fp-be-glng-h8/exceptions"
	"fp-be-glng-h8/helpers"
	"fp-be-glng-h8/models/web"
	"fp-be-glng-h8/responses"
	"fp-be-glng-h8/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(ctx *gin.Context)
}

type UserHandlerImpl struct {
	UserService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &UserHandlerImpl{UserService: userService}
}

var (
	appJSON = "application/json"
)

func (h *UserHandlerImpl) Register(ctx *gin.Context) {
	var userInput web.CreateUserRequest
	contentType := helpers.GetContentType(ctx)

	if contentType == appJSON {
		ctx.ShouldBindJSON(&userInput)
	} else {
		ctx.ShouldBind(&userInput)
	}

	userResp, err := h.UserService.Register(userInput)

	if err != nil {
		exceptions.Errors(ctx, http.StatusBadRequest, "Failed Register User", err.Error())
		return
	}

	responses.ConvertUserStatusResponse(ctx, http.StatusCreated, "Success Register User", userResp)

}
