package services

import (
	"errors"
	"log"

	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type TeacherService struct {
	Repo interfaces.TeacherRepoInterface
}

func NewTeacherService(repo interfaces.TeacherRepoInterface) interfaces.TeacherServiceInterface {
	return &TeacherService{
		Repo: repo,
	}
}

func (s *TeacherService) CreateTeacher(userUuid string, req *request.Teacher) error {

	user, err := s.Repo.FindUserBy("uuid", userUuid)
	if err != nil {
		return response.SERVICE_INTERR
	}

	model := models.Teacher{
		DepartmentID: user.Department.ID,
		Name:         req.Name,
	}

	if err := s.Repo.CreateTeacher(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *TeacherService) GetAllTeachers(userUuid string) (*[]models.Teacher, error) {

	result, err := s.Repo.FindTeachers(userUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *TeacherService) GetTeacher(userUuid, uuid string) (*models.Teacher, error) {
	result, err := s.Repo.FindUserTeacher(userUuid, uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Dosen Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *TeacherService) UpdateTeacher(userUuid, uuid string, req *request.Teacher) error {

	result, err := s.GetTeacher(userUuid, uuid)
	if err != nil {
		return err
	}

	model := models.Teacher{
		ID:           result.ID,
		DepartmentID: result.DepartmentID,
		Name:         req.Name,
	}

	if err := s.Repo.UpdateTeacher(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *TeacherService) DeleteTeacher(userUuid, uuid string) error {
	result, err := s.GetTeacher(userUuid, uuid)
	if err != nil {
		return err
	}

	model := models.Teacher{
		ID: result.ID,
	}

	if err := s.Repo.DeleteTeacher(&model); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return response.VIOLATED_ERR
		}
		return response.SERVICE_INTERR
	}

	return nil
}
