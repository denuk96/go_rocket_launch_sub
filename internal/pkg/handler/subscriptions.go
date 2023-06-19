package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createSubscription(c *gin.Context) {
	userId, _ := c.Get("userId")

	subId, error := h.servises.Subscription.Create(userId.(string))
	if error != nil {
		newErrorResponse(c, http.StatusBadRequest, error.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"subscriptionId": subId,
	})
}

func (h *Handler) deleteSubscription(c *gin.Context) {
	userId, _ := c.Get("userId")
	error := h.servises.Subscription.Destroy(userId.(string), c.Param("id"))

	if error != nil {
		newErrorResponse(c, http.StatusBadRequest, error.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})
}

func (h *Handler) allSubscriptions(c *gin.Context) {
	userId, _ := c.Get("userId")
	subList, error := h.servises.Subscription.AllByUser(userId.(string))

	if error != nil {
		newErrorResponse(c, http.StatusBadRequest, error.Error())
		return
	}

	c.JSON(http.StatusOK, subList)
}
