package models

import (
	"github.com/google/uuid"
	"github.com/iki-rumondor/sips/internal/http/response"
	"gorm.io/gorm"
)

type PembimbingAkademik struct {
	ID         uint   `gorm:"primaryKey"`
	Uuid       string `gorm:"not_null;unique;size:64"`
	Nama       string `gorm:"not_null;size:32"`
	Nip        string `gorm:"not_null;unique;size:32"`
	PenggunaID uint `gorm:"not_null"`
	CreatedAt  int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt  int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Pengguna   *Pengguna
}

func (PembimbingAkademik) TableName() string {
	return "penasihat_akademik"
}

func (m *PembimbingAkademik) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}

func (m *PembimbingAkademik) BeforeSave(tx *gorm.DB) error {
	if result := tx.First(&PembimbingAkademik{}, "nip = ? AND uuid != ?", m.Nip, m.Uuid).RowsAffected; result > 0 {
		return response.BADREQ_ERR("Nip Sudah Terdaftar Untuk Pembimbing Akademik")
	}

	return nil
}

func (m *PembimbingAkademik) BeforeUpdate(tx *gorm.DB) error {
	var model PembimbingAkademik
	if err := tx.First(&model, "uuid = ?", m.Uuid).Error; err != nil {
		return err
	}

	pengguna := Pengguna{
		ID:       model.PenggunaID,
		Username: m.Nip,
		Password: m.Nip,
	}

	if err := tx.Updates(&pengguna).Error; err != nil {
		return err
	}

	return nil
}
