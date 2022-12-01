package services

import (
	"fp-be-glng-h8/models/entity"
	"fp-be-glng-h8/models/web"
	"fp-be-glng-h8/repositories"
	"fp-be-glng-h8/responses"
)

type UserService interface {
	Register(userInput web.CreateUserRequest) (web.CreateUserResponse, error)
}

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (s *UserServiceImpl) Register(userInput web.CreateUserRequest) (web.CreateUserResponse, error) {
	user := entity.User{
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: userInput.Password,
		Age:      userInput.Age,
	}

	newUser, err := s.UserRepository.Create(user)

	return responses.ConvertCreateUserResponse(newUser), err

}
