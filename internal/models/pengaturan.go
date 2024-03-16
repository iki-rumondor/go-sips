package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pengaturan struct {
	ID    uint   `gorm:"primaryKey"`
	Uuid  string `gorm:"not_null;unique;size:64"`
	Name  string `gorm:"not_null"`
	Value string `gorm:"not_null"`
}

func (Pengaturan) TableName() string {
	return "pengaturan"
}

func (m *Pengaturan) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
