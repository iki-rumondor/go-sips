package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/sips/internal/http/request"
	"github.com/iki-rumondor/sips/internal/http/response"
	"github.com/iki-rumondor/sips/internal/models"
)

type AdminHandlerInterface interface {
	SignIn(*gin.Context)
	SetMahasiswaPercepatan(*gin.Context)
	GetMahasiswaPercepatan(*gin.Context)
	SetMahasiswaPeringatan(*gin.Context)
	GetMahasiswaPeringatan(*gin.Context)
	GetUser(*gin.Context)

	CreatePembimbing(*gin.Context)
	GetAllPembimbing(*gin.Context)
	GetPembimbing(*gin.Context)
	UpdatePembimbing(*gin.Context)
	DeletePembimbing(*gin.Context)

	UpdateKelas(*gin.Context)
	GetClasses(*gin.Context)
}

type AdminServiceInterface interface {
	VerifyPengguna(*request.SignIn) (string, error)
	SetMahasiswaPercepatan(*request.PercepatanCond) error
	GetMahasiswaPercepatan() (*[]models.Percepatan, error)
	SetMahasiswaPeringatan() error
	GetMahasiswaPeringatan() (*[]models.Peringatan, error)

	GetUser(userUuid string) (*response.User, error)

	CreatePembimbing(req *request.Pembimbing) error
	FindAllPembimbing() (*[]response.Pembimbing, error)
	FindPembimbing(uuid string) (*response.Pembimbing, error)
	UpdatePembimbing(uuid string, req *request.Pembimbing) error
	DeletePembimbing(uuid string) error

	UpdateKelas(req *request.KelasRule) error
	GetClasses() ([]string, error)
}

type AdminRepoInterface interface {
	FindPenggunaBy(column string, value interface{}) (*models.Pengguna, error)
	FindMahasiswaByRule(ipk float64, total_sks, jumlah_error uint) (*[]models.Mahasiswa, error)
	FindMahasiswaByAngkatan(tahun int) (*[]models.Mahasiswa, error)
	FindMahasiswaPercepatan() (*[]models.Percepatan, error)
	CreateMahasiswaPercepatan(*[]models.Mahasiswa) error
	FindMahasiswaPeringatan() (*[]models.Peringatan, error)
	CreateMahasiswaPeringatan(model *[]models.Mahasiswa, year uint) error

	First(data interface{}, condition string) error
	Find(data interface{}, condition string) error
	FindWithOrder(data interface{}, condition, order string) error
	Truncate(tableName string) error
	Distinct(model interface{}, column string, dest *[]string) error
	Create(data interface{}) error
	Update(data interface{}, condition string) error
	Delete(data interface{}, assoc []string) error
}
