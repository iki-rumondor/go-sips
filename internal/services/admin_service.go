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
	var user models.Pengguna
	condition := fmt.Sprintf("uuid = '%s'", req.UserUuid)
	if err := s.Repo.First(&user, condition); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	model := models.PembimbingAkademik{
		Nama:    req.Nama,
		ProdiID: user.Prodi.ID,
		// Pengguna: &models.Pengguna{
		// 	Username: req.Nip,
		// 	Password: req.Nip,
		// 	RoleID:   3,
		// },
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
			Uuid:     item.Uuid,
			Nama:     item.Nama,
			Username: item.Pengguna.Username,
		})
	}

	return &resp, nil
}

func (s *AdminService) FindPembimbingProdi(userUuid string) (*[]response.Pembimbing, error) {
	var user models.Pengguna
	condition := fmt.Sprintf("uuid = '%s'", userUuid)
	if err := s.Repo.First(&user, condition); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	var model []models.PembimbingAkademik
	condition = fmt.Sprintf("prodi_id = '%d'", user.Prodi.ID)
	if err := s.Repo.Find(&model, condition); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Pembimbing
	for _, item := range model {
		resp = append(resp, response.Pembimbing{
			Uuid:     item.Uuid,
			Nama:     item.Nama,
			Username: item.Pengguna.Username,
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
		Uuid:     model.Uuid,
		Nama:     model.Nama,
		Username: model.Pengguna.Username,
	}

	return &resp, nil
}

func (s *AdminService) UpdatePembimbing(uuid string, req *request.UpdatePembimbing) error {
	model := models.PembimbingAkademik{
		Uuid: uuid,
		Nama: req.Nama,
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

func (s *AdminService) CreateProdi(req *request.Prodi) error {
	model := models.Prodi{
		Name:    req.Name,
		Kaprodi: req.Kaprodi,
		Pengguna: &models.Pengguna{
			Username: req.Username,
			Password: req.Username,
			RoleID:   4,
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

func (s *AdminService) FindAllProdi() (*[]response.Prodi, error) {
	var model []models.Prodi

	if err := s.Repo.Find(&model, ""); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Prodi
	for _, item := range model {
		resp = append(resp, response.Prodi{
			Uuid:     item.Uuid,
			Name:     item.Name,
			Kaprodi:  item.Kaprodi,
			Username: item.Pengguna.Username,
		})
	}

	return &resp, nil
}

func (s *AdminService) FindProdi(uuid string) (*response.Prodi, error) {
	var model models.Prodi
	condition := fmt.Sprintf("uuid = '%s'", uuid)
	if err := s.Repo.First(&model, condition); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	resp := response.Prodi{
		Uuid:     model.Uuid,
		Name:     model.Name,
		Kaprodi:  model.Kaprodi,
		Username: model.Pengguna.Username,
	}

	return &resp, nil
}

func (s *AdminService) UpdateProdi(uuid string, req *request.Prodi) error {
	var prodi models.Prodi
	condition := fmt.Sprintf("uuid = '%s'", uuid)
	if err := s.Repo.First(&prodi, condition); err != nil {
		return response.SERVICE_INTERR
	}

	model := models.Prodi{
		ID:      prodi.ID,
		Name:    req.Name,
		Kaprodi: req.Kaprodi,
		Pengguna: &models.Pengguna{
			ID:       prodi.PenggunaID,
			Username: req.Username,
			Password: req.Username,
		},
	}

	if err := s.Repo.Update(&model, ""); err != nil {
		log.Println(err.Error())
		if utils.IsErrorType(err) {
			return err
		}
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *AdminService) DeleteProdi(uuid string) error {
	var model models.Prodi
	condition := fmt.Sprintf("uuid = '%s'", uuid)

	if err := s.Repo.First(&model, condition); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	if err := s.Repo.Delete(&model.Pengguna, []string{"Prodi"}); err != nil {
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

	var prodi models.Prodi
	condition = fmt.Sprintf("id = '%d'", user.PembimbingAkademik.ProdiID)
	if err := s.Repo.First(&prodi, condition); err != nil {
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
		condition = fmt.Sprintf("angkatan = '%s' AND pembimbing_akademik_id = '%d'", item, user.PembimbingAkademik.ID)
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
		"prodi":          prodi.Name,
	}

	return resp, nil
}

func (s *AdminService) GetKaprodiDashboard(userUuid string) (map[string]interface{}, error) {
	var user models.Pengguna
	condition := fmt.Sprintf("uuid = '%s'", userUuid)
	if err := s.Repo.First(&user, condition); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	var listAngkatan []string
	var amountAngkatan []int

	if err := s.Repo.DistinctProdiMahasiswa(user.Prodi.ID, &listAngkatan, "angkatan"); err != nil {
		return nil, response.SERVICE_INTERR
	}

	for _, item := range listAngkatan {
		var mahasiswa []models.Mahasiswa
		condition := fmt.Sprintf("angkatan = '%s'", item)
		if err := s.Repo.FindProdiMahasiswa(user.Prodi.ID, &mahasiswa, condition); err != nil {
			log.Println(err.Error())
			return nil, response.SERVICE_INTERR
		}
		amountAngkatan = append(amountAngkatan, len(mahasiswa))
	}

	var dropOut []models.Mahasiswa
	rule := time.Now().Year() - 5
	condition = fmt.Sprintf("angkatan < '%d' ", rule)
	if err := s.Repo.FindProdiMahasiswa(user.Prodi.ID, &dropOut, condition); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	var percepatan []models.Mahasiswa
	condition = fmt.Sprintf("percepatan = %v", true)
	if err := s.Repo.FindProdiMahasiswa(user.Prodi.ID, &percepatan, condition); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	resp := map[string]interface{}{
		"prodi":          user.Prodi.Name,
		"listAngkatan":   listAngkatan,
		"amountAngkatan": amountAngkatan,
		"do":             len(dropOut),
		"percepatan":     len(percepatan),
	}

	return resp, nil
}

func (s *AdminService) GetKajurDashboard() (map[string]interface{}, error) {

	var listAngkatan []string
	var amountAngkatan []int
	var listProdi []string
	var amountProdi []int

	if err := s.Repo.Distinct(&models.Mahasiswa{}, "angkatan", "", &listAngkatan); err != nil {
		return nil, response.SERVICE_INTERR
	}

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

	var prodi []models.Prodi
	if err := s.Repo.Find(&prodi, ""); err != nil {
		return nil, response.SERVICE_INTERR
	}

	for _, item := range prodi {
		var mahasiswa []models.Mahasiswa
		if err := s.Repo.FindProdiMahasiswa(item.ID, &mahasiswa, ""); err != nil {
			log.Println(err.Error())
			return nil, response.SERVICE_INTERR
		}

		listProdi = append(listProdi, item.Name)
		amountProdi = append(amountProdi, len(mahasiswa))
	}

	resp := map[string]interface{}{
		"listProdi":      listProdi,
		"amountProdi":    amountProdi,
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

func (s *AdminService) GetPengaturanByName(name string) (*response.Pengaturan, error) {
	var model models.Pengaturan

	if err := s.Repo.First(&model, fmt.Sprintf("name = '%s'", name)); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	resp := response.Pengaturan{
		Uuid:  model.Uuid,
		Name:  model.Name,
		Value: model.Value,
	}

	return &resp, nil
}

func (s *AdminService) FindAllUsers() (*[]response.User, error) {
	var model []models.Pengguna

	if err := s.Repo.Find(&model, ""); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	var resp []response.User
	for _, item := range model {
		resp = append(resp, response.User{
			Uuid:     item.Uuid,
			Username: item.Username,
			Role:     item.Role.Nama,
			RoleID:   item.Role.ID,
		})
	}

	return &resp, nil
}

func (s *AdminService) CreateKajur(req *request.Kajur) error {
	if req.Password != req.ConfirmPassword {
		return response.BADREQ_ERR("Password tidak sama")
	}

	model := models.Pengguna{
		Username: req.Username,
		Password: req.Password,
		RoleID:   req.RoleID,
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

func (s *AdminService) UpdateUsername(uuid string, req *request.UpdateUsername) error {
	var user models.Pengguna
	condition := fmt.Sprintf("uuid = '%s'", uuid)
	if err := s.Repo.First(&user, condition); err != nil {
		return response.SERVICE_INTERR
	}

	model := models.Pengguna{
		ID:       user.ID,
		Username: req.Username,
	}

	if err := s.Repo.Update(&model, ""); err != nil {
		log.Println(err.Error())
		if utils.IsErrorType(err) {
			return err
		}
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *AdminService) UpdatePassword(uuid string, req *request.UpdatePassword) error {
	if req.NewPassword != req.ConfirmPassword {
		return response.BADREQ_ERR("Password tidak sama")
	}

	var user models.Pengguna
	condition := fmt.Sprintf("uuid = '%s'", uuid)
	if err := s.Repo.First(&user, condition); err != nil {
		return response.SERVICE_INTERR
	}

	if err := utils.ComparePassword(user.Password, req.CurrentPassword); err != nil {
		return response.BADREQ_ERR("Password lama tidak sesuai")
	}

	model := models.Pengguna{
		ID:       user.ID,
		Password: req.NewPassword,
	}

	if err := s.Repo.Update(&model, ""); err != nil {
		log.Println(err.Error())
		if utils.IsErrorType(err) {
			return err
		}
		return response.SERVICE_INTERR
	}

	return nil
}
