package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type LaboratoryHandlerInterface interface {
	CreateLaboratory(*gin.Context)
	GetAllLaboratories(*gin.Context)
	GetLaboratory(*gin.Context)
	UpdateLaboratory(*gin.Context)
	DeleteLaboratory(*gin.Context)
}

type LaboratoryServiceInterface interface {
	CreateLaboratory(userUuid string, req *request.Laboratory) error
	GetAllLaboratories(userUuid string) (*[]models.Laboratory, error)
	GetLaboratory(userUuid, uuid string) (*models.Laboratory, error)
	UpdateLaboratory(userUuid, uuid string, req *request.Laboratory) error
	DeleteLaboratory(userUuid, uuid string) error
}

type LaboratoryRepoInterface interface {
	FindLaboratories(userUuid string) (*[]models.Laboratory, error)
	FindUserBy(column string, value interface{}) (*models.User, error)
	FindUserLaboratory(userUuid, uuid string) (*models.Laboratory, error)
	CreateLaboratory(*models.Laboratory) error
	UpdateLaboratory(*models.Laboratory) error
	DeleteLaboratory(*models.Laboratory) error
}
