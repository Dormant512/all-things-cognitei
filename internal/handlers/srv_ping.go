package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) SrvPing(ctx *gin.Context) {
	err := s.Rep.RepPing(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"message": "mongoDB ping failed"})
		return
	}
	ctx.JSON(http.StatusOK,
		gin.H{"message": "pong"})
}
