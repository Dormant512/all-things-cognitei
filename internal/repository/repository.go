package repository

import (
	"github.com/Dormant512/all-things-cognitei/internal/config"
	"github.com/Dormant512/all-things-cognitei/internal/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var Repo *Repository

type DataBaser interface {
	RepNewItem(*gin.Context, *util.MegaItem) (interface{}, error)
	RepFixItem() error
	RepGetById() error
	RepGetInCategory() error
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
