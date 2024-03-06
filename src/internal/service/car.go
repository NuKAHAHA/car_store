package service

import (
	"errors"

	"strings"

	"github.com/nukahaha/car_store/src/internal/model"
	"github.com/nukahaha/car_store/src/internal/model/request"
	"github.com/nukahaha/car_store/src/internal/repository"
)

type CarService struct {
	CarRepository *repository.CarRepository
}

func NewCarService(carRepository *repository.CarRepository) *CarService {
	return &CarService{
		CarRepository: carRepository,
	}
}

func (cs *CarService) GetAllCars() ([]model.Car, error) {
	return cs.CarRepository.GetAllCars()
}

func (cs *CarService) AddCar(carRequest *request.CarRequest) error {
	if strings.Trim(carRequest.Model_Car, " ") == "" ||
		strings.Trim(carRequest.Description, " ") == "" ||
		strings.Trim(carRequest.Brand, " ") == "" {
		return errors.New("required fields are empty")
	}

	err := cs.CarRepository.AddCar(&model.Car{
		UserID:      carRequest.UserID,
		Model_Car:   carRequest.Model_Car,
		Year:        carRequest.Year,
		Cost:        carRequest.Cost,
		Description: carRequest.Description,
		Image:       carRequest.Image,
		Brand:       carRequest.Brand,
	})
	if err != nil {
		return errors.New("user couldn't be saved into database")
	}

	return nil
}

// В CarService

func (cs *CarService) GetCarByID(carID string) (*model.Car, error) {
	return cs.CarRepository.GetCarByID(carID)
}

func (cs *CarService) UpdateCar(car *model.Car) error {
	return cs.CarRepository.UpdateCar(car)
}

func (cs *CarService) DeleteCar(carID string) error {
	return cs.CarRepository.DeleteCar(carID)
}
