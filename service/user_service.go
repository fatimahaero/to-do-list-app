package service

import (
	"errors"
	"to-do-list-app/domain"
	"to-do-list-app/dto"
	"to-do-list-app/repository"
	"to-do-list-app/utils"
)

type UserService interface {
	RegisterUser(req dto.UserCreateRequest) (*dto.UserResponse, error)
	LoginUser(req dto.UserLoginRequest) (*dto.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) RegisterUser(req dto.UserCreateRequest) (*dto.UserResponse, error) {
	existingUser, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username already taken")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := domain.User{
		Username: req.Username,
		Password: hashedPassword,
	}

	savedUser, err := s.userRepo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	response := dto.UserResponse{
		ID:       savedUser.ID,
		Username: savedUser.Username,
	}

	return &response, nil
}

func (s *userService) LoginUser(req dto.UserLoginRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil || user == nil {
		return nil, errors.New("invalid username or password")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid username or password")
	}

	response := dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	return &response, nil
}
