package handler

import (
	"github.com/gin-gonic/gin"
	"go_rocket_launch_sub/internal/pkg/model"
	"net/http"
)

type SignInInput struct {
	Email    string `json:"email"    binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var _, err = h.servises.Authorisation.SignUp(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": "STUB token",
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var token, err = h.servises.Authorisation.SignIn(input.Email, input.Password)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) setCurrentUser(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	userId, err := h.servises.Authorisation.ParseToken(authToken)

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "invalid token")
		return
	}

	c.Set("userId", userId)
}
