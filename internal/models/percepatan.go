package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Percepatan struct {
	ID            uint   `gorm:"primaryKey"`
	Uuid          string `gorm:"not_null;unique;size:64"`
	TahunAjaranID uint   `gorm:"not_null"`
	MahasiswaID   uint   `gorm:"not_null"`
	CreatedAt     int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt     int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	TahunAjaran   *TahunAjaran
	Mahasiswa     *Mahasiswa
}

func (Percepatan) TableName() string {
	return "percepatan"
}

func (m *Percepatan) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
