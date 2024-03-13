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

	GetData(*gin.Context)
	GetMahasiswaByUserUuid(*gin.Context)
	GetMahasiswaByPenasihat(*gin.Context)
}

type MahasiswaServiceInterface interface {
	ImportMahasiswa(pembimbingUuid, pathFile string) (*[]response.FailedImport, error)
	GetAllMahasiswa(options map[string]string) (*[]models.Mahasiswa, error)
	GetMahasiswa(uuid string) (*models.Mahasiswa, error)
	UpdateMahasiswa(uuid string, req *request.Mahasiswa) error
	DeleteMahasiswa(uuid string) error

	GetDataMahasiswa(nim string) (*response.DataMahasiswa, error)
	GetMahasiswaByUserUuid(userUuid string) (*response.Mahasiswa, error)
	GetAllMahasiswaByPenasihat(userUuid string, options map[string]string) (*[]response.Mahasiswa, error)
	// GetPercepatanByPenasihat(userUuid string) (*[]response.Mahasiswa, error)
}

type MahasiswaRepoInterface interface {
	CreateMahasiswa(*models.Mahasiswa) error
	FindAllMahasiswa(condition string) (*[]models.Mahasiswa, error)
	FindMahasiswaByUuid(uuid string) (*models.Mahasiswa, error)
	UpdateMahasiswa(model *models.Mahasiswa) error
	DeleteMahasiswa(model *models.Mahasiswa) error

	// FindPercepatanPenasihat(dest *[]models.Percepatan) error
	FindBy(tableName, column string, value interface{}) (map[string]interface{}, error)
	Find(data interface{}, condition, order string) error
	First(data interface{}, condition string) error
}
