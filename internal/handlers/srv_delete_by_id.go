package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) SrvDeleteById(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"message": "no id provided"})
		return
	}

	err := s.Rep.RepDeleteById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("unknown error: %v", err)})
		return
	}
	ctx.JSON(http.StatusOK,
		gin.H{
			"message": "success",
			"deleted": id,
		})
}
