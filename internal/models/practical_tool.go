package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PracticalTool struct {
	ID             uint   `gorm:"primaryKey"`
	Uuid           string `gorm:"not_null;unique;size:64"`
	SubjectID      uint   `gorm:"not_null"`
	AcademicYearID uint   `gorm:"not_null"`
	Available      *bool   `gorm:"not_null"`
	Condition      string `gorm:"not_null;size:16;default:NONE"`
	Note           string `gorm:"not_null;size:144"`
	CreatedAt      int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt      int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Subject        *Subject
	AcademicYear   *AcademicYear
}

func (m *PracticalTool) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
