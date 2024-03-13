package repository

import (
	"fmt"

	"github.com/iki-rumondor/sips/internal/interfaces"
	"github.com/iki-rumondor/sips/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (r *MahasiswaRepository) FindAllMahasiswa(condtions string) (*[]models.Mahasiswa, error) {

	var result []models.Mahasiswa
	if err := r.db.Preload(clause.Associations).Order("nim").Find(&result, condtions).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *MahasiswaRepository) FindMahasiswaByUuid(uuid string) (*models.Mahasiswa, error) {

	var result models.Mahasiswa
	if err := r.db.Preload(clause.Associations).First(&result, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *MahasiswaRepository) UpdateMahasiswa(model *models.Mahasiswa) error {
	return r.db.Updates(model).Error
}

func (r *MahasiswaRepository) DeleteMahasiswa(model *models.Mahasiswa) error {
	return r.db.Select("Mahasiswa").Delete(model.Pengguna).Error
}

func (r *MahasiswaRepository) FindBy(tableName, column string, value interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := r.db.Table(tableName).Take(&result, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *MahasiswaRepository) First(data interface{}, condition string) error {
	return r.db.Preload(clause.Associations).First(data, condition).Error
}

func (r *MahasiswaRepository) Find(data interface{}, condition, order string) error {
	return r.db.Preload(clause.Associations).Order(order).Find(data, condition).Error
}

