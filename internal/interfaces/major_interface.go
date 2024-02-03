package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type MajorHandlerInterface interface {
	CreateMajor(*gin.Context)
	GetAllMajors(*gin.Context)
	GetMajor(*gin.Context)
	UpdateMajor(*gin.Context)
	DeleteMajor(*gin.Context)
}

type MajorServiceInterface interface {
	CreateMajor(*request.Major) error
	GetAllMajors() (*[]models.Major, error)
	GetMajor(string) (*models.Major, error)
	UpdateMajor(string, *request.Major) error
	DeleteMajor(string) error
}

type MajorRepoInterface interface {
	FindMajors() (*[]models.Major, error)
	FindMajorBy(column string, value interface{}) (*models.Major, error)
	CreateMajor(*models.Major) error
	UpdateMajor(*models.Major) error
	DeleteMajor(*models.Major) error
}
