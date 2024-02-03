package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type LaboratoryRepository struct {
	db *gorm.DB
}

func NewLaboratoryRepository(db *gorm.DB) interfaces.LaboratoryRepoInterface {
	return &LaboratoryRepository{
		db: db,
	}
}

func (r *LaboratoryRepository) FindLaboratories(userUuid string) (*[]models.Laboratory, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result []models.Laboratory
	if err := r.db.Preload("Department").Find(&result, "department_id = ?", user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *LaboratoryRepository) FindUserLaboratory(userUuid, uuid string) (*models.Laboratory, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result models.Laboratory
	if err := r.db.Preload("Department").First(&result, "uuid = ? AND department_id = ?", uuid, user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *LaboratoryRepository) FindUserBy(column string, value interface{}) (*models.User, error) {
	var result models.User
	if err := r.db.Preload("Department").First(&result, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *LaboratoryRepository) CreateLaboratory(model *models.Laboratory) error {
	return r.db.Create(model).Error
}

func (r *LaboratoryRepository) UpdateLaboratory(model *models.Laboratory) error {
	return r.db.Updates(model).Error
}

func (r *LaboratoryRepository) DeleteLaboratory(model *models.Laboratory) error {
	return r.db.Delete(model).Error
}
