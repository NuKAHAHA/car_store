package main

import (
	"context"
	"log"
	"time"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nukahaha/car_store/src/internal/configuration"
	"github.com/nukahaha/car_store/src/internal/controller"
	"github.com/nukahaha/car_store/src/internal/middleware"
	"github.com/nukahaha/car_store/src/internal/repository"
	"github.com/nukahaha/car_store/src/internal/service"
)

func main() {
	appConfiguration, err := configuration.NewConfiguration()
	if err != nil {
		log.Fatalf("Configuration error:\n%s", err.Error())
	}

	engine := gin.New()

	engine.HTMLRender = ginview.New(goview.Config{
		Root:      "src/views",
		Extension: ".html",
		Master:    "layouts/main",
	})

	engine.Static("/public", "./public")

	database, err := setDatabaseConnection(appConfiguration)
	if err != nil {
		log.Fatalf("Error occurred during database connection;\n%s", err.Error())
	}

	err = initAPI(appConfiguration, engine, database)
	if err != nil {
		log.Fatal("Couldn't init routes, middlewares, or other systems;\n", err.Error())
	}

	err = engine.Run()
	if err != nil {
		log.Fatal("Error occurred during server startup;\n", err.Error())
	}

	err = database.Close()
	if err != nil {
		log.Fatal("Error occurred during database exit;\n", err.Error())
	}
}

func initAPI(appConfiguration *configuration.Configuration, engine *gin.Engine, database *repository.Database) error {
	client, err := setMongoDBConnection(appConfiguration)
	if err != nil {
		return err
	}

	// Create Repositories
	userRepository := repository.NewUserRepository(client, "users")
	wishlistRepository := repository.NewWishlistRepository(client, "wishlists")
	carRepository := repository.NewCarRepository(client, "car")

	// Create Services
	authService := service.NewAuthService(userRepository)
	wishlistService := service.NewWishlistService(wishlistRepository, carRepository)
	carService := service.NewCarService(carRepository, userRepository)

	// Create Controllers
	homeController := controller.NewHomeController()
	authController := controller.NewAuthController(authService)
	wishlistController := controller.NewWishlistController(wishlistService)
	carController := controller.NewCarController(carService, authService)

	// Middlewares
	engine.Use(setSession(appConfiguration))

	// Create Routes
	engine.POST("/login", authController.PostLogin)
	engine.POST("/register", authController.PostRegister)

	forceNoAuthRequired := engine.Group("/", middleware.ForceNoAuthRequired)
	forceNoAuthRequired.GET("/login", authController.GetLogin)
	forceNoAuthRequired.GET("/register", authController.GetRegister)

	authorized := engine.Group("/", middleware.AuthRequired)
	authorized.GET("/", homeController.GetHome)
	authorized.GET("/logout", authController.GetLogout)

	authorized.GET("/cars", carController.GetAllCars)

	authorized.GET("/cars_post", carController.PostCarPage)
	authorized.POST("/cars_post", carController.PostCar)

	authorized.GET("/cars/:id/edit", carController.GetCarByID)
	authorized.POST("/cars/:id/edit", carController.UpdateCar)
	authorized.POST("/cars/:id/delete", carController.DeleteCar)

	authorized.GET("/profile", carController.CarsPageHandler)

	wishlistGroup := authorized.Group("/wishlist")
	{
		wishlistGroup.POST("/:userID/add/:carID", wishlistController.AddToWishlistHandler)
		wishlistGroup.POST("/:userID/remove/:carID", wishlistController.RemoveFromWishlistHandler)
		wishlistGroup.GET("", wishlistController.GetWishlistHandler)
	}
	return nil
}

func setMongoDBConnection(appConfiguration *configuration.Configuration) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(*appConfiguration.DatabaseConfiguration.ConnectionURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	database := client.Database("your_database_name") // Replace with your actual database name

	return database, nil
}

func setDatabaseConnection(appConfiguration *configuration.Configuration) (*repository.Database, error) {
	// If you're still using Gorm for something else, you can replace this with your Gorm initialization
	database, err := repository.NewDatabase(appConfiguration.DatabaseConfiguration)
	if err != nil {
		return nil, err
	}

	return database, nil
}

func setSession(appConfiguration *configuration.Configuration) gin.HandlerFunc {
	return sessions.Sessions("some.session", sessions.NewCookieStore([]byte(*appConfiguration.SessionConfiguration.Secret)))
}
