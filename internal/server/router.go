package server

import (
	"bankapp/internal/handlers"
	"bankapp/internal/middleware"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.ErrorHandler())
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	auth := r.Group("/api")
	auth.Use(middleware.JWTAuth())
	{
		auth.GET("/profile", handlers.Profile)
		auth.POST("/accounts/:id/deposit", handlers.Deposit)
		auth.POST("/accounts/:id/withdraw", handlers.Withdraw)
		auth.GET("/accounts/:id", handlers.GetBalance)
	}

	return r
}
