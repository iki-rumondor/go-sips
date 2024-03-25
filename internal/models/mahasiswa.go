package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Mahasiswa struct {
	ID                   uint    `gorm:"primaryKey"`
	PembimbingAkademikID uint    `gorm:"not_null"`
	PenggunaID           uint    `gorm:"not_null"`
	Uuid                 string  `gorm:"not_null;size:64"`
	Nim                  string  `gorm:"not_null;unique;size:9"`
	Nama                 string  `gorm:"not_null;size:64"`
	Angkatan             uint    `gorm:"not_null"`
	TotalSks             uint    `gorm:"not_null"`
	Ipk                  float64 `gorm:"not_null"`
	JumlahError          uint    `gorm:"not_null"`
	Class                string  `gorm:"null;size:8"`
	Percepatan           bool    `gorm:"not_null;size:8"`
	CreatedAt            int64   `gorm:"autoCreateTime:milli"`
	UpdatedAt            int64   `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	PembimbingAkademik   *PembimbingAkademik
	Pengguna             *Pengguna
}

func (Mahasiswa) TableName() string {
	return "mahasiswa"
}

func (m *Mahasiswa) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}

func (m *Mahasiswa) AfterDelete(tx *gorm.DB) error {
	if err := tx.Delete(&Pengguna{}, "id = ?", m.PenggunaID).Error; err != nil {
		return err
	}
	return nil
}
