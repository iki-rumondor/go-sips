package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Department struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"not_null;unique;size:64"`
	MajorID   uint   `gorm:"not_null"`
	UserID    uint   `gorm:"not_null"`
	Name      string `gorm:"not_null;size:32"`
	Head      string `gorm:"not_null;size:32"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoCreateTime:milli; autoUpdateTime:milli"`
	Major     *Major
	User      *User
	Teachers  *[]Teacher
	Subjects  *[]Subject
}

func (m *Department) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
