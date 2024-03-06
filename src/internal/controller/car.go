package controller

import (
	"net/http"
	"strconv"

	"github.com/foolin/goview/supports/ginview"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/nukahaha/car_store/src/internal/model"
	"github.com/nukahaha/car_store/src/internal/service"

	"github.com/nukahaha/car_store/src/internal/model/request"
)

type CarController struct {
	carService  *service.CarService
	userService *service.AuthService
}

func NewCarController(carService *service.CarService, userService *service.AuthService) *CarController {
	return &CarController{
		carService:  carService,
		userService: userService,
	}
}

func (cc *CarController) GetAllCars(ctx *gin.Context) {
	cars, err := cc.carService.GetAllCars()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching cars"})
		return
	}

	ginview.HTML(ctx, http.StatusOK, "cars", gin.H{
		"title": "all cars",
		"Cars":  cars,
	})
}

func (cc *CarController) PostCarPage(ctx *gin.Context) {
	cars, err := cc.carService.GetAllCars()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching cars"})
		return
	}

	ginview.HTML(ctx, http.StatusOK, "post", gin.H{
		"title": "Posting cars",
		"Cars":  cars,
	})
}

func (cc *CarController) PostCar(ctx *gin.Context) {

	carRequest := &request.CarRequest{
		Model_Car:   ctx.PostForm("model"),
		Year:        ctx.GetInt("year"),
		Cost:        ctx.GetFloat64("cost"),
		Description: ctx.PostForm("description"),
		Image:       ctx.PostForm("image"),
		Brand:       ctx.PostForm("brand"),
	}

	ctx.MustBindWith(carRequest, binding.FormPost)

	var err error

	carRequest.UserID, err = GetUserID(ctx)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	carRequest.Year, err = strconv.Atoi(ctx.PostForm("year"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year", "details": err.Error()})
		return
	}

	carRequest.Cost, err = strconv.ParseFloat(ctx.PostForm("cost"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cost", "details": err.Error()})
		return
	}

	err = cc.carService.AddCar(carRequest)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/cars?hasError=true&errorMessage="+err.Error())

		return
	}

	ctx.Redirect(http.StatusFound, "/cars")
}

func (cc *CarController) UpdateCar(ctx *gin.Context) {

	// Получаем значения year и cost
	year, err := strconv.Atoi(ctx.PostForm("year"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year"})
		return
	}

	cost, err := strconv.ParseFloat(ctx.PostForm("cost"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cost"})
		return
	}

	carID := ctx.Param("id")
	updatedCarRequest := &request.CarRequest{
		Model_Car:   ctx.PostForm("model"),
		Year:        year,
		Cost:        cost,
		Description: ctx.PostForm("description"),
		Image:       ctx.PostForm("image"),
		Brand:       ctx.PostForm("brand"),
	}
	ctx.MustBindWith(updatedCarRequest, binding.FormPost)

	updatedCar := &model.Car{
		Model_Car:   updatedCarRequest.Model_Car,
		Year:        updatedCarRequest.Year,
		Cost:        updatedCarRequest.Cost,
		Description: updatedCarRequest.Description,
		Image:       updatedCarRequest.Image,
		Brand:       updatedCarRequest.Brand,
	}

	// Преобразуем carID в ObjectID
	objectID, err := primitive.ObjectIDFromHex(carID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}

	// Устанавливаем ObjectID для обновления
	updatedCar.ID = objectID

	err = cc.carService.UpdateCar(updatedCar)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating car"})
		return
	}

	ctx.Redirect(http.StatusFound, "/cars")
}

func (cc *CarController) DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("id")
	err := cc.carService.DeleteCar(carID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting car"})
		return
	}

	ctx.Redirect(http.StatusFound, "/cars")
}
func (cc *CarController) GetCarByID(ctx *gin.Context) {
	carID := ctx.Param("id")
	car, err := cc.carService.GetCarByID(carID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching car"})
		return
	}

	ginview.HTML(ctx, http.StatusOK, "edit_car", gin.H{
		"title": "Car Details",
		"Car":   car,
	})
}

func (cc *CarController) CarsPageHandler(ctx *gin.Context) {
	email, err := GetUserEmail(ctx) // Замените на получение Email из вашего контекста или запроса
	cars, err := cc.carService.GetCarsByUserEmail(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := cc.userService.GetByFieldMail(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Отправьте данные в ваш HTML-шаблон
	ctx.HTML(http.StatusOK, "your_page", gin.H{
		"title": "Cars Page",
		"User":  user,
		"Cars":  cars,
	})
}
