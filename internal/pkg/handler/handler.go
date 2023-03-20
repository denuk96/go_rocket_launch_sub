package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signUp)
	}

	api := router.Group("/api")
	{
		subscription := api.Group("/subscriptions")
		{
			subscription.POST("/:id", h.createSubscription)
			subscription.DELETE("/:id", h.deleteSubscription)
		}
	}

	return router
}
