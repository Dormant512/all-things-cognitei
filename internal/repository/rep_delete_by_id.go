package repository

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) RepDeleteById(ctx *gin.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	r.Mu.Lock()
	_, err = r.Col.DeleteOne(ctx, bson.M{"_id": oid})
	r.Mu.Unlock()
	if err != nil {
		return err
	}
	return nil
}
