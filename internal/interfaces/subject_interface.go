package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type SubjectHandlerInterface interface {
	CreateSubject(*gin.Context)
	GetAllSubjects(*gin.Context)
	GetAllPracticalSubjects(*gin.Context)
	GetSubject(*gin.Context)
	UpdateSubject(*gin.Context)
	DeleteSubject(*gin.Context)
}

type SubjectServiceInterface interface {
	CreateSubject(userUuid string, req *request.Subject) error
	GetAllSubjects(userUuid string) (*[]models.Subject, error)
	GetPracticalSubjects(userUuid string) (*[]models.Subject, error)
	GetSubject(userUuid, uuid string) (*models.Subject, error)
	UpdateSubject(userUuid, uuid string, req *request.Subject) error
	DeleteSubject(userUuid, uuid string) error
}

type SubjectRepoInterface interface {
	FindSubjects(userUuid string) (*[]models.Subject, error)
	FindPracticalSubjects(userUuid string) (*[]models.Subject, error)
	FindUserBy(column string, value interface{}) (*models.User, error)
	FindUserSubject(userUuid, uuid string) (*models.Subject, error)
	CreateSubject(*models.Subject) error
	UpdateSubject(*models.Subject) error
	DeleteSubject(*models.Subject) error
}
