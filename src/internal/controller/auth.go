package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/nukahaha/car_store/src/internal/configuration"
	"github.com/nukahaha/car_store/src/internal/model/request"
	"github.com/nukahaha/car_store/src/internal/service"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) GetLogin(ctx *gin.Context) {
	var errorBool bool

	hasError, ok := ctx.GetQuery("hasError")

	if !ok {
		errorBool = false
	} else {
		var err error
		errorBool, err = strconv.ParseBool(hasError)

		if err != nil {
			errorBool = false
		}
	}

	errorMessage, ok := ctx.GetQuery("errorMessage")

	if !ok {
		errorMessage = "Hata Oluştu!"
	}

	ginview.HTML(ctx, http.StatusOK, "login", gin.H{
		"title":           "Login",
		"hasError":        errorBool,
		"errorMessage":    errorMessage,
		"isAuthenticated": configuration.IsAuthenticated(ctx),
	})

}

func (ac *AuthController) PostLogin(ctx *gin.Context) {
	loginRequest := &request.LoginRequest{
		Email:    ctx.PostForm("email"),
		Password: ctx.PostForm("password"),
	}

	err := ac.authService.Login(loginRequest)
	if err != nil {
		log.Printf("Logging in failed with: %s", err.Error())
		ctx.Redirect(http.StatusFound, fmt.Sprintf("/login?hasError=true&errorMessage=%s", "Login failed"))

		return
	}

	session := sessions.Default(ctx)
	session.Set(configuration.SessionKeyUser, loginRequest.Email)

	if err := session.Save(); err != nil {
		log.Printf("Logging in failed with: %s", err.Error())
		ctx.Redirect(http.StatusFound, fmt.Sprintf("/login?hasError=true&errorMessage=%s", "Login failed"))

		return
	}

	ctx.Redirect(http.StatusFound, "/")
}

func (ac *AuthController) GetRegister(ctx *gin.Context) {
	var errorBool bool

	hasError, ok := ctx.GetQuery("hasError")

	if !ok {
		errorBool = false
	} else {
		var err error
		errorBool, err = strconv.ParseBool(hasError)

		if err != nil {
			errorBool = false
		}
	}

	errorMessage, ok := ctx.GetQuery("errorMessage")

	if !ok {
		errorMessage = "Error occurred"
	}

	ginview.HTML(ctx, http.StatusOK, "register", gin.H{
		"title":           "Register",
		"hasError":        errorBool,
		"errorMessage":    errorMessage,
		"isAuthenticated": configuration.IsAuthenticated(ctx),
	})

}

func (ac *AuthController) PostRegister(ctx *gin.Context) {
	registerRequest := &request.RegisterRequest{
		Email:           ctx.PostForm("email"),
		Password:        ctx.PostForm("password"),
		ConfirmPassword: ctx.PostForm("confirmPassword"),
		Name:            ctx.PostForm("name"),
		Surname:         ctx.PostForm("surname"),
	}

	ctx.BindWith(registerRequest, binding.FormPost)

	var err error
	registerRequest.Birthday, err = time.Parse("2006-01-02", ctx.PostForm("birthday"))
	if err != nil {
		ctx.Redirect(http.StatusBadRequest, "/register?hasError=true&errorMessage="+err.Error())

		return
	}

	err = ac.authService.Register(registerRequest)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/register?hasError=true&errorMessage="+err.Error())

		return
	}

	ctx.Redirect(http.StatusFound, "/login")
}

func (ac *AuthController) GetLogout(ctx *gin.Context) {
	session := sessions.Default(ctx)

	session.Set(configuration.SessionKeyUser, "")
	session.Clear()
	session.Options(sessions.Options{Path: "/", MaxAge: -1}) // this sets the cookie with a MaxAge of 0
	err := session.Save()

	if err != nil {
		log.Printf("Critical error occured during logout, please check;\n%s", err.Error())
	}

	ctx.Redirect(http.StatusTemporaryRedirect, "/")

}

func GetUserID(ctx *gin.Context) (string, error) {
	session := sessions.Default(ctx)
	userID := session.Get(configuration.SessionKeyUser)
	if userID == nil {
		return "0", errors.New("user ID not found in session")
	}

	// Преобразуем userID в uint
	userIDUint, ok := userID.(string)
	if !ok {
		return "0", errors.New("user ID in session is not of type uint")
	}

	return userIDUint, nil
}
