package repository

import (
	"github.com/iki-rumondor/sips/internal/interfaces"
	"github.com/iki-rumondor/sips/internal/models"
	"gorm.io/gorm"
)

type MahasiswaRepository struct {
	db *gorm.DB
}

func NewMahasiswaRepository(db *gorm.DB) interfaces.MahasiswaRepoInterface {
	return &MahasiswaRepository{
		db: db,
	}
}

func (r *MahasiswaRepository) CreateMahasiswa(model *models.Mahasiswa) error {
	return r.db.Create(model).Error
}

func (r *MahasiswaRepository) FindAllMahasiswa() (*[]models.Mahasiswa, error) {

	var result []models.Mahasiswa
	if err := r.db.Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *MahasiswaRepository) FindMahasiswaByUuid(uuid string) (*models.Mahasiswa, error) {

	var result models.Mahasiswa
	if err := r.db.First(&result, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *MahasiswaRepository) UpdateMahasiswa(model *models.Mahasiswa) error {
	return r.db.Updates(model).Error
}

func (r *MahasiswaRepository) DeleteMahasiswa(model *models.Mahasiswa) error {
	return r.db.Delete(model).Error
}
