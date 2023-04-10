package handler

import (
	"github.com/gin-gonic/gin"
	"go_rocket_launch_sub/internal/pkg/service"
)

type Handler struct {
	servises *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{servises: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		subscription := api.Group("/subscriptions")
		{
			subscription.POST("/", h.createSubscription)
			subscription.DELETE("/:id", h.deleteSubscription)
		}
	}

	return router
}
