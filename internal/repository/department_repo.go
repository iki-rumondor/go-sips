package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"github.com/iki-rumondor/go-monev/internal/utils"
	"gorm.io/gorm"
)

type DepartmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) interfaces.DepartmentRepoInterface {
	return &DepartmentRepository{
		db: db,
	}
}

func (r *DepartmentRepository) FindDepartments() (*[]models.Department, error) {
	var departments []models.Department
	if err := r.db.Preload("Major").Preload("Subjects").Preload("User.Role").Find(&departments).Error; err != nil {
		return nil, err
	}

	return &departments, nil
}

func (r *DepartmentRepository) FindDepartmentBy(column string, value interface{}) (*models.Department, error) {
	var department models.Department
	if err := r.db.Preload("Major").Preload("Subjects").Preload("User.Role").First(&department, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &department, nil
}

func (r *DepartmentRepository) FindMajorBy(column string, value interface{}) (*models.Major, error) {
	var major models.Major
	if err := r.db.First(&major, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &major, nil
}

func (r *DepartmentRepository) CreateDepartment(department *models.Department) error {
	
	randString := utils.GenerateRandomString(7)
	user := models.User{
		Username:   randString,
		Password:   randString,
		RoleID:     2,
		Department: department,
	}

	return r.db.Create(&user).Error
}

func (r *DepartmentRepository) UpdateDepartment(department *models.Department) error {
	return r.db.Model(department).Updates(department).Error
}

func (r *DepartmentRepository) DeleteDepartment(department *models.Department) error {
	return r.db.Delete(department).Error
}
