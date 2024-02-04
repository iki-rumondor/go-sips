package models

import (
	"github.com/google/uuid"
	"github.com/iki-rumondor/sips/internal/utils"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"not_null,unique;size:64"`
	Username  string `gorm:"not_null;size:16"`
	Password  string `gorm:"not_null;size:64"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
}

func (Admin) TableName() string {
	return "admin"
}

func (m *Admin) BeforeSave(tx *gorm.DB) error {
	hashPass, err := utils.HashPassword(m.Password)
	if err != nil {
		return err
	}
	m.Password = hashPass
	m.Uuid = uuid.NewString()
	return nil
}
