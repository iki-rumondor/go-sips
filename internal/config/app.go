package config

import (
	"github.com/iki-rumondor/sips/internal/http/handlers"
	"github.com/iki-rumondor/sips/internal/interfaces"
	"github.com/iki-rumondor/sips/internal/repository"
	"github.com/iki-rumondor/sips/internal/services"
	"gorm.io/gorm"
)

type Handlers struct {
	AdminHandler interfaces.AdminHandlerInterface
}

func GetAppHandlers(db *gorm.DB) *Handlers {

	admin_repo := repository.NewAdminRepository(db)
	admin_service := services.NewAdminService(admin_repo)
	admin_handler := handlers.NewAdminHandler(admin_service)

	return &Handlers{
		AdminHandler: admin_handler,
	}
}
