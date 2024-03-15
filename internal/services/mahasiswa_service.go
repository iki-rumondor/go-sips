package services

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/iki-rumondor/sips/internal/http/request"
	"github.com/iki-rumondor/sips/internal/http/response"
	"github.com/iki-rumondor/sips/internal/interfaces"
	"github.com/iki-rumondor/sips/internal/models"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type MahasiswaService struct {
	Repo interfaces.MahasiswaRepoInterface
}

func NewMahasiswaService(repo interfaces.MahasiswaRepoInterface) interfaces.MahasiswaServiceInterface {
	return &MahasiswaService{
		Repo: repo,
	}
}

func (s *MahasiswaService) ImportMahasiswa(pembimbingUuid, pathFile string) (*[]response.FailedImport, error) {
	var pembimbing models.PembimbingAkademik
	condition := fmt.Sprintf("uuid = '%s'", pembimbingUuid)

	if err := s.Repo.First(&pembimbing, condition); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	f, err := excelize.OpenFile(pathFile)
	if err != nil {
		log.Println("Gagal Membuka File")
		return nil, response.SERVICE_INTERR
	}
	defer f.Close()

	rows, err := f.GetRows("Mahasiswa")
	if err != nil {
		log.Println("Failed to get rows Mahasiswa")
		return nil, response.SERVICE_INTERR
	}

	var failedImport []response.FailedImport

	for i := 1; i < len(rows); i++ {
		cols := rows[i]
		angkatan, err := strconv.Atoi(cols[2])
		if err != nil {
			failedImport = append(failedImport, response.FailedImport{
				Nim:   cols[0],
				Nama:  cols[1],
				Pesan: "Angkatan Bukan Angka",
			})
			continue
		}

		totalSks, err := strconv.Atoi(cols[3])
		if err != nil {
			failedImport = append(failedImport, response.FailedImport{
				Nim:   cols[0],
				Nama:  cols[1],
				Pesan: "Total SKS Bukan Angka",
			})
			continue
		}

		ipk, err := strconv.ParseFloat(cols[4], 32)
		if err != nil {
			failedImport = append(failedImport, response.FailedImport{
				Nim:   cols[0],
				Nama:  cols[1],
				Pesan: "IPK Bukan Angka",
			})
			continue
		}

		jumlahError, err := strconv.Atoi(cols[5])
		if err != nil {
			failedImport = append(failedImport, response.FailedImport{
				Nim:   cols[0],
				Nama:  cols[1],
				Pesan: "Jumlah Error Bukan Angka",
			})
			continue
		}

		mahasiswa := models.Mahasiswa{
			Pengguna: &models.Pengguna{
				Username: cols[0],
				Password: cols[0],
				RoleID:   2,
			},
			PembimbingAkademikID: pembimbing.ID,
			Nim:                  cols[0],
			Nama:                 cols[1],
			Angkatan:             uint(angkatan),
			TotalSks:             uint(totalSks),
			Ipk:                  math.Round(ipk*100) / 100,
			JumlahError:          uint(jumlahError),
		}

		if err := s.Repo.CreateMahasiswa(&mahasiswa); err != nil {
			failedImport = append(failedImport, response.FailedImport{
				Nim:   cols[0],
				Nama:  cols[1],
				Pesan: "Gagal Menambahkan Mahasiswa: " + err.Error(),
			})
			continue
		}
	}

	return &failedImport, nil
}

