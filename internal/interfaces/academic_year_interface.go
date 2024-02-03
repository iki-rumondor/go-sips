package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type AcademicYearHandlerInterface interface {
	CreateAcademicYear(*gin.Context)
	GetAllAcademicYears(*gin.Context)
	GetAcademicYear(*gin.Context)
	UpdateAcademicYear(*gin.Context)
	DeleteAcademicYear(*gin.Context)
}

type AcademicYearServiceInterface interface {
	CreateAcademicYear(*request.AcademicYear) error
	GetAllAcademicYears() (*[]models.AcademicYear, error)
	GetAcademicYear(string) (*models.AcademicYear, error)
	UpdateAcademicYear(string, *request.AcademicYear) error
	DeleteAcademicYear(string) error
}

type AcademicYearRepoInterface interface {
	FindAcademicYears() (*[]models.AcademicYear, error)
	FindAcademicYearBy(column string, value interface{}) (*models.AcademicYear, error)
	CreateAcademicYear(*models.AcademicYear) error
	UpdateAcademicYear(*models.AcademicYear) error
	DeleteAcademicYear(*models.AcademicYear) error
}
