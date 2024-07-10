package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/util"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	err := c.Errors.Last()
	if err != nil {
		switch e := err.Err.(type) {
		case util.HttpError:
			c.AbortWithStatusJSON(e.Status, e)
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
			return
		}
	}
}

func AuthHandler(config util.Config) func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		jwtTokens := strings.Split(authHeader, " ")
		if len(jwtTokens) < 2 {
			c.Error(errors.New("Unauthorized"))
			return
		}
		jwtTokenString := jwtTokens[1]

		token, err := util.ParseJWTToken(jwtTokenString, config)
		if err != nil {
			c.Error(errors.New("Unauthorized"))
			return
		}

		c.Set("User", token)
		c.Next()
	}
}
