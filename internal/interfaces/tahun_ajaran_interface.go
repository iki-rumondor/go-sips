package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/sips/internal/http/request"
	"github.com/iki-rumondor/sips/internal/models"
)

type TahunAjaranHandlerInterface interface {
	Create(*gin.Context)
	GetAll(*gin.Context)
	Get(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type TahunAjaranServiceInterface interface {
	CreateTahunAjaran(*request.TahunAjaran) error
	GetAllTahunAjaran() (*[]models.TahunAjaran, error)
	GetTahunAjaran(string) (*models.TahunAjaran, error)
	UpdateTahunAjaran(string, *request.TahunAjaran) error
	DeleteTahunAjaran(string) error
}

type TahunAjaranRepoInterface interface {
	FindTahunAjaran() (*[]models.TahunAjaran, error)
	FindTahunAjaranBy(column string, value interface{}) (*models.TahunAjaran, error)
	CreateTahunAjaran(*models.TahunAjaran) error
	UpdateTahunAjaran(*models.TahunAjaran) error
	DeleteTahunAjaran(*models.TahunAjaran) error
}
