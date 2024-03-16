package services

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"

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

func (s *AdminService) VerifyPengguna(req *request.SignIn) (string, error) {

	pengguna, err := s.Repo.FindPenggunaBy("username", req.Username)
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

	if err := utils.ComparePassword(pengguna.Password, req.Password); err != nil {
		return "", &response.Error{
			Code:    401,
			Message: "Username atau password salah",
		}
	}

	jwt, err := utils.GenerateToken(pengguna.Uuid)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (s *AdminService) GetUser(userUuid string) (*response.User, error) {
	var model models.Pengguna
	condition := fmt.Sprintf("uuid = '%s'", userUuid)
	if err := s.Repo.First(&model, condition); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	var mahasiswa response.DataMahasiswa
	var penasihat response.Pembimbing

	if model.Mahasiswa != nil {
		mahasiswa = response.DataMahasiswa{
			Nim:  model.Mahasiswa.Nim,
			Nama: model.Mahasiswa.Nama,
		}
	}

	if model.PembimbingAkademik != nil {
		penasihat = response.Pembimbing{
			Uuid: model.PembimbingAkademik.Uuid,
			Nama: model.PembimbingAkademik.Nama,
			Nip:  model.PembimbingAkademik.Nip,
		}
	}

	resp := response.User{
		Uuid:      model.Uuid,
		Username:  model.Username,
		Role:      model.Role.Nama,
		Mahasiswa: &mahasiswa,
		Penasihat: &penasihat,
	}

	return &resp, nil
}

func (s *AdminService) CreatePembimbing(req *request.Pembimbing) error {
	model := models.PembimbingAkademik{
		Nama: req.Nama,
		Nip:  req.Nip,
		Pengguna: &models.Pengguna{
			Username: req.Nip,
			Password: req.Nip,
			RoleID:   3,
		},
	}

	if err := s.Repo.Create(&model); err != nil {
		log.Println(err.Error())
		if utils.IsErrorType(err) {
			return err
		}
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *AdminService) FindAllPembimbing() (*[]response.Pembimbing, error) {
	var model []models.PembimbingAkademik

	if err := s.Repo.Find(&model, ""); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Pembimbing
	for _, item := range model {
		resp = append(resp, response.Pembimbing{
			Uuid: item.Uuid,
			Nama: item.Nama,
			Nip:  item.Nip,
		})
	}

	return &resp, nil
}

func (s *AdminService) FindPembimbing(uuid string) (*response.Pembimbing, error) {
	var model models.PembimbingAkademik
	condition := fmt.Sprintf("uuid = '%s'", uuid)
	if err := s.Repo.First(&model, condition); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	resp := response.Pembimbing{
		Uuid: model.Uuid,
		Nama: model.Nama,
		Nip:  model.Nip,
	}

	return &resp, nil
}

func (s *AdminService) UpdatePembimbing(uuid string, req *request.Pembimbing) error {
	model := models.PembimbingAkademik{
		Uuid: uuid,
		Nama: req.Nama,
		Nip:  req.Nip,
	}

	condition := fmt.Sprintf("uuid = '%s'", uuid)

	if err := s.Repo.Update(&model, condition); err != nil {
		log.Println(err.Error())
		if utils.IsErrorType(err) {
			return err
		}
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *AdminService) DeletePembimbing(uuid string) error {
	var model models.PembimbingAkademik
	condition := fmt.Sprintf("uuid = '%s'", uuid)

	if err := s.Repo.First(&model, condition); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	if err := s.Repo.Delete(&model.Pengguna, []string{"PembimbingAkademik"}); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *AdminService) UpdateKelas(req *request.KelasRule) error {

	years := utils.GeneratePastYears(3)
	var yearStrs []string
	for _, year := range years {
		yearStrs = append(yearStrs, strconv.Itoa(year))
	}

	for _, item := range yearStrs {
		var model []models.Mahasiswa
		condition := fmt.Sprintf("angkatan = %s", item)
		order := "nim ASC"

		if err := s.Repo.FindWithOrder(&model, condition, order); err != nil {
			log.Println(err.Error())
			return response.SERVICE_INTERR
		}

		classes := make(map[string][]models.Mahasiswa)
		studentAmount, _ := strconv.Atoi(req.JumlahMahasiswa)
		for i, student := range model {
			class := string(rune('A' + (i / studentAmount)))
			classes[class] = append(classes[class], student)
		}

		for class, students := range classes {
			for _, student := range students {
				model := models.Mahasiswa{
					ID:    student.ID,
					Class: class,
				}
				if err := s.Repo.Update(&model, ""); err != nil {
					return response.SERVICE_INTERR
				}
			}
		}
	}

	return nil
}

func (s *AdminService) GetClasses() ([]string, error) {
	var model models.Mahasiswa
	var resp []string

	if err := s.Repo.Distinct(&model, "class", "", &resp); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	sort.Strings(resp)

	return resp, nil
}

func (s *AdminService) GetPenasihatDashboard(userUuid string) (map[string]interface{}, error) {
	var user models.Pengguna
	condition := fmt.Sprintf("uuid = '%s'", userUuid)
	if err := s.Repo.First(&user, condition); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	var model models.Mahasiswa
	var listAngkatan []string
	var amountAngkatan []int

	condition = fmt.Sprintf("pembimbing_akademik_id = '%d'", user.PembimbingAkademik.ID)
	if err := s.Repo.Distinct(&model, "angkatan", condition, &listAngkatan); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	sort.Strings(listAngkatan)

	for _, item := range listAngkatan {
		var mahasiswa []models.Mahasiswa
		condition = fmt.Sprintf("angkatan = '%s'", item)
		if err := s.Repo.Find(&mahasiswa, condition); err != nil {
			log.Println(err.Error())
			return nil, response.SERVICE_INTERR
		}
		amountAngkatan = append(amountAngkatan, len(mahasiswa))
	}

	var dropOut []models.Mahasiswa
	rule := time.Now().Year() - 5
	condition = fmt.Sprintf("pembimbing_akademik_id = '%d' AND angkatan < '%d' ", user.PembimbingAkademik.ID, rule)
	if err := s.Repo.Find(&dropOut, condition); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	var percepatan []models.Mahasiswa
	condition = fmt.Sprintf("pembimbing_akademik_id = '%d' AND percepatan = %v", user.PembimbingAkademik.ID, true)
	if err := s.Repo.Find(&percepatan, condition); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	resp := map[string]interface{}{
		"listAngkatan":   listAngkatan,
		"amountAngkatan": amountAngkatan,
		"do":             len(dropOut),
		"percepatan":     len(percepatan),
	}

	return resp, nil
}

func (s *AdminService) GetKaprodiDashboard() (map[string]interface{}, error) {

	var model models.Mahasiswa
	var listAngkatan []string
	var amountAngkatan []int

	if err := s.Repo.Distinct(&model, "angkatan", "", &listAngkatan); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	sort.Strings(listAngkatan)

	for _, item := range listAngkatan {
		var mahasiswa []models.Mahasiswa
		condition := fmt.Sprintf("angkatan = '%s'", item)
		if err := s.Repo.Find(&mahasiswa, condition); err != nil {
			log.Println(err.Error())
			return nil, response.SERVICE_INTERR
		}
		amountAngkatan = append(amountAngkatan, len(mahasiswa))
	}

	var dropOut []models.Mahasiswa
	rule := time.Now().Year() - 5
	condition := fmt.Sprintf("angkatan < '%d' ", rule)
	if err := s.Repo.Find(&dropOut, condition); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	var percepatan []models.Mahasiswa
	condition = fmt.Sprintf("percepatan = %v", true)
	if err := s.Repo.Find(&percepatan, condition); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	resp := map[string]interface{}{
		"listAngkatan":   listAngkatan,
		"amountAngkatan": amountAngkatan,
		"do":             len(dropOut),
		"percepatan":     len(percepatan),
	}

	return resp, nil
}

func (s *AdminService) GetPengaturan() (*[]response.Pengaturan, error) {
	var model []models.Pengaturan

	if err := s.Repo.Find(&model, "id"); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Pengaturan
	for _, item := range model {
		resp = append(resp, response.Pengaturan{
			Uuid:  item.Uuid,
			Name:  item.Name,
			Value: item.Value,
		})
	}

	return &resp, nil
}
