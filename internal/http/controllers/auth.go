package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/db/sqlc"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/users"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/util"
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
		c.Error(util.BadRequest(err.Error()))
		return
	}

	user, err := ctrl.svc.RegisterUser(requestBody.Username, requestBody.Email, requestBody.Password)
	if err != nil {
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, RegisterResponse{Message: "Successfully registered new user", User: user})
}

func (ctrl AuthController) Login(c *gin.Context) {
	var requestBody LoginRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.Error(util.BadRequest(err.Error()))
		return
	}

	token, err := ctrl.svc.LoginUser(requestBody.Email, requestBody.Password)
	if err != nil {
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, LoginResponse{Message: "Successfully logged in", Token: token})
}
