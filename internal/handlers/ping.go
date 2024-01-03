package handlers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
)

func Ping(ctx *gin.Context) {
	err := Repo.DB.Ping(ctx, readpref.Primary())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"message": "mongoDB broke"})
		return
	}
	ctx.JSON(http.StatusOK,
		gin.H{"message": "pong"})
}
