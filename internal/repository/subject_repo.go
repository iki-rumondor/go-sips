package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type SubjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) interfaces.SubjectRepoInterface {
	return &SubjectRepository{
		db: db,
	}
}

func (r *SubjectRepository) FindSubjects(userUuid string) (*[]models.Subject, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result []models.Subject
	if err := r.db.Preload("Department").Preload("AcademicPlan.AcademicYear").Find(&result, "department_id = ?", user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *SubjectRepository) FindPracticalSubjects(userUuid string) (*[]models.Subject, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result []models.Subject
	if err := r.db.Preload("PracticalTool.AcademicYear").Preload("PracticalModule.AcademicYear").Find(&result, "department_id = ? AND practical = 1", user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *SubjectRepository) FindUserSubject(userUuid, uuid string) (*models.Subject, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result models.Subject
	if err := r.db.Preload("Department").Preload("AcademicPlan.AcademicYear").First(&result, "uuid = ? AND department_id = ?", uuid, user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *SubjectRepository) FindUserBy(column string, value interface{}) (*models.User, error) {
	var result models.User
	if err := r.db.Preload("Department").First(&result, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *SubjectRepository) CreateSubject(model *models.Subject) error {
	return r.db.Create(model).Error
}

func (r *SubjectRepository) UpdateSubject(model *models.Subject) error {
	return r.db.Updates(model).Error
}

func (r *SubjectRepository) DeleteSubject(model *models.Subject) error {

	return r.db.Delete(model).Error
}
