package repository

import (
	"errors"
	"github.com/Dormant512/all-things-cognitei/internal/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) RepGetById(ctx *gin.Context, id string) (*util.MegaItem, error) {
	var findRes util.MegaItem
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	r.Mu.RLock()
	err = r.Col.FindOne(ctx, bson.M{"_id": oid}).Decode(&findRes)
	r.Mu.RUnlock()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// document not found
			return nil, util.DocWithWithIdNotFoundError{Id: id}
		}
		return nil, err
	}
	return &findRes, nil
}
