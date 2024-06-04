package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"payment-system-two/internal/api"
	"payment-system-two/internal/middleware"
	"payment-system-two/internal/ports"
	"time"
)

// SetupRouter is where router endpoints are called
func SetupRouter(handler *api.HTTPHandler, repository ports.Repository) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r := router.Group("/")
	{
		r.GET("/", handler.Readiness)
		r.POST("/create", handler.CreateUser)
		r.POST("/login", handler.LoginUser)
		r.POST("/createAdmin", handler.CreateAdmin)
		r.POST("/adminlogin", handler.LoginAdmin)

	}
	authorizeUser := r.Group("/user")
	{

	}
	authorizeUser.Use(middleware.AuthorizeAdmin(repository.FindUserByEmail, repository.TokenInBlacklist))
	{
		authorizeUser.POST("/transfer", handler.TransferFunds)
		authorizeUser.POST("/addfunds", handler.AddFunds)
		authorizeUser.GET("/balance", handler.BalanceCheck)
		authorizeUser.GET("/transaction", handler.TransactionHistory)
	}

	// authorizeAdmin authorizes all authorized users handlers
	authorizeAdmin := r.Group("/admin")
	authorizeAdmin.Use(middleware.AuthorizeAdmin(repository.FindUserByEmail, repository.TokenInBlacklist))
	{
		authorizeAdmin.GET("/user", handler.GetUserByEmail)

	}

	return router
}
