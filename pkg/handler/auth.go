package handler

import (
	"net/http"

	"github.com/404th/todolist/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]int{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var sign_in_input model.SignInInput
	if err := c.BindJSON(&sign_in_input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	jwt, err := h.service.GetUser(sign_in_input)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"token": jwt,
	})
}
