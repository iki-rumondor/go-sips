package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type TeacherRepository struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) interfaces.TeacherRepoInterface {
	return &TeacherRepository{
		db: db,
	}
}

func (r *TeacherRepository) FindTeachers(userUuid string) (*[]models.Teacher, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result []models.Teacher
	if err := r.db.Preload("Department").Find(&result, "department_id = ?", user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *TeacherRepository) FindUserTeacher(userUuid, uuid string) (*models.Teacher, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result models.Teacher
	if err := r.db.Preload("Department").First(&result, "uuid = ? AND department_id = ?", uuid, user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *TeacherRepository) FindUserBy(column string, value interface{}) (*models.User, error) {
	var result models.User
	if err := r.db.Preload("Department").First(&result, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *TeacherRepository) CreateTeacher(model *models.Teacher) error {
	return r.db.Create(model).Error
}

func (r *TeacherRepository) UpdateTeacher(model *models.Teacher) error {
	return r.db.Updates(model).Error
}

func (r *TeacherRepository) DeleteTeacher(model *models.Teacher) error {
	return r.db.Delete(model).Error
}
