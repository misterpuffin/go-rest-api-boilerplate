package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/config"
	custom_errors "github.com/misterpuffin/go-rest-api-boilerplate/internal/errors"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/util"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	err := c.Errors.Last()
	if err != nil {
		switch e := err.Err.(type) {
		case custom_errors.HttpError:
			c.AbortWithStatusJSON(e.Status, map[string]string{"message": e.Message})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
			return
		}
	}
}

func AuthHandler(config config.Config) func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		jwtTokens := strings.Split(authHeader, " ")
		if len(jwtTokens) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
			return
		}
		jwtTokenString := jwtTokens[1]

		token, err := util.ParseJWTToken(jwtTokenString, config)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
			return
		}

		parsedId, err := uuid.Parse(token.UserId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
			return
		}

		c.Set("UserId", parsedId)
		c.Next()
	}
}
