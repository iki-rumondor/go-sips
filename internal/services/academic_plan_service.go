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

type AcademicPlanService struct {
	Repo interfaces.AcademicPlanRepoInterface
}

func NewAcademicPlanService(repo interfaces.AcademicPlanRepoInterface) interfaces.AcademicPlanServiceInterface {
	return &AcademicPlanService{
		Repo: repo,
	}
}

func (s *AcademicPlanService) CreateAcademicPlan(userUuid string, req *request.AcademicPlan) error {

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

	model := models.AcademicPlan{
		Available:      req.Available,
		Note:           req.Note,
		SubjectID:      subject.ID,
		AcademicYearID: academic_year.ID,
	}

	if err := s.Repo.CreateAcademicPlan(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *AcademicPlanService) GetAllAcademicPlans(userUuid string) (*[]models.AcademicPlan, error) {

	// department, err := s.Repo.FindUserBy("uuid")
	// if err != nil{
	// 	log.Println(err.Error())
	// 	return nil, response.SERVICE_INTERR
	// }

	result, err := s.Repo.FindAcademicPlans(userUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *AcademicPlanService) GetAcademicPlan(userUuid, uuid string) (*models.AcademicPlan, error) {
	result, err := s.Repo.FindUserAcademicPlan(userUuid, uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("RPS Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *AcademicPlanService) UpdateAcademicPlan(userUuid, uuid string, req *request.UpdateAcademicPlan) error {

	result, err := s.GetAcademicPlan(userUuid, uuid)
	if err != nil {
		return err
	}

	model := models.AcademicPlan{
		ID:        result.ID,
		Available: req.Available,
		Note:      req.Note,
	}

	if err := s.Repo.UpdateAcademicPlan(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *AcademicPlanService) DeleteAcademicPlan(userUuid, uuid string) error {
	result, err := s.GetAcademicPlan(userUuid, uuid)
	if err != nil {
		return err
	}

	model := models.AcademicPlan{
		ID: result.ID,
	}

	if err := s.Repo.DeleteAcademicPlan(&model); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return response.VIOLATED_ERR
		}
		return response.SERVICE_INTERR
	}

	return nil
}
