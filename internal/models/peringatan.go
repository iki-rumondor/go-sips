package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Peringatan struct {
	ID            uint   `gorm:"primaryKey"`
	Uuid          string `gorm:"not_null;unique;size:64"`
	TahunAjaranID uint   `gorm:"not_null"`
	MahasiswaID   uint   `gorm:"not_null"`
	CreatedAt     int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt     int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	TahunAjaran   *TahunAjaran
	Mahasiswa     *Mahasiswa
}

func (Peringatan) TableName() string {
	return "peringatan"
}

func (m *Peringatan) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
