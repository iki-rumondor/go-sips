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

func (r *AdminRepository) FindMahasiswaPercepatan() (*[]models.Percepatan, error) {
	var result []models.Percepatan
	if err := r.db.Preload("Mahasiswa").Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *AdminRepository) CreateMahasiswaPercepatan(mahasiswa *[]models.Mahasiswa) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range *mahasiswa {
			if err := tx.Create(&models.Percepatan{MahasiswaID: item.ID}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *AdminRepository) Truncate(tableName string) error {
	return r.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", tableName)).Error
}

func (r *AdminRepository) FindMahasiswaPeringatan() (*[]models.Peringatan, error) {
	var result []models.Peringatan
	if err := r.db.Preload("Mahasiswa").Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *AdminRepository) CreateMahasiswaPeringatan(mahasiswa *[]models.Mahasiswa, year uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("TRUNCATE TABLE peringatan").Error; err != nil {
			return err
		}

		for _, item := range *mahasiswa {
			var status uint

			switch item.Angkatan {
			case year:
				status = 1
			case year - 1:
				status = 2
			case year - 2:
				status = 3
			default:
				status = 4
			}

			if err := tx.Create(&models.Peringatan{MahasiswaID: item.ID, Status: uint(status)}).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// func (r *AdminRepository) CreatePembimbing(data models.PembimbingAkademik) error {
// 	return r.db.Transaction(func(tx *gorm.DB) error {
// 		pengguna := models.Pengguna{
// 			Username: ,
// 		}
// 		if err := tx.Create(pengguna)
// 		return nil
// 	})
// }

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

func (r *AdminRepository) FindPenasihatPercepatan(dest *[]models.Percepatan, penasihatID uint) error {
	return r.db.Preload(clause.Associations).Joins("Mahasiswa").Find(dest, "mahasiswa.pembimbing_akademik_id = ?", penasihatID).Error
}
