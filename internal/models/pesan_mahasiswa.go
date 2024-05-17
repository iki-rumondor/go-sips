package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PesanMahasiswa struct {
	ID                   uint   `gorm:"primaryKey"`
	Uuid                 string `gorm:"not_null;unique;size:64"`
	MahasiswaID          uint   `gorm:"not_null"`
	PembimbingAkademikID uint   `gorm:"not_null"`
	Status               uint   `gorm:"not_null;size:2"`
	Message              string `gorm:"not_null;size:255"`
	CreatedAt            int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt            int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	PembimbingAkademik   *PembimbingAkademik
	Mahasiswa            *Mahasiswa
}

func (PesanMahasiswa) TableName() string {
	return "pesan_mahasiswa"
}

func (m *PesanMahasiswa) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
