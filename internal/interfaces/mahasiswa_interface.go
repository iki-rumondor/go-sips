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
	GetMahasiswaPercepatan(*gin.Context)
	GetMahasiswaByUserUuid(*gin.Context)
	GetMahasiswaByPenasihat(*gin.Context)
	UpdatePengaturan(*gin.Context)
}

type MahasiswaServiceInterface interface {
	ImportMahasiswa(pembimbingUuid, pathFile string) (*[]response.FailedImport, error)
	GetAllMahasiswa(options map[string]string) (*[]response.Mahasiswa, error)
	GetMahasiswa(uuid string) (*models.Mahasiswa, error)
	UpdateMahasiswa(uuid string, req *request.Mahasiswa) error
	DeleteMahasiswa(uuid string) error

	GetMahasiswaPercepatan() (*[]response.Mahasiswa, error)
	GetDataMahasiswa(nim string) (*response.Mahasiswa, error)
	GetMahasiswaByUserUuid(userUuid string) (*response.Mahasiswa, error)
	GetAllMahasiswaByPenasihat(userUuid string, options map[string]string) (*[]response.Mahasiswa, error)
	UpdatePengaturan(req *request.Pengaturan) error
}

type MahasiswaRepoInterface interface {
	CreateMahasiswa(*models.Mahasiswa) error
	FindAllMahasiswa(condition string) (*[]models.Mahasiswa, error)
	FindMahasiswaByUuid(uuid string) (*models.Mahasiswa, error)
	UpdateMahasiswa(model *models.Mahasiswa) error
	DeleteMahasiswa(model *models.Mahasiswa) error

	FindBy(tableName, column string, value interface{}) (map[string]interface{}, error)
	Find(data interface{}, condition, order string) error
	First(data interface{}, condition string) error
	UpdatePengaturan(model *[]models.Pengaturan) error
	UpdateKelas() error
	UpdatePercepatan() error
}
