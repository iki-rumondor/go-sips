package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PembimbingAkademik struct {
	ID         uint   `gorm:"primaryKey"`
	Uuid       string `gorm:"not_null;unique;size:64"`
	ProdiID    uint   `gorm:"not_null"`
	PenggunaID uint   `gorm:"not_null"`
	Nama       string `gorm:"not_null;size:64"`
	CreatedAt  int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt  int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Pengguna   *Pengguna
	Prodi      *Prodi
	Mahasiswa  *[]Mahasiswa
}

func (PembimbingAkademik) TableName() string {
	return "penasihat_akademik"
}

func (m *PembimbingAkademik) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}

func (m *PembimbingAkademik) BeforeUpdate(tx *gorm.DB) error {
	var model PembimbingAkademik
	if err := tx.First(&model, "uuid = ?", m.Uuid).Error; err != nil {
		return err
	}

	return nil
}