func (s *MahasiswaService) GetAllMahasiswa(options map[string]string) (*[]response.Mahasiswa, error) {
	var model []models.Mahasiswa

	condition := fmt.Sprintf("angkatan LIKE '%%%s%%' AND class LIKE '%%%s%%'", options["angkatan"], options["class"])
	if err := s.Repo.Find(&model, condition, "nim"); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Mahasiswa

	for _, item := range model {
		resp = append(resp, response.Mahasiswa{
			Uuid:        item.Uuid,
			Nim:         item.Nim,
			Nama:        item.Nama,
			Kelas:       item.Class,
			Angkatan:    fmt.Sprintf("%d", item.Angkatan),
			Ipk:         fmt.Sprintf("%.2f", item.Ipk),
			TotalSks:    fmt.Sprintf("%d", item.TotalSks),
			JumlahError: fmt.Sprintf("%d", item.JumlahError),
			Pembimbing: &response.Pembimbing{
				Uuid: item.PembimbingAkademik.Uuid,
				Nama: item.PembimbingAkademik.Nama,
				Nip:  item.PembimbingAkademik.Nip,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})

	}

	return &resp, nil
}

func (s *MahasiswaService) GetMahasiswa(uuid string) (*models.Mahasiswa, error) {
	result, err := s.Repo.FindMahasiswaByUuid(uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Mahasiswa Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *MahasiswaService) UpdateMahasiswa(uuid string, req *request.Mahasiswa) error {
	var pembimbing models.PembimbingAkademik
	condition := fmt.Sprintf("uuid = '%s'", req.PembimbingUuid)

	if err := s.Repo.First(&pembimbing, condition); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	result, err := s.GetMahasiswa(uuid)
	if err != nil {
		return err
	}

	angkatan, _ := strconv.Atoi(req.Angkatan)
	totalSks, _ := strconv.Atoi(req.TotalSks)
	ipk, _ := strconv.ParseFloat(req.Ipk, 64)
	jumlahError, _ := strconv.Atoi(req.JumlahError)

	model := models.Mahasiswa{
		PembimbingAkademikID: pembimbing.ID,
		ID:                   result.ID,
		Nim:                  req.Nim,
		Nama:                 req.Nama,
		Angkatan:             uint(angkatan),
		TotalSks:             uint(totalSks),
		Ipk:                  ipk,
		JumlahError:          uint(jumlahError),
	}

	if err := s.Repo.UpdateMahasiswa(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *MahasiswaService) DeleteMahasiswa(uuid string) error {
	result, err := s.GetMahasiswa(uuid)
	if err != nil {
		return err
	}

	if err := s.Repo.DeleteMahasiswa(result); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return response.VIOLATED_ERR
		}
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *MahasiswaService) GetDataMahasiswa(nim string) (*response.DataMahasiswa, error) {
	result, err := s.Repo.FindBy("mahasiswa", "nim", nim)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Mahasiswa Dengan NIM Tersebut Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	resp := response.DataMahasiswa{
		Nim:         result["nim"].(string),
		Nama:        result["nama"].(string),
		Angkatan:    fmt.Sprintf("%d", result["angkatan"]),
		Ipk:         fmt.Sprintf("%.2f", result["ipk"]),
		TotalSks:    fmt.Sprintf("%d", result["total_sks"]),
		JumlahError: fmt.Sprintf("%d", result["jumlah_error"]),
		CreatedAt:   result["created_at"].(int64),
	}

	return &resp, nil
}

func (s *MahasiswaService) GetMahasiswaByUserUuid(userUuid string) (*response.Mahasiswa, error) {
	var user models.Pengguna
	condition := fmt.Sprintf("uuid = '%s'", userUuid)
	if err := s.Repo.First(&user, condition); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	var result models.Mahasiswa
	condition = fmt.Sprintf("id = '%d'", user.Mahasiswa.ID)
	if err := s.Repo.First(&result, condition); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	percepatan := true

	condition = fmt.Sprintf("mahasiswa_id = '%d'", result.ID)
	if err := s.Repo.First(&models.Percepatan{}, condition); err != nil {
		percepatan = false
	}

	resp := response.Mahasiswa{
		Uuid:        result.Uuid,
		Nim:         result.Nim,
		Nama:        result.Nama,
		Kelas:       result.Class,
		JumlahError: fmt.Sprintf("%d", result.JumlahError),
		Angkatan:    fmt.Sprintf("%d", result.Angkatan),
		Ipk:         fmt.Sprintf("%.2f", result.Ipk),
		TotalSks:    fmt.Sprintf("%d", result.TotalSks),
		Percepatan:  percepatan,
		Pembimbing: &response.Pembimbing{
			Uuid: result.PembimbingAkademik.Uuid,
			Nama: result.PembimbingAkademik.Nama,
			Nip:  result.PembimbingAkademik.Nip,
		},
	}

	return &resp, nil
}

func (s *MahasiswaService) GetAllMahasiswaByPenasihat(userUuid string, options map[string]string) (*[]response.Mahasiswa, error) {
	var user models.Pengguna
	condition := fmt.Sprintf("uuid = '%s'", userUuid)
	if err := s.Repo.First(&user, condition); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	var penasihat models.PembimbingAkademik
	condition = fmt.Sprintf("pengguna_id = '%d'", user.ID)
	if err := s.Repo.First(&penasihat, condition); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	var result []models.Mahasiswa

	condition = fmt.Sprintf("pembimbing_akademik_id = '%d' AND angkatan LIKE '%%%s%%' AND class LIKE '%%%s%%'", penasihat.ID, options["angkatan"], options["class"])

	if options["min_angkatan"] != "" {
		condition = fmt.Sprintf("pembimbing_akademik_id = '%d' AND angkatan LIKE '%%%s%%' AND class LIKE '%%%s%%' AND angkatan > '%s'", penasihat.ID, options["angkatan"], options["class"], options["min_angkatan"])
	}

	if err := s.Repo.Find(&result, condition, "nim"); err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Mahasiswa
	for _, item := range result {
		percepatan := true

		condition = fmt.Sprintf("mahasiswa_id = '%d'", item.ID)
		if err := s.Repo.First(&models.Percepatan{}, condition); err != nil {
			percepatan = false
		}

		resp = append(resp, response.Mahasiswa{
			Uuid:        item.Uuid,
			Nim:         item.Nim,
			Nama:        item.Nama,
			Kelas:       item.Class,
			JumlahError: fmt.Sprintf("%d", item.JumlahError),
			Angkatan:    fmt.Sprintf("%d", item.Angkatan),
			Ipk:         fmt.Sprintf("%.2f", item.Ipk),
			TotalSks:    fmt.Sprintf("%d", item.TotalSks),
			Percepatan:  percepatan,
			Pembimbing: &response.Pembimbing{
				Uuid: item.PembimbingAkademik.Uuid,
				Nama: item.PembimbingAkademik.Nama,
				Nip:  item.PembimbingAkademik.Nip,
			},
		})
	}

	return &resp, nil
}

// func (s *MahasiswaService) GetPercepatanByPenasihat(userUuid string) (*[]response.Mahasiswa, error) {
// 	var user models.Pengguna
// 	condition := fmt.Sprintf("uuid = '%s'", userUuid)
// 	if err := s.Repo.First(&user, condition); err != nil {
// 		log.Println(err.Error())
// 		return nil, response.SERVICE_INTERR
// 	}

// 	var penasihat models.PembimbingAkademik
// 	condition = fmt.Sprintf("pengguna_id = '%d'", user.ID)
// 	if err := s.Repo.First(&penasihat, condition); err != nil {
// 		log.Println(err.Error())
// 		return nil, response.SERVICE_INTERR
// 	}

// 	var result []models.Percepatan

// 	if err := s.Repo.FindPercepatanPenasihat(&result); err != nil {
// 		log.Println(err.Error())
// 		return nil, response.SERVICE_INTERR
// 	}

// 	var resp []response.Mahasiswa
// 	for _, item := range result {
// 		percepatan := true

// 		condition = fmt.Sprintf("mahasiswa_id = '%d'", item.ID)
// 		if err := s.Repo.First(&models.Percepatan{}, condition); err != nil {
// 			percepatan = false
// 		}

// 		resp = append(resp, response.Mahasiswa{
// 			Uuid:        item.Uuid,
// 			Nim:         item.Nim,
// 			Nama:        item.Nama,
// 			Kelas:       item.Class,
// 			JumlahError: fmt.Sprintf("%d", item.JumlahError),
// 			Angkatan:    fmt.Sprintf("%d", item.Angkatan),
// 			Ipk:         fmt.Sprintf("%.2f", item.Ipk),
// 			TotalSks:    fmt.Sprintf("%d", item.TotalSks),
// 			Percepatan:  percepatan,
// 			Pembimbing: &response.Pembimbing{
// 				Uuid: item.PembimbingAkademik.Uuid,
// 				Nama: item.PembimbingAkademik.Nama,
// 				Nip:  item.PembimbingAkademik.Nip,
// 			},
// 		})
// 	}

// 	return &resp, nil
// }
