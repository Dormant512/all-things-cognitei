package repository

import (
	"github.com/Dormant512/all-things-cognitei/internal/config"
	"github.com/Dormant512/all-things-cognitei/internal/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type DataBaser interface {
	RepPing(ctx *gin.Context) error
	RepNewItem(*gin.Context, *util.MegaItem) (*primitive.ObjectID, error)
	RepGetById(*gin.Context, string) (*util.MegaItem, error)
	RepGetInCategory() error
	RepDeleteById(*gin.Context, string) error
}

type Repository struct {
	DB  *mongo.Client
	Col *mongo.Collection
	Cfg *config.Config
	Mu  *sync.RWMutex
}

func NewRepository(database *mongo.Client, con *config.Config) *Repository {
	adminDB := database.Database("admin")
	collection := adminDB.Collection(con.MGCollection)
	mu := &sync.RWMutex{}
	repo := Repository{
		DB:  database,
		Col: collection,
		Cfg: con,
		Mu:  mu,
	}
	return &repo
}
