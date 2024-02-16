package repository

import (
	"fmt"

	"github.com/iki-rumondor/sips/internal/interfaces"
	"github.com/iki-rumondor/sips/internal/models"
	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepoInterface {
	return &AdminRepository{
		db: db,
	}
}

func (r *AdminRepository) FindAdminBy(column string, value interface{}) (*models.Admin, error) {
	var model models.Admin
	if err := r.db.First(&model, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *AdminRepository) FindMahasiswaByRule(ipk float64, total_sks, jumlah_error, angkatan uint) (*[]models.Mahasiswa, error) {
	var mahasiswa []models.Mahasiswa
	rules := "angkatan = ? AND total_sks >= ? AND jumlah_error >= ? AND ipk >= ?"

	if err := r.db.Find(&mahasiswa, rules, angkatan, total_sks, jumlah_error, ipk).Error; err != nil {
		return nil, err
	}

	return &mahasiswa, nil
}

func (r *AdminRepository) FindMahasiswaPercepatan() (*[]models.Percepatan, error) {
	var result []models.Percepatan
	if err := r.db.Preload("Mahasiswa").Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *AdminRepository) CreateMahasiswaPercepatan(mahasiswa *[]models.Mahasiswa) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("TRUNCATE TABLE percepatan").Error; err != nil {
			return err
		}
		for _, item := range *mahasiswa {
			if err := tx.Create(&models.Percepatan{MahasiswaID: item.ID}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
