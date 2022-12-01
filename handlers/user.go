package handlers

import (
	"fp-be-glng-h8/exceptions"
	"fp-be-glng-h8/helpers"
	"fp-be-glng-h8/models/web"
	"fp-be-glng-h8/responses"
	"fp-be-glng-h8/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Profile(ctx *gin.Context)
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

func (h *UserHandlerImpl) Login(ctx *gin.Context) {
	var LoginInput web.LoginUserRequest
	contentType := helpers.GetContentType(ctx)

	if contentType == appJSON {
		ctx.ShouldBindJSON(&LoginInput)
	} else {
		ctx.ShouldBind(&LoginInput)
	}

	user, err := h.UserService.Login(LoginInput)

	if err != nil {
		exceptions.Errors(ctx, http.StatusNotFound, "User Not Found", err.Error())
		return
	}

	validPass := helpers.ComparePass([]byte(user.Password), []byte(LoginInput.Password))
	if !validPass {
		exceptions.Errors(ctx, http.StatusUnauthorized, "Password Failed", "Unauthenthicated")
		return
	}

	genToken := helpers.GenerateToken(user.ID, user.Email)

	responses.ConvertUserStatusResponse(ctx, http.StatusOK, "Login Success", genToken)
}

func (h *UserHandlerImpl) Profile(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	user, err := h.UserService.Profile(userId)

	if err != nil {
		exceptions.Errors(ctx, http.StatusNotFound, "User Not Found", err.Error())
		return
	}

	responses.ConvertUserStatusResponse(ctx, http.StatusOK, "Success found user", user)

}
