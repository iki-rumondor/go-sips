package models

import (
	"github.com/google/uuid"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"gorm.io/gorm"
)

type Subject struct {
	ID              uint   `gorm:"primaryKey"`
	Uuid            string `gorm:"not_null;unique;size:64"`
	Name            string `gorm:"not_null;size:32"`
	Code            string `gorm:"not_null;size:16"`
	DepartmentID    uint   `gorm:"not_null"`
	CreatedAt       int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt       int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Practical       *bool
	Department      *Department
	AcademicPlan    *AcademicPlan
	PracticalTool   *PracticalTool
	PracticalModule *PracticalModule
}

func (m *Subject) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}

func (m *Subject) BeforeSave(tx *gorm.DB) error {
	if result := tx.First(&Subject{}, "code = ? AND department_id = ? AND id != ?", m.Code, m.DepartmentID, m.ID).RowsAffected; result > 0 {
		return response.BADREQ_ERR("Kode Mata Kuliah Sudah Terdaftar")
	}
	return nil
}
