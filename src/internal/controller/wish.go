// controller/wishlist_controller.go
package controller

import (
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/nukahaha/car_store/src/internal/service"
	"net/http"
)

type WishlistController struct {
	wishlistService *service.WishlistService
}

func NewWishlistController(wishlistService *service.WishlistService) *WishlistController {
	return &WishlistController{wishlistService}
}

func (c *WishlistController) AddToWishlistHandler(ctx *gin.Context) {
	userID, err := GetUserID(ctx) // Замените на получение UserID из вашего контекста или запроса
	carID := ctx.Param("carID")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.wishlistService.AddToWishlist(userID, carID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Car added to wishlist successfully"})
}

func (c *WishlistController) RemoveFromWishlistHandler(ctx *gin.Context) {
	userID, err := GetUserID(ctx) // Замените на получение UserID из вашего контекста или запроса
	carID := ctx.Param("carID")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.wishlistService.RemoveFromWishlist(userID, carID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Car removed from wishlist successfully"})
}

func (c *WishlistController) GetWishlistHandler(ctx *gin.Context) {
	userID, err := GetUserID(ctx) // Замените на получение UserID из вашего контекста или запроса

	cars, err := c.wishlistService.GetCarsByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginview.HTML(ctx, http.StatusOK, "wishlist", gin.H{
		"title": "all cars",
		"Cars":  cars,
	})
}
