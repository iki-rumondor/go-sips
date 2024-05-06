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
	GetPembimbingProdi(*gin.Context)
	GetPembimbing(*gin.Context)
	UpdatePembimbing(*gin.Context)
	DeletePembimbing(*gin.Context)

	CreateProdi(*gin.Context)
	GetAllProdi(*gin.Context)
	GetProdi(*gin.Context)
	UpdateProdi(*gin.Context)
	DeleteProdi(*gin.Context)

	UpdateKelas(*gin.Context)
	GetClasses(*gin.Context)
	GetPenasihatDashboard(*gin.Context)
	GetKaprodiDashboard(*gin.Context)
	GetKajurDashboard(*gin.Context)
	GetPengaturan(*gin.Context)
	GetPengaturanByName(*gin.Context)

	GetAllUsers(*gin.Context)
	CreateKajur(*gin.Context)
	UpdateUsername(*gin.Context)
	UpdatePassword(*gin.Context)
	RekomendasiMahasiswa(*gin.Context)
}

type AdminServiceInterface interface {
	VerifyPengguna(*request.SignIn) (string, error)
	GetUser(userUuid string) (*response.User, error)

	CreatePembimbing(req *request.Pembimbing) error
	FindAllPembimbing() (*[]response.Pembimbing, error)
	FindPembimbingProdi(userUuid string) (*[]response.Pembimbing, error)
	FindPembimbing(uuid string) (*response.Pembimbing, error)
	UpdatePembimbing(uuid string, req *request.UpdatePembimbing) error
	DeletePembimbing(uuid string) error

	CreateProdi(req *request.Prodi) error
	FindAllProdi() (*[]response.Prodi, error)
	FindProdi(uuid string) (*response.Prodi, error)
	UpdateProdi(uuid string, req *request.Prodi) error
	DeleteProdi(uuid string) error

	UpdateKelas(req *request.KelasRule) error
	GetClasses() ([]string, error)
	GetPenasihatDashboard(userUuid string) (map[string]interface{}, error)
	GetKaprodiDashboard(userUuid string) (map[string]interface{}, error)
	GetKajurDashboard() (map[string]interface{}, error)
	GetPengaturan() (*[]response.Pengaturan, error)
	GetPengaturanByName(name string) (*response.Pengaturan, error)

	FindAllUsers() (*[]response.User, error)
	CreateKajur(req *request.Kajur) error
	UpdateUsername(uuid string, req *request.UpdateUsername) error
	UpdatePassword(uuid string, req *request.UpdatePassword) error
	
	RekomendasiMahasiswa(req *request.RekomendasiMahasiswa) error
}

type AdminRepoInterface interface {
	FindPenggunaBy(column string, value interface{}) (*models.Pengguna, error)
	FindMahasiswaByAngkatan(tahun int) (*[]models.Mahasiswa, error)
	DistinctProdiMahasiswa(prodiID uint, dest *[]string, column string) error
	FindProdiMahasiswa(prodiID uint, dest *[]models.Mahasiswa, condition string) error

	First(data interface{}, condition string) error
	Find(data interface{}, condition string) error
	FindWithOrder(data interface{}, condition, order string) error
	Truncate(tableName string) error
	Distinct(model interface{}, column, condition string, dest *[]string) error
	Create(data interface{}) error
	Update(data interface{}, condition string) error
	Delete(data interface{}, assoc []string) error
}
