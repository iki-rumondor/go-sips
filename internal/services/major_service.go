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

type MajorService struct {
	Repo interfaces.MajorRepoInterface
}

func NewMajorService(repo interfaces.MajorRepoInterface) interfaces.MajorServiceInterface {
	return &MajorService{
		Repo: repo,
	}
}

func (s *MajorService) CreateMajor(req *request.Major) error {

	model := models.Major{
		Name: req.Name,
	}

	if err := s.Repo.CreateMajor(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *MajorService) GetAllMajors() (*[]models.Major, error) {
	result, err := s.Repo.FindMajors()
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *MajorService) GetMajor(uuid string) (*models.Major, error) {
	result, err := s.Repo.FindMajorBy("uuid", uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Jurusan Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *MajorService) UpdateMajor(uuid string, req *request.Major) error {
	major, err := s.GetMajor(uuid)
	if err != nil {
		return err
	}

	model := models.Major{
		ID:   major.ID,
		Name: req.Name,
	}

	if err := s.Repo.UpdateMajor(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *MajorService) DeleteMajor(uuid string) error {
	major, err := s.GetMajor(uuid)
	if err != nil {
		return err
	}

	model := models.Major{
		ID: major.ID,
	}

	if err := s.Repo.DeleteMajor(&model); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return response.VIOLATED_ERR
		}
		return response.SERVICE_INTERR
	}

	return nil
}
