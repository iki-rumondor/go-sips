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

type DepartmentService struct {
	Repo interfaces.DepartmentRepoInterface
}

func NewDepartmentService(repo interfaces.DepartmentRepoInterface) interfaces.DepartmentServiceInterface {
	return &DepartmentService{
		Repo: repo,
	}
}

func (s *DepartmentService) CreateDepartment(req *request.Department) error {
	major, err := s.Repo.FindMajorBy("uuid", req.MajorUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Jurusan tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	department := models.Department{
		MajorID: major.ID,
		Name:    req.Name,
		Head:    req.Head,
	}

	if err := s.Repo.CreateDepartment(&department); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *DepartmentService) GetAllDepartments() (*[]models.Department, error) {
	departments, err := s.Repo.FindDepartments()
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return departments, nil
}

func (s *DepartmentService) GetDepartment(uuid string) (*models.Department, error) {
	department, err := s.Repo.FindDepartmentBy("uuid", uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Program Studi Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return department, nil
}

func (s *DepartmentService) UpdateDepartment(uuid string, req *request.Department) error {
	major, err := s.Repo.FindMajorBy("uuid", req.MajorUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Jurusan tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	department, err := s.GetDepartment(uuid)
	if err != nil {
		return err
	}

	model := models.Department{
		ID:      department.ID,
		MajorID: major.ID,
		Name:    req.Name,
		Head:    req.Head,
	}

	if err := s.Repo.UpdateDepartment(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *DepartmentService) DeleteDepartment(uuid string) error {
	department, err := s.GetDepartment(uuid)
	if err != nil {
		return err
	}

	model := models.Department{
		ID: department.ID,
	}

	if err := s.Repo.DeleteDepartment(&model); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return response.VIOLATED_ERR
		}
		return response.SERVICE_INTERR
	}

	return nil
}
