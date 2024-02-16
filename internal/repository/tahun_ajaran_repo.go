package repository

import (
	"fmt"

	"github.com/iki-rumondor/sips/internal/interfaces"
	"github.com/iki-rumondor/sips/internal/models"
	"gorm.io/gorm"
)

type TahunAjaranRepository struct {
	db *gorm.DB
}

func NewTahunAjaranRepository(db *gorm.DB) interfaces.TahunAjaranRepoInterface {
	return &TahunAjaranRepository{
		db: db,
	}
}

func (r *TahunAjaranRepository) FindTahunAjaran() (*[]models.TahunAjaran, error) {
	var model []models.TahunAjaran
	if err := r.db.Find(&model).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *TahunAjaranRepository) FindTahunAjaranBy(column string, value interface{}) (*models.TahunAjaran, error) {
	var model models.TahunAjaran
	if err := r.db.First(&model, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *TahunAjaranRepository) CreateTahunAjaran(model *models.TahunAjaran) error {
	return r.db.Create(model).Error
}

func (r *TahunAjaranRepository) UpdateTahunAjaran(model *models.TahunAjaran) error {
	return r.db.Updates(model).Error
}

func (r *TahunAjaranRepository) DeleteTahunAjaran(model *models.TahunAjaran) error {
	return r.db.Delete(model).Error
}
