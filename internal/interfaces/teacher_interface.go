package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type TeacherHandlerInterface interface {
	CreateTeacher(*gin.Context)
	GetAllTeachers(*gin.Context)
	GetTeacher(*gin.Context)
	UpdateTeacher(*gin.Context)
	DeleteTeacher(*gin.Context)
}

type TeacherServiceInterface interface {
	CreateTeacher(userUuid string, req *request.Teacher) error
	GetAllTeachers(userUuid string) (*[]models.Teacher, error)
	GetTeacher(userUuid, uuid string) (*models.Teacher, error)
	UpdateTeacher(userUuid, uuid string, req *request.Teacher) error
	DeleteTeacher(userUuid, uuid string) error
}

type TeacherRepoInterface interface {
	FindTeachers(userUuid string) (*[]models.Teacher, error)
	FindUserBy(column string, value interface{}) (*models.User, error)
	FindUserTeacher(userUuid, uuid string) (*models.Teacher, error)
	CreateTeacher(*models.Teacher) error
	UpdateTeacher(*models.Teacher) error
	DeleteTeacher(*models.Teacher) error
}
