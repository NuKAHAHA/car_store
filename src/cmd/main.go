package main

import (
	"log"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/nukahaha/car_store/src/internal/configuration"
	"github.com/nukahaha/car_store/src/internal/controller"
	"github.com/nukahaha/car_store/src/internal/middleware"
	"github.com/nukahaha/car_store/src/internal/repository"
	"github.com/nukahaha/car_store/src/internal/service"
)

// pussy
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
	//
	engine.Static("/public", "./public")

	database, err := setDatabaseConnection(appConfiguration)
	if err != nil {
		log.Fatalf("Error occurred during database connection;\n%s", err.Error())
	}

	err = initAPI(appConfiguration, engine, database)
	if err != nil {
		log.Fatal("Couldn't init routes, middlewares or other systems;\n", err.Error())
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
	// Create Repositories
	userRepository, err := repository.NewUserRepository(database)

	if err != nil {
		return err
	}

	// Create Services
	authService := service.NewAuthService(userRepository)

	// Create Controllers
	homeController := controller.NewHomeController()
	authController := controller.NewAuthController(authService)

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

	return nil
}

func setDatabaseConnection(appConfiguration *configuration.Configuration) (*repository.Database, error) {
	database, err := repository.NewDatabase(appConfiguration.DatabaseConfiguration)

	if err != nil {
		return nil, err
	}

	return database, nil
}

func setSession(appConfiguration *configuration.Configuration) gin.HandlerFunc {
	return sessions.Sessions("some.session", sessions.NewCookieStore([]byte(*appConfiguration.SessionConfiguration.Secret)))
}
