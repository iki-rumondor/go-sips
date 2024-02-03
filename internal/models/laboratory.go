package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Laboratory struct {
	ID           uint   `gorm:"primaryKey"`
	Uuid         string `gorm:"not_null;unique;size:64"`
	DepartmentID uint   `gorm:"not_null"`
	Name         string `gorm:"not_null;size:32"`
	CreatedAt    int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt    int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Department   *Department
}

func (m *Laboratory) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
