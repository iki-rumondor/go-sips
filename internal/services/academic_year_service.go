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

type AcademicYearService struct {
	Repo interfaces.AcademicYearRepoInterface
}

func NewAcademicYearService(repo interfaces.AcademicYearRepoInterface) interfaces.AcademicYearServiceInterface {
	return &AcademicYearService{
		Repo: repo,
	}
}

func (s *AcademicYearService) CreateAcademicYear(req *request.AcademicYear) error {

	model := models.AcademicYear{
		Name: req.Name,
	}

	if err := s.Repo.CreateAcademicYear(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *AcademicYearService) GetAllAcademicYears() (*[]models.AcademicYear, error) {
	result, err := s.Repo.FindAcademicYears()
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *AcademicYearService) GetAcademicYear(uuid string) (*models.AcademicYear, error) {
	result, err := s.Repo.FindAcademicYearBy("uuid", uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Lab Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *AcademicYearService) UpdateAcademicYear(uuid string, req *request.AcademicYear) error {

	result, err := s.GetAcademicYear(uuid)
	if err != nil {
		return err
	}

	model := models.AcademicYear{
		ID:   result.ID,
		Name: req.Name,
	}

	if err := s.Repo.UpdateAcademicYear(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *AcademicYearService) DeleteAcademicYear(uuid string) error {
	result, err := s.GetAcademicYear(uuid)
	if err != nil {
		return err
	}

	model := models.AcademicYear{
		ID: result.ID,
	}

	if err := s.Repo.DeleteAcademicYear(&model); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return response.VIOLATED_ERR
		}
		return response.SERVICE_INTERR
	}

	return nil
}
