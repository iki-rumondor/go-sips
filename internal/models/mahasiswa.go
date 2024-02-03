package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Mahasiswa struct {
	ID          uint    `gorm:"primaryKey"`
	Uuid        string  `gorm:"not_null;size:16"`
	Nim         string  `gorm:"not_null;unique;size:9"`
	Nama        string  `gorm:"not_null;size:64"`
	Angkatan    uint    `gorm:"not_null;size:4"`
	TotalSks    byte    `gorm:"not_null;size:3"`
	IPK         float32 `gorm:"not_null;size:3"`
	JumlahError byte    `gorm:"not_null;size:3"`
	CreatedAt   int64   `gorm:"autoCreateTime:milli"`
	UpdatedAt   int64   `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
}

func (Mahasiswa) TableName() string {
	return "mahasiswa"
}

func (m *Mahasiswa) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
