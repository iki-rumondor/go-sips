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

type LaboratoryService struct {
	Repo interfaces.LaboratoryRepoInterface
}

func NewLaboratoryService(repo interfaces.LaboratoryRepoInterface) interfaces.LaboratoryServiceInterface {
	return &LaboratoryService{
		Repo: repo,
	}
}

func (s *LaboratoryService) CreateLaboratory(userUuid string, req *request.Laboratory) error {

	user, err := s.Repo.FindUserBy("uuid", userUuid)
	if err != nil {
		return err
	}

	model := models.Laboratory{
		DepartmentID: user.Department.ID,
		Name:         req.Name,
	}

	if err := s.Repo.CreateLaboratory(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *LaboratoryService) GetAllLaboratories(userUuid string) (*[]models.Laboratory, error) {

	result, err := s.Repo.FindLaboratories(userUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *LaboratoryService) GetLaboratory(userUuid, uuid string) (*models.Laboratory, error) {
	result, err := s.Repo.FindUserLaboratory(userUuid, uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Laboratorium Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *LaboratoryService) UpdateLaboratory(userUuid, uuid string, req *request.Laboratory) error {

	result, err := s.GetLaboratory(userUuid, uuid)
	if err != nil {
		return err
	}

	model := models.Laboratory{
		ID:           result.ID,
		DepartmentID: result.DepartmentID,
		Name:         req.Name,
	}

	if err := s.Repo.UpdateLaboratory(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *LaboratoryService) DeleteLaboratory(userUuid, uuid string) error {
	result, err := s.GetLaboratory(userUuid, uuid)
	if err != nil {
		return err
	}

	model := models.Laboratory{
		ID: result.ID,
	}

	if err := s.Repo.DeleteLaboratory(&model); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return response.VIOLATED_ERR
		}
		return response.SERVICE_INTERR
	}

	return nil
}
