package service

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/nukahaha/car_store/src/internal/model"
	"github.com/nukahaha/car_store/src/internal/model/request"
	"github.com/nukahaha/car_store/src/internal/repository"
)

type AuthService struct {
	UserRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (as *AuthService) Login(loginRequest *request.LoginRequest) error {
	if strings.Trim(loginRequest.Email, " ") == "" || strings.Trim(loginRequest.Password, " ") == "" {
		return errors.New("username and password cannot be empty")
	}

	user, err := as.UserRepository.GetByFieldMail(loginRequest.Email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return err
	}

	if loginRequest.Email == user.Mail {
		return nil
	}

	return errors.New("error occurred during login")
}

func (as *AuthService) Register(registerRequest *request.RegisterRequest) error {
	if strings.Trim(registerRequest.Email, " ") == "" ||
		strings.Trim(registerRequest.Password, " ") == "" ||
		strings.Trim(registerRequest.ConfirmPassword, " ") == "" {
		return errors.New("required fields are empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = as.UserRepository.Register(&model.User{
		Mail:     registerRequest.Email,
		Password: string(hashedPassword),
		Name:     registerRequest.Name,
		Surname:  registerRequest.Surname,
		Birthday: registerRequest.Birthday,
	})
	if err != nil {
		return errors.New("user couldn't be saved into database")
	}

	return nil
}