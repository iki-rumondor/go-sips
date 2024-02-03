package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type AcademicYearRepository struct {
	db *gorm.DB
}

func NewAcademicYearRepository(db *gorm.DB) interfaces.AcademicYearRepoInterface {
	return &AcademicYearRepository{
		db: db,
	}
}

func (r *AcademicYearRepository) FindAcademicYears() (*[]models.AcademicYear, error) {
	var model []models.AcademicYear
	if err := r.db.Preload("AcademicPlans").Preload("PracticalTools").Preload("PracticalModules").Find(&model).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *AcademicYearRepository) FindAcademicYearBy(column string, value interface{}) (*models.AcademicYear, error) {
	var model models.AcademicYear
	if err := r.db.First(&model, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *AcademicYearRepository) CreateAcademicYear(model *models.AcademicYear) error {
	return r.db.Create(model).Error
}

func (r *AcademicYearRepository) UpdateAcademicYear(model *models.AcademicYear) error {
	return r.db.Updates(model).Error
}

func (r *AcademicYearRepository) DeleteAcademicYear(model *models.AcademicYear) error {
	return r.db.Delete(model).Error
}
