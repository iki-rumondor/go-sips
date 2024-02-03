package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type TeacherSkillHandlerInterface interface {
	CreateTeacherSkill(*gin.Context)
	GetAllTeacherSkills(*gin.Context)
	GetTeacherSkill(*gin.Context)
	UpdateTeacherSkill(*gin.Context)
	DeleteTeacherSkill(*gin.Context)
}

type TeacherSkillServiceInterface interface {
	CreateTeacherSkill(userUuid string, req *request.TeacherSkill) error
	GetAllTeacherSkills(userUuid string) (*[]models.TeacherSkill, error)
	GetTeacherSkill(userUuid, uuid string) (*models.TeacherSkill, error)
	UpdateTeacherSkill(userUuid, uuid string, req *request.TeacherSkill) error
	DeleteTeacherSkill(userUuid, uuid string) error
}

type TeacherSkillRepoInterface interface {
	FindTeacherSkills(userUuid string) (*[]models.TeacherSkill, error)
	FindUserTeacherSkill(userUuid, uuid string) (*models.TeacherSkill, error)
	FindTeacherBy(column string, value interface{}) (*models.Teacher, error)
	CreateTeacherSkill(*models.TeacherSkill) error
	UpdateTeacherSkill(*models.TeacherSkill) error
	DeleteTeacherSkill(*models.TeacherSkill) error
}
