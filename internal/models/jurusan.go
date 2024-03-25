package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Jurusan struct {
	ID          uint   `gorm:"primaryKey"`
	Uuid        string `gorm:"not_null;unique;size:64"`
	PenggunaID  uint   `gorm:"not_null"`
	KajurName   uint   `gorm:"not_null"`
	CreatedAt   int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt   int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Mahasiswa   *Mahasiswa
}

func (Jurusan) TableName() string {
	return "jurusan"
}

func (m *Jurusan) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
