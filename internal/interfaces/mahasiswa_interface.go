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
	DeleteAll(*gin.Context)

	GetData(*gin.Context)
	GetMahasiswaProdi(*gin.Context)
	GetMahasiswaPercepatan(*gin.Context)
	GetMahasiswaByUserUuid(*gin.Context)
	GetMahasiswaByPenasihat(*gin.Context)
	UpdatePengaturan(*gin.Context)
}

type MahasiswaServiceInterface interface {
	CreateMahasiswaCSV(userUuid, pathFile string) (*[]response.FailedImport, error)
	ImportMahasiswa(pembimbingUuid, pathFile string) (*[]response.FailedImport, error)
	GetAllMahasiswa(options map[string]string) (*[]response.Mahasiswa, error)
	GetMahasiswaProdi(userUuid string) (*[]response.Mahasiswa, error)
	GetMahasiswa(uuid string) (*models.Mahasiswa, error)
	UpdateMahasiswa(uuid string, req *request.Mahasiswa) error
	DeleteMahasiswa(uuid string) error
	DeleteAllMahasiswa(userUuid string) error

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
	FindLimit(data interface{}, condition, order, limit string) error
	First(data interface{}, condition string) error
	FirstOrCreate(dest, model interface{}) error
	Truncate(tableName string) error
	UpdatePengaturan(model *[]models.Pengaturan) error
	UpdateKelas() error
	UpdatePercepatan() error
	DeleteMahasiswaPengguna(data *[]models.Mahasiswa) error
	Create(data interface{}) error
}
