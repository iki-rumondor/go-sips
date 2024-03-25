package repository

import (
	"fmt"

	"github.com/iki-rumondor/sips/internal/interfaces"
	"github.com/iki-rumondor/sips/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepoInterface {
	return &AdminRepository{
		db: db,
	}
}

func (r *AdminRepository) FindPenggunaBy(column string, value interface{}) (*models.Pengguna, error) {
	var model models.Pengguna
	if err := r.db.First(&model, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *AdminRepository) FindMahasiswaByRule(ipk float64, total_sks, jumlah_error uint) (*[]models.Mahasiswa, error) {
	var mahasiswa []models.Mahasiswa
	yearRule := r.db.NowFunc().Year() - 3
	rules := "total_sks >= ? AND jumlah_error <= ? AND ipk >= ? AND angkatan >= ?"

	if err := r.db.Find(&mahasiswa, rules, total_sks, jumlah_error, ipk, yearRule).Order("angkatan DESC").Error; err != nil {
		return nil, err
	}

	return &mahasiswa, nil
}

func (r *AdminRepository) FindMahasiswaByAngkatan(tahun int) (*[]models.Mahasiswa, error) {
	var mahasiswa []models.Mahasiswa

	if err := r.db.Find(&mahasiswa, "angkatan <= ?", tahun).Order("angkatan DESC").Error; err != nil {
		return nil, err
	}

	return &mahasiswa, nil
}

func (r *AdminRepository) Truncate(tableName string) error {
	return r.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", tableName)).Error
}


func (r *AdminRepository) Create(data interface{}) error {
	return r.db.Create(data).Error
}

func (r *AdminRepository) Find(data interface{}, condition string) error {
	return r.db.Preload(clause.Associations).Find(data, condition).Error
}

func (r *AdminRepository) FindWithOrder(data interface{}, condition, order string) error {
	return r.db.Preload(clause.Associations).Order(order).Find(data, condition).Error
}

func (r *AdminRepository) First(data interface{}, condition string) error {
	return r.db.Preload(clause.Associations).First(data, condition).Error
}

func (r *AdminRepository) Update(data interface{}, condition string) error {
	return r.db.Where(condition).Updates(data).Error
}

func (r *AdminRepository) Delete(data interface{}, assoc []string) error {
	return r.db.Select(assoc).Delete(data).Error
}

func (r *AdminRepository) Distinct(model interface{}, column, condition string, dest *[]string) error {
	return r.db.Model(model).Distinct().Where(condition).Pluck(column, dest).Error
}
