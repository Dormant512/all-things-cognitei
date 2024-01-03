package repository

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (r *Repository) RepPing(ctx *gin.Context) error {
	return r.DB.Ping(ctx, readpref.Primary())
}
