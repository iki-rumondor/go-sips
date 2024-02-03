package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type AcademicPlanRepository struct {
	db *gorm.DB
}

func NewAcademicPlanRepository(db *gorm.DB) interfaces.AcademicPlanRepoInterface {
	return &AcademicPlanRepository{
		db: db,
	}
}

func (r *AcademicPlanRepository) FindAcademicPlans(userUuid string) (*[]models.AcademicPlan, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result []models.AcademicPlan

	if err := r.db.Joins("Subject").Preload("AcademicYear").Find(&result, "subject.department_id = ?", user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *AcademicPlanRepository) FindUserAcademicPlan(userUuid, uuid string) (*models.AcademicPlan, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var model models.AcademicPlan
	if err := r.db.Preload("AcademicYear").Joins("Subject").First(&model, "academic_plans.uuid = ? AND subject.department_id = ?", uuid, user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *AcademicPlanRepository) FindSubjectBy(column string, value interface{}) (*models.Subject, error) {
	var model models.Subject
	if err := r.db.Preload("Department.User").First(&model, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *AcademicPlanRepository) FindAcademicYearBy(column string, value interface{}) (*models.AcademicYear, error) {
	var model models.AcademicYear
	if err := r.db.First(&model, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *AcademicPlanRepository) CreateAcademicPlan(model *models.AcademicPlan) error {
	return r.db.Create(model).Error
}

func (r *AcademicPlanRepository) UpdateAcademicPlan(model *models.AcademicPlan) error {
	return r.db.Updates(model).Error
}

func (r *AcademicPlanRepository) DeleteAcademicPlan(model *models.AcademicPlan) error {
	return r.db.Delete(model).Error
}
