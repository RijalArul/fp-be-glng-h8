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
	Update(updateInput web.UpdateUserRequest, userId uint) (web.UpdateUserResponse, error)
	DeleteUser(userId uint) error
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

func (s *UserServiceImpl) Update(updateInput web.UpdateUserRequest, userId uint) (web.UpdateUserResponse, error) {
	user := entity.User{
		Username: updateInput.Username,
		Email:    updateInput.Email,
	}

	updateUser, err := s.UserRepository.UpdateUser(user, userId)
	userResp := responses.ConvertUpdateUserResponse(updateUser)
	return userResp, err
}

func (s *UserServiceImpl) DeleteUser(userId uint) error {
	err := s.UserRepository.DeleteUser(userId)
	return err
}
