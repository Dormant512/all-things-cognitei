package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) SrvDeleteById(ctx *gin.Context) {
	// TODO: implement
	ctx.JSON(http.StatusOK,
		gin.H{"message": "PLACEHOLDER"})
}
