package util

import (
	"fmt"

	"net/http"
)

type HttpError struct {
	Message string
	Status  int
}

func (e HttpError) Error() string {
	return fmt.Sprintf("Description: %s", e.Message)
}

func BadRequest(message string) HttpError {
	return HttpError{Message: message, Status: http.StatusBadRequest}
}