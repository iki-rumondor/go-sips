package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/sips/internal/http/request"
	"github.com/iki-rumondor/sips/internal/http/response"
	"github.com/iki-rumondor/sips/internal/models"
)

type MahasiswaHandlerInterface interface {
	Import(*gin.Context)
	GetAll(*gin.Context)
	Get(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type MahasiswaServiceInterface interface {
	ImportMahasiswa(pathFile string) (*[]response.FailedImport, error)
	GetAllMahasiswa() (*[]models.Mahasiswa, error)
	GetMahasiswa(uuid string) (*models.Mahasiswa, error)
	UpdateMahasiswa(uuid string, req *request.Mahasiswa) error
	DeleteMahasiswa(uuid string) error
}

type MahasiswaRepoInterface interface {
	CreateMahasiswa(*models.Mahasiswa) error
	FindAllMahasiswa() (*[]models.Mahasiswa, error)
	FindMahasiswaByUuid(uuid string) (*models.Mahasiswa, error)
	UpdateMahasiswa(model *models.Mahasiswa) error
	DeleteMahasiswa(model *models.Mahasiswa) error
}
