package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	Message string
	Status  int
	Data    any
}
type HttpError struct {
	Message string
	Status  int
	Data    any
}

func (e HttpError) Error() string {
	return fmt.Sprintf("Description: %s, Data: %s", e.Message, e.Data)
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		switch e := err.Err.(type) {
		case HttpError:
			c.AbortWithStatusJSON(e.Status, e)
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
			return
		}
	}
}
