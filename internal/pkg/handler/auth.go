package handler

import (
	"github.com/gin-gonic/gin"
	"go_rocket_launch_sub/internal/pkg/model"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var id, err = h.servises.Authorisation.CreateUser(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"userId": id,
		"token":  "STUB token",
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
