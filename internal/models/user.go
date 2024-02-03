package models

import (
	"github.com/google/uuid"
	"github.com/iki-rumondor/go-monev/internal/utils"
	"gorm.io/gorm"
)

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Uuid       string `gorm:"not_null,unique;size:64"`
	Username   string `gorm:"not_null;size:16"`
	Password   string `gorm:"not_null;size:64"`
	RoleID     uint   `gorm:"not_null"`
	CreatedAt  int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt  int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Role       *Role
	Department *Department
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	hashPass, err := utils.HashPassword(m.Password)
	if err != nil {
		return err
	}
	m.Password = hashPass
	m.Uuid = uuid.NewString()
	return nil
}
