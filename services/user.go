package services

import (
	"fp-be-glng-h8/models/entity"
	"fp-be-glng-h8/models/web"
	"fp-be-glng-h8/repositories"
	"fp-be-glng-h8/responses"
)

type UserService interface {
	Register(registerInput web.CreateUserRequest) (web.CreateUserResponse, error)
	Login(loginInput web.LoginUserRequest) (entity.User, error)
	Profile(userId uint) (web.CreateUserResponse, error)
}

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (s *UserServiceImpl) Register(registerInput web.CreateUserRequest) (web.CreateUserResponse, error) {
	user := entity.User{
		Username: registerInput.Username,
		Email:    registerInput.Email,
		Password: registerInput.Password,
		Age:      registerInput.Age,
	}

	newUser, err := s.UserRepository.Create(user)

	return responses.ConvertCreateUserResponse(newUser), err

}

func (s *UserServiceImpl) Login(loginInput web.LoginUserRequest) (entity.User, error) {
	inputUser := entity.User{
		Email:    loginInput.Email,
		Password: loginInput.Password,
	}
	user, err := s.UserRepository.FindUserByEmail(inputUser)
	return user, err
}

func (s *UserServiceImpl) Profile(userId uint) (web.CreateUserResponse, error) {
	user, err := s.UserRepository.FindUserByID(userId)
	userResp := responses.ConvertCreateUserResponse(user)
	return userResp, err
}
