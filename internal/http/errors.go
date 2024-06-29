package http

import (
	"net/http"

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
