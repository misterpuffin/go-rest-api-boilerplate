package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/db/sqlc"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/errors"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/patterns"
)

type PatternController struct {
	svc *patterns.Service
}

type PostPatternBody struct {
	Instructions string `json:"instructions" binding:"required"`
}
type PostPatternResponse struct {
	Message string
	Pattern db.Pattern
}
type GetPatternResponse struct {
	Message string
	Pattern db.Pattern
}

func NewPatternController(patternService *patterns.Service) *PatternController {
	return &PatternController{patternService}
}

func (ctrl PatternController) Post(c *gin.Context) {
	var requestBody PostPatternBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.Error(errors.BadRequest(err.Error()))
		return
	}

	userId, ok := c.Get("UserId")
	if !ok {
		errors.Unauthorized()
	}
	pattern, err := ctrl.svc.CreatePattern(userId.(uuid.UUID), requestBody.Instructions)
	if err != nil {
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, PostPatternResponse{Message: "Successfully created new pattern", Pattern: pattern})
}

func (ctrl PatternController) Get(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.Error(errors.BadRequest("Please enter a pattern id"))
	}

	pattern, err := ctrl.svc.GetPattern(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, GetPatternResponse{Message: "Pattern found", Pattern: pattern})
}
