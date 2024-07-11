package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/db/sqlc"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/errors"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/users"
)

type AuthController struct {
	svc *users.Service
}

type RegisterRequestBody struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type RegisterResponse struct {
	Message string
	User    db.User
}
type LoginRequestBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginResponse struct {
	Message string
	Token   string
}

func NewAuthController(userService *users.Service) *AuthController {
	return &AuthController{userService}
}

func (ctrl AuthController) Register(c *gin.Context) {
	var requestBody RegisterRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		_ = c.Error(errors.BadRequest(err.Error())) // ignore error since we handle it
		return
	}

	user, err := ctrl.svc.RegisterUser(requestBody.Username, requestBody.Email, requestBody.Password)
	if err != nil {
		_ = c.Error(err) // ignore error since we handle it
		return
	}

	c.IndentedJSON(http.StatusOK, RegisterResponse{Message: "Successfully registered new user", User: user})
}

func (ctrl AuthController) Login(c *gin.Context) {
	var requestBody LoginRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		_ = c.Error(errors.BadRequest(err.Error())) // ignore error since we handle it
		return
	}

	token, err := ctrl.svc.LoginUser(requestBody.Email, requestBody.Password)
	if err != nil {
		_ = c.Error(err) // ignore error since we handle it
		return
	}

	c.IndentedJSON(http.StatusOK, LoginResponse{Message: "Successfully logged in", Token: token})
}
