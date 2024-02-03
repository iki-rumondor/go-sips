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

type PracticalToolService struct {
	Repo interfaces.PracticalToolRepoInterface
}

func NewPracticalToolService(repo interfaces.PracticalToolRepoInterface) interfaces.PracticalToolServiceInterface {
	return &PracticalToolService{
		Repo: repo,
	}
}

func (s *PracticalToolService) CreatePracticalTool(userUuid string, req *request.PracticalTool) error {

	subject, err := s.Repo.FindSubjectBy("uuid", req.SubjectUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Mata Kuliah Tidak Ditemukan")
		}
		return response.SERVICE_INTERR
	}

	if subject.Department.User.Uuid != userUuid {
		return response.NOTFOUND_ERR("RPS Tidak Ditemukan")
	}

	academic_year, err := s.Repo.FindAcademicYearBy("uuid", req.AcademicYearUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Tahun Ajaran Tidak Ditemukan")
		}
		return response.SERVICE_INTERR
	}

	model := models.PracticalTool{
		Condition:      req.Condition,
		Available:      req.Available,
		Note:           req.Note,
		SubjectID:      subject.ID,
		AcademicYearID: academic_year.ID,
	}

	if err := s.Repo.CreatePracticalTool(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *PracticalToolService) GetAllPracticalTools(userUuid string) (*[]models.PracticalTool, error) {

	// department, err := s.Repo.FindUserBy("uuid")
	// if err != nil{
	// 	log.Println(err.Error())
	// 	return nil, response.SERVICE_INTERR
	// }

	result, err := s.Repo.FindPracticalTools(userUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *PracticalToolService) GetPracticalTool(userUuid, uuid string) (*models.PracticalTool, error) {
	result, err := s.Repo.FindUserPracticalTool(userUuid, uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("RPS Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *PracticalToolService) UpdatePracticalTool(userUuid, uuid string, req *request.UpdatePracticalTool) error {

	result, err := s.GetPracticalTool(userUuid, uuid)
	if err != nil {
		return err
	}

	if req.Condition == ""{
		req.Condition = "NONE"
	}

	model := models.PracticalTool{
		ID:             result.ID,
		Condition:      req.Condition,
		Available:      req.Available,
		Note:           req.Note,
	}

	if err := s.Repo.UpdatePracticalTool(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *PracticalToolService) DeletePracticalTool(userUuid, uuid string) error {
	result, err := s.GetPracticalTool(userUuid, uuid)
	if err != nil {
		return err
	}

	model := models.PracticalTool{
		ID: result.ID,
	}

	if err := s.Repo.DeletePracticalTool(&model); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return response.VIOLATED_ERR
		}
		return response.SERVICE_INTERR
	}

	return nil
}
