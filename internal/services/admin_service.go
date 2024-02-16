package services

import (
	"errors"
	"log"
	"strconv"

	"github.com/iki-rumondor/sips/internal/http/request"
	"github.com/iki-rumondor/sips/internal/http/response"
	"github.com/iki-rumondor/sips/internal/interfaces"
	"github.com/iki-rumondor/sips/internal/models"
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

func (s *AdminService) SetMahasiswaPercepatan(req *request.PercepatanCond) error {
	angkatan, _ := strconv.Atoi(req.Angkatan)
	totalSks, _ := strconv.Atoi(req.TotalSks)
	jumlahError, _ := strconv.Atoi(req.JumlahError)

	ipk, err := utils.StringToFloat(req.Ipk)
	if err != nil {
		return response.BADREQ_ERR("Nilai Ipk Tidak Valid")
	}

	mahasiswa, err := s.Repo.FindMahasiswaByRule(ipk, uint(totalSks), uint(jumlahError), uint(angkatan))
	if err != nil {
		return response.SERVICE_INTERR
	}

	if len(*mahasiswa) == 0 {
		return response.NOTFOUND_ERR("Mahasiswa Dengan Kriteria Tersebut Tidak Ditemukan")
	}

	if err := s.Repo.CreateMahasiswaPercepatan(mahasiswa); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *AdminService) GetMahasiswaPercepatan() (*[]models.Percepatan, error) {
	result, err := s.Repo.FindMahasiswaPercepatan()
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}
