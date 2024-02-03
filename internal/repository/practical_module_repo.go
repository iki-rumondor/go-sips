package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type PracticalModuleRepository struct {
	db *gorm.DB
}

func NewPracticalModuleRepository(db *gorm.DB) interfaces.PracticalModuleRepoInterface {
	return &PracticalModuleRepository{
		db: db,
	}
}

func (r *PracticalModuleRepository) FindPracticalModules(userUuid string) (*[]models.PracticalModule, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result []models.PracticalModule

	if err := r.db.Joins("Subject").Preload("AcademicYear").Preload("Laboratory").Find(&result, "subject.department_id = ?", user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *PracticalModuleRepository) FindUserPracticalModule(userUuid, uuid string) (*models.PracticalModule, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var model models.PracticalModule
	if err := r.db.Joins("Subject").Preload("AcademicYear").Preload("Laboratory").First(&model, "practical_modules.uuid = ? AND subject.department_id = ?", uuid, user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *PracticalModuleRepository) FindSubjectBy(column string, value interface{}) (*models.Subject, error) {
	var model models.Subject
	if err := r.db.Preload("Department.User").First(&model, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *PracticalModuleRepository) FindAcademicYearBy(column string, value interface{}) (*models.AcademicYear, error) {
	var model models.AcademicYear
	if err := r.db.First(&model, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *PracticalModuleRepository) FindLaboratoryBy(column string, value interface{}) (*models.Laboratory, error) {
	var model models.Laboratory
	if err := r.db.First(&model, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *PracticalModuleRepository) CreatePracticalModule(model *models.PracticalModule) error {
	return r.db.Create(model).Error
}

func (r *PracticalModuleRepository) UpdatePracticalModule(model *models.PracticalModule) error {
	return r.db.Updates(model).Error
}

func (r *PracticalModuleRepository) DeletePracticalModule(model *models.PracticalModule) error {
	return r.db.Delete(model).Error
}
