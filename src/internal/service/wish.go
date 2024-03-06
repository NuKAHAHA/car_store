// service/wishlist_service.go
package service

import (
	"github.com/nukahaha/car_store/src/internal/model"
	"github.com/nukahaha/car_store/src/internal/repository"
)

type WishlistService struct {
	wishlistRepo *repository.WishlistRepository
	carRepo      *repository.CarRepository
}

func NewWishlistService(wishlistRepo *repository.WishlistRepository, carRepo *repository.CarRepository) *WishlistService {
	return &WishlistService{wishlistRepo, carRepo}
}

func (s *WishlistService) AddToWishlist(userID, carID string) error {
	wishlist := &model.Wishlist{UserID: userID, CarID: carID}
	return s.wishlistRepo.AddToWishlist(wishlist)
}

func (s *WishlistService) RemoveFromWishlist(userID, carID string) error {
	return s.wishlistRepo.RemoveFromWishlist(userID, carID)
}

func (s *WishlistService) GetCarsByUserID(userID string) ([]model.Car, error) {
	wishlist, err := s.wishlistRepo.GetWishlistByUserID(userID)
	if err != nil {
		return nil, err
	}

	var cars []model.Car
	for _, item := range wishlist {
		car, err := s.carRepo.GetCarByID(item.CarID)
		if err != nil {
			return nil, err
		}
		cars = append(cars, *car)
	}

	return cars, nil
}
