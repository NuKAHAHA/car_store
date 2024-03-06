package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/contrib/sessions"

	"github.com/nukahaha/car_store/src/internal/configuration"
)

func AuthRequired(ctx *gin.Context) {
	if !configuration.IsAuthenticated(ctx) {
		ctx.Redirect(http.StatusFound, "/login")

		return
	}


	ctx.Next()
}

func ForceNoAuthRequired(ctx *gin.Context) {
	if configuration.IsAuthenticated(ctx) {
		ctx.Redirect(http.StatusFound, "/")
		return
	}

	ctx.Next()
}
