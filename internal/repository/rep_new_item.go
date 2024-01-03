package repository

import (
	"errors"
	"github.com/Dormant512/all-things-cognitei/internal/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) RepNewItem(ctx *gin.Context, item *util.MegaItem) (*primitive.ObjectID, error) {
	// try to find doc with the name
	itemName := item.ItemName
	var findRes util.MegaItem
	r.Mu.RLock()
	err := r.Col.FindOne(ctx, bson.M{"itemName": itemName}).Decode(&findRes)
	r.Mu.RUnlock()
	if err == nil {
		// document found
		return nil, util.DocWithNameExistsError{ItemName: *itemName}
	}
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	// try to add new item to database
	r.Mu.Lock()
	insRes, err := r.Col.InsertOne(ctx, item)
	r.Mu.Unlock()
	if err != nil {
		return nil, err
	}
	var oid = insRes.InsertedID.(primitive.ObjectID)
	return &oid, nil
}
