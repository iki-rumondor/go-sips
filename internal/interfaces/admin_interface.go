package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/sips/internal/http/request"
	"github.com/iki-rumondor/sips/internal/http/response"
	"github.com/iki-rumondor/sips/internal/models"
)

type AdminHandlerInterface interface {
	SignIn(*gin.Context)
	GetUser(*gin.Context)

	CreatePembimbing(*gin.Context)
	GetAllPembimbing(*gin.Context)
	GetPembimbing(*gin.Context)
	UpdatePembimbing(*gin.Context)
	DeletePembimbing(*gin.Context)

	UpdateKelas(*gin.Context)
	GetClasses(*gin.Context)
	GetPenasihatDashboard(*gin.Context)
	GetKaprodiDashboard(*gin.Context)
	GetPengaturan(*gin.Context)
}

type AdminServiceInterface interface {
	VerifyPengguna(*request.SignIn) (string, error)
	GetUser(userUuid string) (*response.User, error)

	CreatePembimbing(req *request.Pembimbing) error
	FindAllPembimbing() (*[]response.Pembimbing, error)
	FindPembimbing(uuid string) (*response.Pembimbing, error)
	UpdatePembimbing(uuid string, req *request.Pembimbing) error
	DeletePembimbing(uuid string) error

	UpdateKelas(req *request.KelasRule) error
	GetClasses() ([]string, error)
	GetPenasihatDashboard(userUuid string) (map[string]interface{}, error)
	GetKaprodiDashboard() (map[string]interface{}, error)
	GetPengaturan() (*[]response.Pengaturan, error)
}

type AdminRepoInterface interface {
	FindPenggunaBy(column string, value interface{}) (*models.Pengguna, error)
	FindMahasiswaByAngkatan(tahun int) (*[]models.Mahasiswa, error)
	FindPenasihatPercepatan(dest *[]models.Percepatan, penasihatID uint) error

	First(data interface{}, condition string) error
	Find(data interface{}, condition string) error
	FindWithOrder(data interface{}, condition, order string) error
	Truncate(tableName string) error
	Distinct(model interface{}, column, condition string, dest *[]string) error
	Create(data interface{}) error
	Update(data interface{}, condition string) error
	Delete(data interface{}, assoc []string) error
}
