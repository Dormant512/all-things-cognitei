package handlers

import (
	"errors"
	"fmt"
	"github.com/Dormant512/all-things-cognitei/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) SrvGetById(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"message": "no id provided"})
		return
	}

	item, err := s.Rep.RepGetById(ctx, id)
	if err != nil {
		if errors.Is(err, util.DocWithWithIdNotFoundError{Id: id}) {
			ctx.JSON(http.StatusNotFound,
				gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("unknown error: %v", err)})
		return
	}
	ctx.JSON(http.StatusOK,
		gin.H{
			"message":  "success",
			"document": item,
		})
}
