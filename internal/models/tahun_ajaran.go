package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TahunAjaran struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"not_null;unique;size:64"`
	Tahun     uint   `gorm:"not_null"`
	Semester  string `gorm:"not_null;size:16"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
}

func (TahunAjaran) TableName() string {
	return "tahun_ajaran"
}

func (m *TahunAjaran) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
