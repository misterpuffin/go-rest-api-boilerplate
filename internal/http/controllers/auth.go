package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/users"
)

type AuthController struct {
	userService *users.Service
}

type RegisterRequestBody struct {
	Email    string
	Password string
}
type RegisterResponse struct {
	Message string
}

func NewAuthController(userService *users.Service) *AuthController {
	return &AuthController{userService}
}

func (ctrl AuthController) Register(c *gin.Context) {
	var requestBody RegisterRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, RegisterResponse{Message: "ok!"})
}
