package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type MajorRepository struct {
	db *gorm.DB
}

func NewMajorRepository(db *gorm.DB) interfaces.MajorRepoInterface {
	return &MajorRepository{
		db: db,
	}
}

func (r *MajorRepository) FindMajors() (*[]models.Major, error) {
	var majors []models.Major
	if err := r.db.Find(&majors).Error; err != nil {
		return nil, err
	}

	return &majors, nil
}

func (r *MajorRepository) FindMajorBy(column string, value interface{}) (*models.Major, error) {
	var major models.Major
	if err := r.db.First(&major, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &major, nil
}

func (r *MajorRepository) CreateMajor(model *models.Major) error {
	return r.db.Create(model).Error
}

func (r *MajorRepository) UpdateMajor(model *models.Major) error {
	return r.db.Updates(model).Error
}

func (r *MajorRepository) DeleteMajor(model *models.Major) error {
	return r.db.Delete(model).Error
}
