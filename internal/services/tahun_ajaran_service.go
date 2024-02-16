package services

import (
	"errors"
	"log"
	"strconv"

	"github.com/iki-rumondor/sips/internal/http/request"
	"github.com/iki-rumondor/sips/internal/http/response"
	"github.com/iki-rumondor/sips/internal/interfaces"
	"github.com/iki-rumondor/sips/internal/models"
	"gorm.io/gorm"
)

type TahunAjaranService struct {
	Repo interfaces.TahunAjaranRepoInterface
}

func NewTahunAjaranService(repo interfaces.TahunAjaranRepoInterface) interfaces.TahunAjaranServiceInterface {
	return &TahunAjaranService{
		Repo: repo,
	}
}

func (s *TahunAjaranService) CreateTahunAjaran(req *request.TahunAjaran) error {

	tahun, _ := strconv.Atoi(req.Tahun)
	model := models.TahunAjaran{
		Tahun:    uint(tahun),
		Semester: req.Semester,
	}

	if err := s.Repo.CreateTahunAjaran(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *TahunAjaranService) GetAllTahunAjaran() (*[]models.TahunAjaran, error) {
	result, err := s.Repo.FindTahunAjaran()
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *TahunAjaranService) GetTahunAjaran(uuid string) (*models.TahunAjaran, error) {
	result, err := s.Repo.FindTahunAjaranBy("uuid", uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Tahun Ajaran Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *TahunAjaranService) UpdateTahunAjaran(uuid string, req *request.TahunAjaran) error {

	result, err := s.GetTahunAjaran(uuid)
	if err != nil {
		return err
	}

	tahun, _ := strconv.Atoi(req.Tahun)
	model := models.TahunAjaran{
		ID:       result.ID,
		Tahun:    uint(tahun),
		Semester: req.Semester,
	}

	if err := s.Repo.UpdateTahunAjaran(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *TahunAjaranService) DeleteTahunAjaran(uuid string) error {
	result, err := s.GetTahunAjaran(uuid)
	if err != nil {
		return err
	}

	model := models.TahunAjaran{
		ID: result.ID,
	}

	if err := s.Repo.DeleteTahunAjaran(&model); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return response.VIOLATED_ERR
		}
		return response.SERVICE_INTERR
	}

	return nil
}
