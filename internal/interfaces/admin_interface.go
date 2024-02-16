package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/sips/internal/http/request"
	"github.com/iki-rumondor/sips/internal/models"
)

type AdminHandlerInterface interface {
	SignIn(*gin.Context)
	SetMahasiswaPercepatan(*gin.Context)
	GetMahasiswaPercepatan(*gin.Context)
}

type AdminServiceInterface interface {
	VerifyAdmin(*request.SignIn) (string, error)
	SetMahasiswaPercepatan(*request.PercepatanCond) error
	GetMahasiswaPercepatan() (*[]models.Percepatan, error)
}

type AdminRepoInterface interface {
	FindAdminBy(column string, value interface{}) (*models.Admin, error)
	FindMahasiswaByRule(ipk float64, total_sks, jumlah_error, angkatan uint) (*[]models.Mahasiswa, error)
	FindMahasiswaPercepatan() (*[]models.Percepatan, error)
	CreateMahasiswaPercepatan(*[]models.Mahasiswa) error
}
