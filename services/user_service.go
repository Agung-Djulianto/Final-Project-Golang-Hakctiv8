package services

import (
	"Project-Akhir/helpers"
	"Project-Akhir/models"
	"Project-Akhir/repositori"
)

type UserService struct {
	UserRepository repositori.IUserRepository
}

func NewUserService(UserRepository repositori.IUserRepository) *UserService {
	return &UserService{
		UserRepository: UserRepository,
	}
}

func (us *UserService) Add(request models.UserRegisterRequest) (models.UserRegisterResponse, error) {

	id := helpers.GenerateID()
	hashed_password, err := helpers.HashPassword(request.Password)

	if err != nil {
		return models.UserRegisterResponse{}, err
	}

	newUser := models.User{
		ID:       id,
		Email:    request.Email,
		UserName: request.Username,
		Password: hashed_password,
		Age:      request.Age,
	}

	res, err := us.UserRepository.Save(newUser)
	if err != nil {
		return models.UserRegisterResponse{}, err
	}

	return models.UserRegisterResponse{
		ID:        res.ID,
		UserName:  res.UserName,
		CreatedAt: res.CreatedAt,
	}, nil
}

func (us *UserService) Login(request models.UserLoginRequest) (models.UserLoginResponse, error) {

	result, err := us.UserRepository.GetByUsername(request.Username)

	if err != nil {
		return models.UserLoginResponse{}, err
	}

	valid := helpers.CekPasswordHash(request.Password, result.Password)

	if !valid {
		return models.UserLoginResponse{}, models.ErrorInvalidEmailOrPassword
	}

	token, err := helpers.GenerateToken(result.ID)
	if err != nil {
		return models.UserLoginResponse{}, models.ErrorInvalidToken
	}

	return models.UserLoginResponse{
		Token: token,
	}, nil
}
