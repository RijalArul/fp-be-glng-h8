package responses

import (
	"fp-be-glng-h8/models/entity"
	"fp-be-glng-h8/models/web"
)

func ConvertCreateUserResponse(user entity.User) web.CreateUserResponse {
	return web.CreateUserResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}
}

func ConvertLoginUserResponse(token string) web.LoginUserResponse {
	return web.LoginUserResponse{
		Token: token,
	}
}
