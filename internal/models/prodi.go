package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Prodi struct {
	ID         uint   `gorm:"primaryKey"`
	Uuid       string `gorm:"not_null;unique;size:64"`
	PenggunaID uint   `gorm:"not_null"`
	Name       string `gorm:"not_null;size:32"`
	Kaprodi    string `gorm:"not_null;size:32"`
	CreatedAt  int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt  int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Pengguna   *Pengguna
}

func (Prodi) TableName() string {
	return "prodi"
}

func (m *Prodi) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}

func (m *Prodi) BeforeUpdate(tx *gorm.DB) error {
	if err := tx.Updates(m.Pengguna).Error; err != nil {
		return err
	}
	return nil
}
