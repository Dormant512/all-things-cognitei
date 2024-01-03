package repository

import (
	"errors"
	"github.com/Dormant512/all-things-cognitei/internal/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) RepNewItem(ctx *gin.Context, item *util.MegaItem) (interface{}, error) {
	// try to find doc with the name
	itemName := item.ItemName
	var findRes bson.M
	Repo.Mu.RLock()
	err := r.Col.FindOne(ctx, bson.M{"itemName": itemName}).Decode(&findRes)
	Repo.Mu.RUnlock()
	if err == nil {
		// document found
		return nil, util.DocWithNameExistsError{ItemName: *itemName}
	}
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	// try to add new item to database
	// TODO: remove placeholder
	insRes, err := r.Col.InsertOne(ctx, bson.M{"itemName": itemName, "placeholder": true})
	if err != nil {
		return nil, err
	}
	return insRes.InsertedID, nil
}
