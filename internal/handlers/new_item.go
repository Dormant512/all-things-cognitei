package handlers

import (
	"errors"
	"fmt"
	"github.com/Dormant512/all-things-cognitei/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewItem(ctx *gin.Context) {
	var megaItem util.MegaItem
	err := ctx.BindJSON(&megaItem)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"message": "invalid input json"})
		return
	}

	// TODO: validate itemType of MegaItem, send only useful info to repo
	err = megaItem.FillUtil()
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"message": err.Error()})
		return
	}

	insId, err := Repo.RepNewItem(ctx, &megaItem)
	if err != nil {
		if errors.Is(err, util.DocWithNameExistsError{ItemName: *megaItem.ItemName}) {
			ctx.JSON(http.StatusBadRequest,
				gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("unknown error: %v", err)})
		return
	}
	ctx.JSON(http.StatusCreated,
		gin.H{
			"message":    "success",
			"insertedId": insId,
		})
}
