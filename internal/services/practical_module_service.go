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

type PracticalModuleService struct {
	Repo interfaces.PracticalModuleRepoInterface
}

func NewPracticalModuleService(repo interfaces.PracticalModuleRepoInterface) interfaces.PracticalModuleServiceInterface {
	return &PracticalModuleService{
		Repo: repo,
	}
}

func (s *PracticalModuleService) CreatePracticalModule(userUuid string, req *request.PracticalModule) error {

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

	laboratory, err := s.Repo.FindLaboratoryBy("uuid", req.LaboratoryUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Laboratorium Tidak Ditemukan")
		}
		return response.SERVICE_INTERR
	}

	model := models.PracticalModule{
		Available:      req.Available,
		Note:           req.Note,
		SubjectID:      subject.ID,
		AcademicYearID: academic_year.ID,
		LaboratoryID:   laboratory.ID,
	}

	if err := s.Repo.CreatePracticalModule(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *PracticalModuleService) GetAllPracticalModules(userUuid string) (*[]models.PracticalModule, error) {

	// department, err := s.Repo.FindUserBy("uuid")
	// if err != nil{
	// 	log.Println(err.Error())
	// 	return nil, response.SERVICE_INTERR
	// }

	result, err := s.Repo.FindPracticalModules(userUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *PracticalModuleService) GetPracticalModule(userUuid, uuid string) (*models.PracticalModule, error) {
	result, err := s.Repo.FindUserPracticalModule(userUuid, uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("RPS Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *PracticalModuleService) UpdatePracticalModule(userUuid, uuid string, req *request.UpdatePracticalModule) error {

	result, err := s.GetPracticalModule(userUuid, uuid)
	if err != nil {
		return err
	}

	laboratory, err := s.Repo.FindLaboratoryBy("uuid", req.LaboratoryUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Laboratorium Tidak Ditemukan")
		}
		return response.SERVICE_INTERR
	}

	model := models.PracticalModule{
		ID:           result.ID,
		Available:    req.Available,
		Note:         req.Note,
		LaboratoryID: laboratory.ID,
	}

	if err := s.Repo.UpdatePracticalModule(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *PracticalModuleService) DeletePracticalModule(userUuid, uuid string) error {
	result, err := s.GetPracticalModule(userUuid, uuid)
	if err != nil {
		return err
	}

	model := models.PracticalModule{
		ID: result.ID,
	}

	if err := s.Repo.DeletePracticalModule(&model); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return response.VIOLATED_ERR
		}
		return response.SERVICE_INTERR
	}

	return nil
}
