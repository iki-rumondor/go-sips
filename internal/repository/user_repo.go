package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepoInterface {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindUserBy(column string, value interface{}) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Role").First(&user, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindSubjects() (*[]models.Subject, error) {
	var model []models.Subject
	if err := r.db.Find(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *UserRepository) FindPracticalSubjects() (*[]models.Subject, error) {
	var model []models.Subject
	if err := r.db.Find(&model, "practical = 1").Error; err != nil {
		return nil, err
	}
	return &model, nil
}
