package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type AcademicPlanHandlerInterface interface {
	CreateAcademicPlan(*gin.Context)
	GetAllAcademicPlans(*gin.Context)
	GetAcademicPlan(*gin.Context)
	UpdateAcademicPlan(*gin.Context)
	DeleteAcademicPlan(*gin.Context)
}

type AcademicPlanServiceInterface interface {
	CreateAcademicPlan(userUuid string, req *request.AcademicPlan) error
	GetAllAcademicPlans(userUuid string) (*[]models.AcademicPlan, error)
	GetAcademicPlan(userUuid, uuid string) (*models.AcademicPlan, error)
	UpdateAcademicPlan(userUuid, uuid string, req *request.UpdateAcademicPlan) error
	DeleteAcademicPlan(userUuid, uuid string) error
}

type AcademicPlanRepoInterface interface {
	FindAcademicPlans(userUuid string) (*[]models.AcademicPlan, error)
	FindUserAcademicPlan(userUuid, uuid string) (*models.AcademicPlan, error)
	FindSubjectBy(column string, value interface{}) (*models.Subject, error)
	FindAcademicYearBy(column string, value interface{}) (*models.AcademicYear, error)
	CreateAcademicPlan(*models.AcademicPlan) error
	UpdateAcademicPlan(*models.AcademicPlan) error
	DeleteAcademicPlan(*models.AcademicPlan) error
}
