package models

type Role struct {
	ID       uint   `gorm:"primaryKey"`
	Nama     string `gorm:"not_null;size:16"`
	Pengguna *[]Pengguna
}
