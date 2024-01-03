package repository

import (
	"github.com/Dormant512/all-things-cognitei/internal/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) RepGetInCategory(ctx *gin.Context, catsToFetch *util.CategoriesToFetch) ([]util.MegaItem, error) {
	var res []util.MegaItem
	filter := bson.M{
		"utilValidTypes": bson.M{
			"$all": catsToFetch.Include,
			"$nin": catsToFetch.Exclude,
		},
	}
	r.Mu.RLock()
	cursor, err := r.Col.Find(ctx, filter)
	r.Mu.RUnlock()
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
