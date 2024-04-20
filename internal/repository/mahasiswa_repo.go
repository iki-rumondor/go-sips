package repository

import (
	"fmt"
	"strconv"

	"github.com/iki-rumondor/sips/internal/interfaces"
	"github.com/iki-rumondor/sips/internal/models"
	"github.com/iki-rumondor/sips/internal/utils"
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
	if err := r.db.Preload(clause.Associations).Preload("PembimbingAkademik.Prodi").Order("nim").Find(&result, condtions).Error; err != nil {
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

func (r *MahasiswaRepository) FindLimit(data interface{}, condition, order string, limit int) error {
	return r.db.Preload(clause.Associations).Order(order).Limit(limit).Find(data, condition).Error
}

func (r *MahasiswaRepository) FindMahasiswaPercepatan(data *[]models.Mahasiswa, prodiID uint, limit int, order string) error {
	subQuery := r.db.Where("percepatan = true AND prodi_id = ?", prodiID).Model(&models.PembimbingAkademik{}).Select("id")
	return r.db.Preload(clause.Associations).Order(order).Limit(limit).Find(data, "pembimbing_akademik_id IN (?)", subQuery).Error
}


func (r *MahasiswaRepository) UpdatePengaturan(model *[]models.Pengaturan) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range *model {
			if err := tx.Model(&item).Where("name = ?", item.Name).Update("value", item.Value).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *MahasiswaRepository) UpdateKelas() error {
	var jumlahMahasiswa models.Pengaturan
	if err := r.db.Find(&jumlahMahasiswa, "name = ?", "jumlah_mahasiswa").Error; err != nil {
		return err
	}

	var angkatan models.Pengaturan
	if err := r.db.Find(&angkatan, "name = ?", "angkatan_kelas").Error; err != nil {
		return err
	}

	var prodi []models.Prodi
	if err := r.db.Find(&prodi).Error; err != nil {
		return err
	}

	jmlMahasiswa, _ := strconv.Atoi(jumlahMahasiswa.Value)
	angkatanInt, _ := strconv.Atoi(angkatan.Value)

	years := utils.GenerateYearsUntil(angkatanInt)

	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Mahasiswa{}).Where("1 = 1").Update("class", "").Error; err != nil {
			return err
		}

		for _, p := range prodi {
			subQuery := tx.Model(&models.PembimbingAkademik{}).Where("prodi_id = ?", p.ID).Select("id")
			for _, item := range years {
				var mahasiswa []models.Mahasiswa
				if err := tx.Order("nim ASC").Find(&mahasiswa, "angkatan = ? AND pembimbing_akademik_id IN (?)", item, subQuery).Error; err != nil {
					return err
				}

				classes := make(map[string][]models.Mahasiswa)
				for i, item := range mahasiswa {
					class := string(rune('A' + (i / jmlMahasiswa)))
					classes[class] = append(classes[class], item)
				}

				for class, students := range classes {
					for _, student := range students {
						model := models.Mahasiswa{
							ID:    student.ID,
							Class: class,
						}
						if err := tx.Updates(&model).Error; err != nil {
							return err
						}
					}
				}
			}

		}

		return nil
	})
}

func (r *MahasiswaRepository) UpdatePercepatan() error {
	var sks models.Pengaturan
	if err := r.db.Find(&sks, "name = ?", "total_sks").Error; err != nil {
		return err
	}

	var ipk models.Pengaturan
	if err := r.db.Find(&ipk, "name = ?", "ipk").Error; err != nil {
		return err
	}

	// var jumlahError models.Pengaturan
	// if err := r.db.Find(&jumlahError, "name = ?", "jumlah_error").Error; err != nil {
	// 	return err
	// }

	var angkatan models.Pengaturan
	if err := r.db.Find(&angkatan, "name = ?", "angkatan_percepatan").Error; err != nil {
		return err
	}

	// jmlError, _ := strconv.Atoi(jumlahError.Value)
	totalSks, _ := strconv.Atoi(sks.Value)
	ipkFloat, _ := strconv.ParseFloat(ipk.Value, 64)
	angkatanInt, _ := strconv.Atoi(angkatan.Value)

	var mahasiswa []models.Mahasiswa
	if err := r.db.Find(&mahasiswa, "total_sks >= ?  AND ipk >= ? AND angkatan >= ?", totalSks, ipkFloat, angkatanInt).Error; err != nil {
		return err
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Mahasiswa{}).Where("1 = 1").Update("percepatan", false).Error; err != nil {
			return err
		}

		for _, item := range mahasiswa {
			model := models.Mahasiswa{
				ID:         item.ID,
				Percepatan: true,
			}

			if err := tx.Updates(&model).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *MahasiswaRepository) Truncate(tableName string) error {
	return r.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", tableName)).Error
}

func (r *MahasiswaRepository) Delete(data interface{}, assoc []string) error {
	return r.db.Select(assoc).Delete(data).Error
}

func (r *MahasiswaRepository) FirstOrCreate(dest, model interface{}) error {
	return r.db.Where(model).FirstOrCreate(dest).Error
}

func (r *MahasiswaRepository) Create(data interface{}) error {
	return r.db.Create(data).Error
}

func (r *MahasiswaRepository) DeleteMahasiswaPengguna(data *[]models.Mahasiswa) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range *data {
			if err := r.db.Delete(item).Error; err != nil {
				return err
			}
			if err := r.db.Delete(&models.Pengguna{}, "id = ?", item.PenggunaID).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
