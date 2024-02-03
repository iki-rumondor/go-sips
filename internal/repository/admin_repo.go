package repository

import (
	"fmt"

	"github.com/iki-rumondor/sips/internal/interfaces"
	"github.com/iki-rumondor/sips/internal/models"
	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepoInterface {
	return &AdminRepository{
		db: db,
	}
}

func (r *AdminRepository) FindAdminBy(column string, value interface{}) (*models.Admin, error) {
	var model models.Admin
	if err := r.db.Preload("Role").First(&model, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}
	return &model, nil
}
