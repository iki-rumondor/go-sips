package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type DepartmentHandlerInterface interface {
	CreateDepartment(*gin.Context)
	GetAllDepartments(*gin.Context)
	GetDepartment(*gin.Context)
	UpdateDepartment(*gin.Context)
	DeleteDepartment(*gin.Context)
}

type DepartmentServiceInterface interface {
	CreateDepartment(*request.Department) error
	GetAllDepartments() (*[]models.Department, error)
	GetDepartment(string) (*models.Department, error)
	UpdateDepartment(string, *request.Department) error
	DeleteDepartment(string) error
}

type DepartmentRepoInterface interface {
	FindDepartments() (*[]models.Department, error)
	FindDepartmentBy(column string, value interface{}) (*models.Department, error)
	FindMajorBy(column string, value interface{}) (*models.Major, error)
	CreateDepartment(*models.Department) error
	UpdateDepartment(*models.Department) error
	DeleteDepartment(*models.Department) error
}
