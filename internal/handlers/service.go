package handlers

import (
	"github.com/Dormant512/all-things-cognitei/internal/repository"
	"github.com/gin-gonic/gin"
)

type BusinessLogic interface {
	SrvPing(*gin.Context)
	SrvNewItem(*gin.Context)
	SrvFixItem(*gin.Context)
	SrvGetById(*gin.Context)
	SrvGetInCategory(*gin.Context)
	SrvDeleteById(*gin.Context)
}

type Service struct {
	Rep *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	srv := Service{
		Rep: repo,
	}
	return &srv
}
