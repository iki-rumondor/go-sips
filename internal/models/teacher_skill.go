package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeacherSkill struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"not_null,unique;size:64"`
	Skill     string `gorm:"not_null;size:64"`
	TeacherID uint   `gorm:"not_null"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Teacher   *Teacher
}

func (m *TeacherSkill) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
