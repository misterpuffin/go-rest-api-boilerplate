package controllers

import (
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/users"
)

type AuthController struct {
	userService users.Service
}

func CreateUser(c *gin.Context) {

}
