package handler

import (
	"go_rocket_launch_sub/internal/pkg/handler/middlewares"
	"go_rocket_launch_sub/internal/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	servises *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{servises: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(middlewares.LoggingMiddleware())
	router.Use(gin.Recovery()) // return 500 in case of error

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		subscription := api.Group("/subscriptions", h.setCurrentUser)
		{
			subscription.POST("/", h.createSubscription)
			subscription.DELETE("/:id", h.deleteSubscription)
		}
	}

	return router
}
