package migrate

import (
	"fmt"
	"log"

	"github.com/iki-rumondor/sips/internal/models"
	"gorm.io/gorm"
)

func ReadTerminal(db *gorm.DB, args []string) {
	switch {
	case args[1] == "fresh":
		if err := freshDatabase(db); err != nil {
			log.Fatal(err.Error())
		}
	case args[1] == "migrate":
		if err := migrateDatabase(db); err != nil {
			log.Fatal(err.Error())
		}
	case args[1] == "seed":
		if err := seedDatabase(db); err != nil {
			log.Fatal(err.Error())
		}
	default:
		fmt.Println("Hello, Nice to meet you")
	}
}

func freshDatabase(db *gorm.DB) error {
	for _, model := range GetAllModels() {
		if err := db.Migrator().DropTable(model.Model); err != nil {
			return err
		}
	}

	for _, model := range GetAllModels() {
		if err := db.Debug().AutoMigrate(model.Model); err != nil {
			return err
		}
	}

	db.Create(&models.Role{
		Nama: "ADMIN",
		Pengguna: &[]models.Pengguna{
			{
				Username: "kajur",
				Password: "123",
			},
		},
	})

	db.Create(&models.Role{
		Nama: "MAHASISWA",
	})

	db.Create(&models.Role{
		Nama: "PA",
	})

	db.Create(&models.Role{
		Nama: "KAPRODI",
	})

	db.Create([]models.Pengaturan{
		{
			Name:  "angkatan_percepatan",
			Value: fmt.Sprintf("%d", db.NowFunc().Year()-3),
		},
		{
			Name:  "angkatan_kelas",
			Value: fmt.Sprintf("%d", db.NowFunc().Year()-2),
		},
		{
			Name:  "total_sks",
			Value: "120",
		},
		{
			Name:  "jumlah_error",
			Value: "0",
		},
		{
			Name:  "ipk",
			Value: "3",
		},
		{
			Name:  "jumlah_mahasiswa",
			Value: "20",
		},
		{
			Name:  "maksimal_percepatan",
			Value: "10",
		},
	})

	return nil
}

func migrateDatabase(db *gorm.DB) error {
	for _, model := range GetAllModels() {
		if err := db.Debug().AutoMigrate(model.Model); err != nil {
			return err
		}
	}
	return nil
}

func seedDatabase(db *gorm.DB) error {
	sisfor := models.Prodi{
		Name:    "Sistem Informasi",
		Kaprodi: "Muchlis Polin",
		Pengguna: &models.Pengguna{
			Username: "sis123",
			Password: "sis123",
			RoleID:   4,
		},
	}

	pti := models.Prodi{
		Name:    "Pendidikan Teknologi Informasi",
		Kaprodi: "Arif Dwinanto",
		Pengguna: &models.Pengguna{
			Username: "pti123",
			Password: "pti123",
			RoleID:   4,
		},
	}

	if err := db.Create(&sisfor).Error; err != nil {
		return err
	}

	if err := db.Create(&pti).Error; err != nil {
		return err
	}

	return nil
}
