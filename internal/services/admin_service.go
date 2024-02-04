package services

import (
	"errors"
	"log"

	"github.com/iki-rumondor/sips/internal/http/request"
	"github.com/iki-rumondor/sips/internal/http/response"
	"github.com/iki-rumondor/sips/internal/interfaces"
	"github.com/iki-rumondor/sips/internal/utils"
	"gorm.io/gorm"
)

type AdminService struct {
	Repo interfaces.AdminRepoInterface
}

func NewAdminService(repo interfaces.AdminRepoInterface) interfaces.AdminServiceInterface {
	return &AdminService{
		Repo: repo,
	}
}

func (s *AdminService) VerifyAdmin(req *request.SignIn) (string, error) {

	admin, err := s.Repo.FindAdminBy("username", req.Username)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", &response.Error{
				Code:    401,
				Message: "Username atau Password Salah",
			}
		}
		return "", response.SERVICE_INTERR
	}

	if err := utils.ComparePassword(admin.Password, req.Password); err != nil {
		return "", &response.Error{
			Code:    401,
			Message: "Username atau password salah",
		}
	}

	jwt, err := utils.GenerateToken(admin.Uuid)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

