package config

import (
	"github.com/iki-rumondor/sips/internal/http/handlers"
	"github.com/iki-rumondor/sips/internal/interfaces"
	"github.com/iki-rumondor/sips/internal/repository"
	"github.com/iki-rumondor/sips/internal/services"
	"gorm.io/gorm"
)

type Handlers struct {
	AdminHandler       interfaces.AdminHandlerInterface
	MahasiswaHandler   interfaces.MahasiswaHandlerInterface
	TahunAjaranHandler interfaces.TahunAjaranHandlerInterface
}

func GetAppHandlers(db *gorm.DB) *Handlers {

	admin_repo := repository.NewAdminRepository(db)
	admin_service := services.NewAdminService(admin_repo)
	admin_handler := handlers.NewAdminHandler(admin_service)

	mahasiswa_repo := repository.NewMahasiswaRepository(db)
	mahasiswa_service := services.NewMahasiswaService(mahasiswa_repo)
	mahasiswa_handler := handlers.NewMahasiswaHandler(mahasiswa_service)

	tahun_ajaran_repo := repository.NewTahunAjaranRepository(db)
	tahun_ajaran_service := services.NewTahunAjaranService(tahun_ajaran_repo)
	tahun_ajaran_handler := handlers.NewTahunAjaranHandler(tahun_ajaran_service)

	return &Handlers{
		AdminHandler:       admin_handler,
		MahasiswaHandler:   mahasiswa_handler,
		TahunAjaranHandler: tahun_ajaran_handler,
	}
}
