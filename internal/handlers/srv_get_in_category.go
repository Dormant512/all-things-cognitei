package handlers

import (
	"fmt"
	"github.com/Dormant512/all-things-cognitei/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) SrvGetInCategory(ctx *gin.Context) {
	var catsToFetch util.CategoriesToFetch
	err := ctx.BindJSON(&catsToFetch)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"message": "invalid categories json"})
		return
	}
	err = catsToFetch.MakeValid()
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"message": fmt.Sprintf("invalid categories json: %v", err)})
		return
	}

	items, err := s.Rep.RepGetInCategory(ctx, &catsToFetch)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("unknown error: %v", err)})
		return
	}
	ctx.JSON(http.StatusOK,
		gin.H{
			"message":   "success",
			"documents": items,
		})
}
